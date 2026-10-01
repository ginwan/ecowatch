"use client";

import { useForm, type UseFormRegister } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { createFacility } from "@/lib/api";

const schema = z.object({
    name: z.string().min(2, "Name is too short"),
    type: z.string().min(2, "Type is required"),
    location: z.string().min(2, "Location is required"),
    status: z.enum(["active", "inactive"]),
});

type FormData = z.infer<typeof schema>;

export default function NewFacilityPage() {
    const router = useRouter();
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<FormData>({
        resolver: zodResolver(schema),
        defaultValues: { status: "active" },
    });

    async function onSubmit(data: FormData) {
        await createFacility(data);
        router.push("/facilities");
        router.refresh();
    }

    return (
        <div className="p-10 max-w-3xl mx-auto">
            <Link href="/facilities" className="text-xs text-muted hover:text-text transition-colors">
                ← Back to facilities
            </Link>

            <h1 className="text-2xl font-semibold tracking-tight mt-4 mb-1">New Facility</h1>
            <p className="text-sm text-muted mb-8">Register a new site to the monitoring network.</p>

            <form onSubmit={handleSubmit(onSubmit)} className="bg-surface border border-border rounded-xl p-6">
                <div className="flex flex-col gap-5">
                    <Field label="Facility name" error={errors.name?.message}>
                        <input {...register("name")} placeholder="Ruwais Refinery" className={inputClass} />
                    </Field>

                    <div className="grid grid-cols-2 gap-4">
                        <Field label="Type" error={errors.type?.message}>
                            <input {...register("type")} placeholder="Refinery" className={inputClass} />
                        </Field>

                        <Field label="Location" error={errors.location?.message}>
                            <input {...register("location")} placeholder="Abu Dhabi" className={inputClass} />
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
                    <Link
                        href="/facilities"
                        className="text-sm text-muted hover:text-text px-4 py-2 rounded-lg transition-colors"
                    >
                        Cancel
                    </Link>
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="bg-accent text-white text-sm font-medium px-5 py-2 rounded-lg hover:opacity-90 transition-opacity disabled:opacity-50"
                    >
                        {isSubmitting ? "Creating…" : "Create Facility"}
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

function StatusOption({
    value,
    register,
}: {
    value: "active" | "inactive";
    register: UseFormRegister<FormData>;
}) {
    return (
        <label className="flex-1 relative">
            <input
                type="radio"
                value={value}
                {...register("status")}
                className="peer sr-only"
            />
            <div className="text-center text-sm py-2 rounded-lg border border-border cursor-pointer peer-checked:border-accent peer-checked:bg-accent/10 peer-checked:text-accent transition-colors capitalize">
                {value}
            </div>
        </label>
    );
}