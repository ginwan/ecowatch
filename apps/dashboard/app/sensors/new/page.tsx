"use client";

import { useEffect, useState } from "react";
import { useForm, UseFormRegister } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { createSensor, getFacilities } from "@/lib/api";
import { Facility } from "@/types";

const schema = z.object({
    name: z.string().min(2, "Name is too short"),
    type: z.string().min(2, "Type is required"),
    facility_id: z.string().min(1, "Select a facility"),
    unit: z.string().min(1, "Unit is required"),
    min_threshold: z.string().min(1, "Required"),
    max_threshold: z.string().min(1, "Required"),
    status: z.enum(["active", "inactive"]),
});

type FormData = z.infer<typeof schema>;

export default function NewSensorPage() {
    const router = useRouter();
    const [facilities, setFacilities] = useState<Facility[]>([]);

    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<FormData>({
        resolver: zodResolver(schema),
        defaultValues: { status: "active" },
    });

    useEffect(() => {
        getFacilities().then(setFacilities);
    }, []);

    async function onSubmit(data: FormData) {
        if (Number(data.max_threshold) <= Number(data.min_threshold)) {
            alert("Max threshold must be greater than min threshold");
            return;
        }
        await createSensor({
            ...data,
            facility_id: Number(data.facility_id),
            min_threshold: Number(data.min_threshold),
            max_threshold: Number(data.max_threshold),
        });
        router.push("/sensors");
        router.refresh();
    }

    return (
        <div className="p-10 max-w-3xl mx-auto">
            <Link href="/sensors" className="text-xs text-muted hover:text-text transition-colors">
                ← Back to sensors
            </Link>

            <h1 className="text-2xl font-semibold tracking-tight mt-4 mb-1">New Sensor</h1>
            <p className="text-sm text-muted mb-8">Register a new sensor to a facility.</p>

            <form onSubmit={handleSubmit(onSubmit)} className="bg-surface border border-border rounded-xl p-6">
                <div className="flex flex-col gap-5">
                    <Field label="Sensor name" error={errors.name?.message}>
                        <input {...register("name")} placeholder="Boiler Temp Sensor" className={inputClass} />
                    </Field>

                    <div className="grid grid-cols-2 gap-4">
                        <Field label="Type" error={errors.type?.message}>
                            <input {...register("type")} placeholder="temperature" className={inputClass} />
                        </Field>
                        <Field label="Unit" error={errors.unit?.message}>
                            <input {...register("unit")} placeholder="°C" className={inputClass} />
                        </Field>
                    </div>

                    <Field label="Facility" error={errors.facility_id?.message}>
                        <select {...register("facility_id")} className={inputClass}>
                            <option value="">Select a facility</option>
                            {facilities.map((f) => (
                                <option key={f.id} value={f.id}>{f.name}</option>
                            ))}
                        </select>
                    </Field>

                    <div className="grid grid-cols-2 gap-4">
                        <Field label="Min threshold" error={errors.min_threshold?.message}>
                            <input type="number" step="any" {...register("min_threshold")} className={inputClass} />
                        </Field>
                        <Field label="Max threshold" error={errors.max_threshold?.message}>
                            <input type="number" step="any" {...register("max_threshold")} className={inputClass} />
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
                    <Link href="/sensors" className="text-sm text-muted hover:text-text px-4 py-2 rounded-lg transition-colors">
                        Cancel
                    </Link>
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="bg-accent text-white text-sm font-medium px-5 py-2 rounded-lg hover:opacity-90 transition-opacity disabled:opacity-50"
                    >
                        {isSubmitting ? "Creating…" : "Create Sensor"}
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