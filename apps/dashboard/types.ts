export type Sensor = {
    id: number;
    name: string;
    type: string;
    facility_id: number;
    unit: string;
    min_threshold: number;
    max_threshold: number;
    status: string;
};

export type Facility = {
    id: number;
    name: string;
    type: string;
    location: string;
    status: string;
};

export type Reading = {
    id: number;
    sensor_id: number;
    value: number;
    recorded_at: string;
    received_at: string;
};

export type Alert = {
    id: number;
    sensor_id: number;
    value: number;
    status: string;
    message: string;
    created_at: string;
};