# EcoWatch

## Industrial Environmental Monitoring Platform

EcoWatch is an engineering project that monitors environmental data from
industrial facilities.

It simulates sensor readings, detects unusual values, stores data, and
displays alerts through a monitoring dashboard.

> EcoWatch is an educational project inspired by industrial monitoring systems.
> It is not a certified industrial safety or compliance system.


## Project Goals

- Practice backend development with Go.
- Understand systems-oriented software architecture.
- Work with PostgreSQL and relational data modeling.
- Build APIs for data ingestion and monitoring.
- Implement threshold-based anomaly detection.
- Explore monitoring, alerting, testing, and reliability.
- Apply software engineering concepts to energy and environmental technology.

## Architecture

EcoWatch consists of three main components:

- **Sensor Simulator:** Generates simulated environmental sensor readings.
- **Backend API:** Receives, validates, and processes sensor data.
- **Dashboard:** Displays sensor readings and alerts.

The backend stores the data in a PostgreSQL database.

## Tech Stack

| Component        | Technology          | Purpose                              |
| ---------------- | ------------------- | ------------------------------------ |
| Backend API      | Go                  | Receive and process sensor data      |
| Sensor Simulator | Go                  | Generate simulated sensor readings   |
| Database         | PostgreSQL          | Store sensor data and alerts         |
| Dashboard        | Next.js, TypeScript | Display readings and alerts          |
| Containerization | Docker              | Run the application and its services |


## Monitored Data

EcoWatch currently simulates three types of environmental sensor data:

- **Temperature** (°C)
- **Pressure** (bar)
- **Gas Concentration** (ppm)

The simulator generates normal readings along with occasional abnormal values
to test the monitoring and alerting system.


## Repository Structure

```text
ecowatch/
├── apps/
│   ├── api/              # Go backend API
│   ├── simulator/        # Go sensor simulator
│   └── dashboard/        # Next.js dashboard
│
├── db/
│   ├── migrations/       # Database migrations
│   └── seeds/            # Initial development data
│
├── docs/                 # Project documentation
├── deployments/          # Deployment configuration
├── .gitignore
├── README.md
└── docker-compose.yml