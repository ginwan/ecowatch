import { ArrowUp, ArrowDown, Clock } from "lucide-react";
import { getAlerts } from "@/lib/api";
import { Alert } from "@/types";

export default async function AlertsPage() {
    const alerts: Alert[] = await getAlerts();
    const sorted = [...alerts].sort(
        (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
    );

    return (
        <div className="p-10">
            <h1 className="text-2xl font-semibold mb-1 tracking-tight">Alerts</h1>
            <p className="text-sm text-muted mb-10">{alerts.length} recorded events</p>

            <div className="flex flex-col gap-3">
                {sorted.map((alert) => {
                    const isHigh = alert.status === "high";
                    const time = new Date(alert.created_at).toLocaleString();
                    const Icon = isHigh ? ArrowUp : ArrowDown;

                    return (
                        <div
                            key={alert.id}
                            className={`flex items-center gap-4 bg-surface border border-border rounded-lg px-5 py-4 border-l-[3px] ${isHigh ? "border-l-danger" : "border-l-warning"
                                }`}
                        >
                            <span className={`shrink-0 flex items-center gap-1 text-xs font-medium px-2 py-0.5 rounded-full border uppercase tracking-wide ${isHigh
                                ? "border-danger/30 text-danger bg-danger/10"
                                : "border-warning/30 text-warning bg-warning/10"
                                }`}>
                                <Icon size={12} />
                                {alert.status}
                            </span>

                            <div className="flex-1 min-w-0">
                                <p className="text-sm">{alert.message}</p>
                                <p className="text-xs text-muted mt-0.5">Sensor #{alert.sensor_id} · {alert.value}</p>
                            </div>

                            <span className="shrink-0 flex items-center gap-1 text-xs text-text-muted/70 font-mono">
                                <Clock size={12} />
                                {time}
                            </span>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}