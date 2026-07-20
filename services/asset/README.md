# Asset Management Service

The **PRAHARI Asset Management Service** (`services/asset/`) is the master system of record owning all industrial assets metadata profiles, hierarchies structures, calculated criticality scoring, and condition index predictions logs.

## Features

- **Hierarchical Asset Tree:** Model parent-child structural equipment trees to trace components.
- **Computed Criticality Rating:** Formulate priority metrics based on safety, business, environmental and production impact scores.
- **State Machine Rules:** Validate status transitions (`Registered` → `Operational` → `Maintenance`).
- **Scan tags:** Support QR/Barcodes identifier metadata tagging.
- **Policies Checks:** Preconditions checking warranty limits, Mutability restrictions, and safety review triggers.

## Tech Stack

- **Core:** Go 1.21+
- **Architecture:** Clean Hexagonal domain-driven layout
- **Database:** PostgreSQL (hierarchical recursive queries & dynamic spec columns)
- **Messaging:** Kafka Event-Driven integration pipelines
- **Caching:** Redis Cache
