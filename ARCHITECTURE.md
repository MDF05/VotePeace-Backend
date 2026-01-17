# System Architecture

## Overview

VotePeace-Backend follows a **Modular Monolith** architecture pattern using the **Model-View-Controller (MVC)** design principle (adapted for API-first design where "View" is the JSON response).

The system is structured to separate concerns, ensuring that business logic, data access, and HTTP handling are distinct layers.

## High-Level Architecture

```mermaid
graph TD
    Client[Frontend Client] -->|HTTP Request| Router[Router (Fiber)]
    Router -->|Route Handler| Controller[Controller Layer]
    Controller -->|Business Logic| Service[Service Logic \n(Embedded in Controller)]
    Service -->|Data Mapping| Model[Data Models (GORM)]
    Model -->|SQL Query| DB[(SQLite Database)]
    
    subgraph "VotePeace Backend"
    Router
    Controller
    Model
    end
```

## Directory Structure

```
VotePeace-Backend/
├── controllers/    # Request handlers and business logic
├── database/       # Database connection and seeding logic
├── models/         # GORM Struct definitions (Database Schema)
├── routes/         # API Route definitions
├── main.go         # Application entry point
├── go.mod          # Go module definitions
└── votepeace.db    # SQLite Database file
```

## Core Components

### 1. HTTP Interface (Routes)
Located in `/routes`, this layer defines the API endpoints and maps them to specific controller functions. It serves as the entry point for all external requests.

### 2. Controller Layer
Located in `/controllers`, this layer handles:
- Parsing request body and parameters.
- Validating input.
- executing core business logic (e.g., calculating votes, registering users).
- Returning formatted JSON responses.

### 3. Data Layer (Models)
Located in `/models`, this layer defines the data structures (Structs) that map to database tables. We use **GORM** to abstract SQL queries and handle relationships (Associations).

### 4. Database
We currently use **SQLite** for simplicity and portability. The database connection logic is centralized in `/database/connect.go`, which also handles:
- **Auto-Migration**: Automatically creating tables based on Structs.
- **Seeding**: Populating initial Admin and Campaign data.

## Data Flow

1.  **Request**: User sends `POST /login` with credentials.
2.  **Route**: Fiber router matches `/login` and calls `controllers.Login`.
3.  **Controller**:
    *   Parses JSON body.
    *   Queries `models.User` via GORM to find the user.
    *   Compares password hash (bcrypt).
    *   Generates a Session/Token (if applicable) or success message.
4.  **Response**: Controller returns `200 OK` JSON with user details.

## Design Decisions

-   **Fiber Framework**: Chosen for its high performance and resemblance to Express.js, making it easier for Node.js developers to adapt.
-   **SQLite**: Chosen for zero-configuration, self-contained deployment suitable for this scale. (Can be easily swapped for PostgreSQL via GORM driver change).
-   **Direct Controller Logic**: For the current scale, business logic sits within controllers to avoid over-engineering with a separate Service layer. As the app grows, logic will be extracted to a `/services` package.
