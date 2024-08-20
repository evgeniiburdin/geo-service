# 🌍 Geo-Service

Geo-Service is a microservices-based project designed to handle user data, geocoding services, and proxying requests between different services. It leverages Go, Docker, and Prometheus for efficient operation and monitoring.

## ✨ Features
- **Auth Service**: Manages user authorization (JWT + Postgres)
- **User Service**: Manages user data storage (Redis + Postgres)
- **Geo Service**: Handles geocoding requests (OpenCage API)
- **Proxy Service**: Routes requests between different services (gRPC)

## 🚀 Getting Started

### Prerequisites
- Docker
- Go (if running locally)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/evgeniiburdin/geo-service.git
   cd geo-service
   
2. Prepare the database:
    ```bash
   ./db_prepare.sh

3. Start the services using Docker Compose:
    ```bash
   docker-compose up --build