import Link from "next/link";
import { ArrowLeft, MapPin, Pencil } from "lucide-react";
import { getFacility } from "@/lib/api";
import DeleteFacilityButton from "./DeleteButton";

export default async function FacilityDetailPage({ params }: { params: Promise<{ id: string }> }) {
    const { id } = await params;
    const facility = await getFacility(Number(id));
    const isActive = facility.status === "active";

    return (
        <div className="p-10 max-w-lg">
            <Link href="/facilities" className="flex items-center gap-1 text-xs text-muted hover:text-text transition-colors mb-6">
                <ArrowLeft size={14} />
                Back to facilities
            </Link>

            <div className="bg-surface border border-border rounded-xl p-6">
                <div className="flex items-start justify-between mb-6">
                    <div>
                        <h1 className="text-xl font-semibold tracking-tight">{facility.name}</h1>
                        <p className="flex items-center gap-1 text-sm text-muted mt-1">
                            <MapPin size={13} />
                            {facility.location}
                        </p>
                    </div>
                    <span className={`shrink-0 text-xs px-2 py-0.5 rounded-full border ${isActive
                        ? "border-safe/30 text-safe bg-safe/10"
                        : "border-text-muted/30 text-muted bg-text-muted/10"
                        }`}>
                        {facility.status}
                    </span>
                </div>

                <div className="text-[11px] uppercase tracking-wide text-text-muted/70 mb-6">
                    {facility.type}
                </div>

                <div className="h-px bg-border mb-6" />

                <div className="flex justify-end gap-3">
                    <DeleteFacilityButton id={facility.id} />
                    <Link
                        href={`/facilities/${facility.id}/edit`}
                        className="flex items-center gap-2 bg-accent text-white text-sm font-medium px-4 py-2 rounded-lg hover:opacity-90 transition-opacity"
                    >
                        <Pencil size={14} />
                        Edit
                    </Link>
                </div>
            </div>
        </div>
    );
}