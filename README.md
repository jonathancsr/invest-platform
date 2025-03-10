# invest-platform

# Invest Platform

A microservices-based investment platform for processing trade orders.

## Overview

This platform processes trade orders through a Kafka-based event-driven architecture. It includes:

- Order processing service
- Real-time trade matching engine
- Event-based communication

## Architecture

The system uses Apache Kafka for message queuing and event streaming, with the following components:

- **Kafka**: Message broker for order processing
- **Zookeeper**: Coordination service for Kafka
- **Control Center**: Web UI for managing Kafka

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.16+

### Running the Application

1. Start the infrastructure:

   ```
   cd order-service
   docker-compose up -d
   ```

2. Run the trade service:
   ```
   go run cmd/trade/main.go
   ```
