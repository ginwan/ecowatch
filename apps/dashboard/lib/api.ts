import { Facility, Sensor } from "@/types";

const API_URL = "http://localhost:8080/api/v1";

export async function getSensors() {
    const res = await fetch(`${API_URL}/sensors`);
    return res.json();
}

export async function getSensor(id: number) {
    const res = await fetch(`${API_URL}/sensors/${id}`);
    return res.json();
}

export async function createSensor(data: Omit<Sensor, "id">) {
    const res = await fetch(`${API_URL}/sensors`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
    return res.json();
}

export async function updateSensor(id: number, data: Omit<Sensor, "id">) {
    const res = await fetch(`${API_URL}/sensors/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
    return res.json();
}

export async function deleteSensor(id: number) {
    await fetch(`${API_URL}/sensors/${id}`, { method: "DELETE" });
}

export async function getFacilities() {
    const res = await fetch(`${API_URL}/facilities`);
    return res.json();
}

export async function getFacility(id: number) {
    const res = await fetch(`${API_URL}/facilities/${id}`);
    return res.json();
}

export async function createFacility(data: Omit<Facility, "id">) {
    const res = await fetch(`${API_URL}/facilities`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
    return res.json();
}

export async function updateFacility(id: number, data: Omit<Facility, "id">) {
    const res = await fetch(`${API_URL}/facilities/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
    return res.json();
}

export async function deleteFacility(id: number) {
    await fetch(`${API_URL}/facilities/${id}`, { method: "DELETE" });
}

export async function getReadings() {
    const res = await fetch(`${API_URL}/readings`);
    return res.json();
}

export async function getAlerts() {
    const res = await fetch(`${API_URL}/alerts`);
    return res.json();
}