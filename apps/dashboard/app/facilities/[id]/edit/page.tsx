"use client";

import { useEffect, useState } from "react";
import { useForm, type UseFormRegister } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useRouter, useParams } from "next/navigation";
import Link from "next/link";
import { getFacility, updateFacility } from "@/lib/api";

const schema = z.object({
    name: z.string().min(2, "Name is too short"),
    type: z.string().min(2, "Type is required"),
    location: z.string().min(2, "Location is required"),
    status: z.enum(["active", "inactive"]),
});

type FormData = z.infer<typeof schema>;

export default function EditFacilityPage() {
    const router = useRouter();
    const params = useParams();
    const id = Number(params.id);
    const [loaded, setLoaded] = useState(false);

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors, isSubmitting },
    } = useForm<FormData>({ resolver: zodResolver(schema) });

    useEffect(() => {
        getFacility(id).then((facility) => {
            reset({
                name: facility.name,
                type: facility.type,
                location: facility.location,
                status: facility.status,
            });
            setLoaded(true);
        });
    }, [id, reset]);

    async function onSubmit(data: FormData) {
        await updateFacility(id, data);
        router.push(`/facilities/${id}`);
        router.refresh();
    }

    if (!loaded) {
        return <div className="p-10 text-sm text-muted">Loading…</div>;
    }

    return (
        <div className="p-10 max-w-lg mx-auto">
            <Link href={`/facilities/${id}`} className="text-xs text-muted hover:text-text transition-colors">
                ← Back
            </Link>

            <h1 className="text-2xl font-semibold tracking-tight mt-4 mb-1">Edit Facility</h1>
            <p className="text-sm text-muted mb-8">Update site details.</p>

            <form onSubmit={handleSubmit(onSubmit)} className="bg-surface border border-border rounded-xl p-6">
                <div className="flex flex-col gap-5">
                    <Field label="Facility name" error={errors.name?.message}>
                        <input {...register("name")} className={inputClass} />
                    </Field>

                    <div className="grid grid-cols-2 gap-4">
                        <Field label="Type" error={errors.type?.message}>
                            <input {...register("type")} className={inputClass} />
                        </Field>
                        <Field label="Location" error={errors.location?.message}>
                            <input {...register("location")} className={inputClass} />
                        </Field>
                    </div>

                    <Field label="Status" error={errors.status?.message}>
                        <div className="flex gap-2">
                            <StatusOption value="active" register={register} />
                            <StatusOption value="inactive" register={register} />
                        </div>
                    </Field>
                </div>

                <div className="h-px bg-border my-6" />

                <div className="flex justify-end gap-3">
                    <Link href={`/facilities/${id}`} className="text-sm text-muted hover:text-text px-4 py-2 rounded-lg transition-colors">
                        Cancel
                    </Link>
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="bg-accent text-white text-sm font-medium px-5 py-2 rounded-lg hover:opacity-90 transition-opacity disabled:opacity-50"
                    >
                        {isSubmitting ? "Saving…" : "Save Changes"}
                    </button>
                </div>
            </form>
        </div>
    );
}

const inputClass =
    "w-full bg-bg border border-border rounded-lg px-3 py-2.5 text-sm placeholder:text-text-muted/40 focus:outline-none focus:border-accent transition-colors";

function Field({ label, error, children }: { label: string; error?: string; children: React.ReactNode }) {
    return (
        <div>
            <label className="block text-xs font-medium text-text-muted mb-2">{label}</label>
            {children}
            {error && <p className="text-xs text-danger mt-1.5">{error}</p>}
        </div>
    );
}

function StatusOption({ value, register }: { value: "active" | "inactive"; register: UseFormRegister<FormData> }) {
    return (
        <label className="flex-1 relative">
            <input type="radio" value={value} {...register("status")} className="peer sr-only" />
            <div className="text-center text-sm py-2 rounded-lg border border-border cursor-pointer peer-checked:border-accent peer-checked:bg-accent/10 peer-checked:text-accent transition-colors capitalize">
                {value}
            </div>
        </label>
    );
}