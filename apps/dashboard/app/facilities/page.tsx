import Link from "next/link";
import { Plus, MapPin } from "lucide-react";
import { getFacilities } from "@/lib/api";
import { Facility } from "@/types";

export default async function FacilitiesPage() {
    const facilities: Facility[] = await getFacilities();

    return (
        <div className="p-10">
            <div className="flex justify-between items-center mb-10">
                <div>
                    <h1 className="text-2xl font-semibold tracking-tight">Facilities</h1>
                    <p className="text-sm text-muted mt-1">{facilities.length} monitored sites across the network</p>
                </div>
                <Link
                    href="/facilities/new"
                    className="flex items-center gap-2 bg-accent text-white text-sm font-medium px-4 py-2 rounded-lg hover:opacity-90 transition-opacity"
                >
                    <Plus size={16} />
                    New Facility
                </Link>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
                {facilities.map((facility) => {
                    const isActive = facility.status === "active";
                    return (
                        <Link
                            key={facility.id}
                            href={`/facilities/${facility.id}`}
                            className="group relative bg-surface border border-border rounded-xl p-5 overflow-hidden transition-all hover:border-accent/40 hover:-translate-y-0.5"
                        >
                            <div className={`absolute left-0 top-0 bottom-0 w-0.75 ${isActive ? "bg-safe" : "bg-text-muted"}`} />

                            <div className="flex items-start justify-between mb-4">
                                <h2 className="font-medium text-[15px] leading-snug pr-2">{facility.name}</h2>
                                <span className={`shrink-0 text-xs px-2 py-0.5 rounded-full border ${isActive
                                    ? "border-safe/30 text-safe bg-safe/10"
                                    : "border-text-muted/30 text-muted bg-text-muted/10"
                                    }`}>
                                    {facility.status}
                                </span>
                            </div>

                            <div className="flex items-center gap-2 text-sm text-muted">
                                <span className="uppercase tracking-wide text-[11px] text-text-muted/70">{facility.type}</span>
                                <span className="w-1 h-1 rounded-full bg-border" />
                                <span className="flex items-center gap-1">
                                    <MapPin size={12} />
                                    {facility.location}
                                </span>
                            </div>
                        </Link>
                    );
                })}
            </div>
        </div>
    );
}