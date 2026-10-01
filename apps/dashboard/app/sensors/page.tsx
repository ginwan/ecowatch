import Link from "next/link";
import { Plus, Thermometer, Gauge, Wind, Radio } from "lucide-react";
import { getSensors } from "@/lib/api";
import { Sensor } from "@/types";

const typeIcons: Record<string, typeof Thermometer> = {
    temperature: Thermometer,
    pressure: Gauge,
    gas: Wind,
};

export default async function SensorsPage() {
    const sensors: Sensor[] = await getSensors();

    return (
        <div className="p-10">
            <div className="flex justify-between items-center mb-10">
                <div>
                    <h1 className="text-2xl font-semibold tracking-tight">Sensors</h1>
                    <p className="text-sm text-muted mt-1">{sensors.length} active sensors</p>
                </div>
                <Link
                    href="/sensors/new"
                    className="flex items-center gap-2 bg-accent text-white text-sm font-medium px-4 py-2 rounded-lg hover:opacity-90 transition-opacity"
                >
                    <Plus size={16} />
                    New Sensor
                </Link>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
                {sensors.map((sensor) => {
                    const isActive = sensor.status === "active";
                    const Icon = typeIcons[sensor.type] ?? Radio;
                    return (
                        <Link
                            key={sensor.id}
                            href={`/sensors/${sensor.id}`}
                            className="relative bg-surface border border-border rounded-xl p-5 overflow-hidden transition-all hover:border-accent/40 hover:-translate-y-0.5"
                        >
                            <div className={`absolute left-0 top-0 bottom-0 w-0.75 ${isActive ? "bg-safe" : "bg-text-muted"}`} />

                            <div className="flex items-start justify-between mb-4">
                                <div className="flex items-center gap-2">
                                    <Icon size={16} className="text-text-muted" />
                                    <h2 className="font-medium text-[15px] leading-snug">{sensor.name}</h2>
                                </div>
                                <span className={`shrink-0 text-xs px-2 py-0.5 rounded-full border ${isActive
                                    ? "border-safe/30 text-safe bg-safe/10"
                                    : "border-text-muted/30 text-muted bg-text-muted/10"
                                    }`}>
                                    {sensor.status}
                                </span>
                            </div>

                            <div className="text-[11px] uppercase tracking-wide text-text-muted/70 mb-3">
                                {sensor.type}
                            </div>

                            <div className="flex items-baseline gap-1 font-mono text-sm text-muted">
                                <span>{sensor.min_threshold}</span>
                                <span className="text-text-muted/50">—</span>
                                <span>{sensor.max_threshold}</span>
                                <span className="text-text-muted/70 ml-1">{sensor.unit}</span>
                            </div>
                        </Link>
                    );
                })}
            </div>
        </div>
    );
}