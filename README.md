# PRAHARI: Enterprise EHS Platform

PRAHARI is a production-ready, enterprise-grade Environment, Health, and Safety (EHS) platform. Designed for high-risk industrial environments (manufacturing, chemical plants, utilities), it fuses a microservices transactional backend with AI assistance, real-time Computer Vision surveillance, and Digital Twin layouts simulation workspaces.

---

## Key Capabilities

1. **44 Microservices Backend Architecture**: Modular Go core handling lockouts (LOTO), dynamic safety permits, audits, environmental incidents tracking, and work authorizations.
2. **Shared UI Platform**: Consistent Design Tokens and accessible custom Tailwind + Radix UI inputs mapping.
3. **React Enterprise Web Portal**: Central console providing operational views, permits workflows, visual camera streams, and twin topology controls.
4. **Android & iOS Field Operations Clients**: Offline-first mobile applications with local SQLite caching, biometric checks (Face ID), Core Location tags, and evidence camera uploads.
5. **Computer Vision Perception Engine**: RTSP/WebRTC streams with pluggable AI models evaluating restricted zone boundaries.
6. **Digital Twin Visual Plane**: Real-time sensor state integrations and predictive what-if pressure simulations.

---

## Repository Folder Structure

```text
PRAHARI/
├── services/               # 44 Backend Go microservices
├── shared/                 # Shared Go packages
├── frontend/
│   ├── platform/           # UI Design Tokens & Component Library
│   └── web/                # React Enterprise Web Portal
├── mobile/
│   ├── android/            # React Native Android client
│   └── ios/                # React Native iOS client
└── deployments/            # Terraform, Helm, Argo CD & Observability templates
```

---

## Installation & Local Development

### 1. Backend Microservices
Run the services using the built-in Go compile toolchain:
```bash
GOWORK=off go build ./...
GOWORK=off go test -v -race ./...
```

### 2. Frontend Web Portal
Launch Vite dev server for the React portal:
```bash
cd frontend/web
npm install
npm run dev
```

### 3. Mobile Applications
Compile mobile clients:
```bash
cd mobile/android
npm install --legacy-peer-deps
npm run build
```

### 4. Terraform Infrastructure Validation
Lint manifests targetting AWS EKS:
```bash
cd deployments/terraform
terraform init
terraform validate
```

---

## License

This project is licensed under the [MIT License](LICENSE).
