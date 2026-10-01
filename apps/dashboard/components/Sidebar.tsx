"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Building2, Radio, AlertTriangle, Leaf } from "lucide-react";

const navItems = [
    { href: "/facilities", label: "Facilities", icon: Building2 },
    { href: "/sensors", label: "Sensors", icon: Radio },
    { href: "/alerts", label: "Alerts", icon: AlertTriangle },
];

export default function Sidebar() {
    const pathname = usePathname();

    return (
        <aside className="w-60 shrink-0 border-r border-border flex flex-col">
            <div className="flex items-center gap-2 px-6 py-6 border-b border-border">
                <div className="w-7 h-7 rounded-md bg-accent/15 flex items-center justify-center">
                    <Leaf size={16} className="text-accent" />
                </div>
                <span className="text-sm font-semibold tracking-tight">EcoWatch</span>
            </div>

            <nav className="flex flex-col gap-1 p-4">
                {navItems.map(({ href, label, icon: Icon }) => {
                    const isActive = pathname.startsWith(href);
                    return (
                        <Link
                            key={href}
                            href={href}
                            className={`flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors ${isActive
                                    ? "bg-accent/10 text-accent"
                                    : "text-text-muted hover:bg-surface hover:text-text"
                                }`}
                        >
                            <Icon size={16} />
                            {label}
                        </Link>
                    );
                })}
            </nav>
        </aside>
    );
}