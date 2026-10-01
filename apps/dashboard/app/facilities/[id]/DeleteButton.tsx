"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Trash2 } from "lucide-react";
import { deleteFacility } from "@/lib/api";

export default function DeleteFacilityButton({ id }: { id: number }) {
    const router = useRouter();
    const [confirming, setConfirming] = useState(false);
    const [loading, setLoading] = useState(false);

    async function handleDelete() {
        setLoading(true);
        await deleteFacility(id);
        router.push("/facilities");
        router.refresh();
    }

    if (confirming) {
        return (
            <div className="flex items-center gap-2">
                <span className="text-xs text-muted">Delete this facility?</span>
                <button
                    onClick={handleDelete}
                    disabled={loading}
                    className="text-sm text-white bg-danger px-3 py-2 rounded-lg hover:opacity-90 transition-opacity"
                >
                    {loading ? "Deleting…" : "Confirm"}
                </button>
                <button
                    onClick={() => setConfirming(false)}
                    className="text-sm text-muted px-3 py-2 rounded-lg hover:text-text transition-colors"
                >
                    Cancel
                </button>
            </div>
        );
    }

    return (
        <button
            onClick={() => setConfirming(true)}
            className="flex items-center gap-2 text-sm text-danger px-4 py-2 rounded-lg border border-danger/30 hover:bg-danger/10 transition-colors"
        >
            <Trash2 size={14} />
            Delete
        </button>
    );
}