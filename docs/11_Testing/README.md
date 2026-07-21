# PRAHARI Platform: Testing Strategy and Framework

This document outlines the testing methodologies, tools, and validation rules used to verify the reliability, security, and performance of the PRAHARI Platform.

---

## 1. Multi-Tier Testing Pyramid
PRAHARI uses a comprehensive test suite to identify issues early, structured as follows:

```
                  ┌──────────────────────┐
                  │  Acceptance / E2E    │  <-- Cypress & Playwright (UI flows)
                  ├──────────────────────┤
                  │   Chaos & Load       │  <-- Chaos Mesh & k6 performance
                  ├──────────────────────┤
                  │ Integration/Contract │  <-- Testcontainers & Pact
                  ├──────────────────────┤
                  │     Unit Tests       │  <-- Go testing, stretchr/testify
                  └──────────────────────┘
```

---

## 2. Unit Testing Conventions
- **Framework**: Standard Go `testing` library with **testify** assertions.
- **Isolation**: Domain layers must be tested in isolation from network adapters or databases. Use **mockgen** to generate mock implementations for repositories and external clients.
- **Coverage Target**: Minimum of **80% code coverage** enforced at the pull request level.

---

## 3. Integration Testing (Testcontainers)
To ensure database operations and event streaming function correctly under real-world conditions, integration tests run against ephemeral infrastructure spun up in Docker containers via **Testcontainers Go**.

### Example: Kafka Integration Test Setup
```go
import (
    "context"
    "testing"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/kafka"
)

func TestKafkaPublisher(t *testing.T) {
    ctx := context.Background()
    kafkaContainer, err := kafka.RunContainer(ctx, testcontainers.WithImage("confluentinc/cp-kafka:7.4.0"))
    if err != nil {
        t.Fatalf("failed to start kafka: %s", err)
    }
    defer kafkaContainer.Terminate(ctx)
    
    brokers, _ := kafkaContainer.Brokers(ctx)
    // Initialize application publisher passing the test broker coordinates...
}
```

---

## 4. Contract & Performance Testing

### 4.1 API Contract Testing (Pact)
Contract testing via **Pact Go** ensures API contracts between the frontend (Consumer) and backend microservices (Provider) remain compatible. The Pact broker registers contract specs and verifies them during release pipelines to prevent breaking changes.

### 4.2 Performance & Load Testing (k6)
Load tests are executed using **k6** scripts to evaluate performance under peak conditions.
- **Throughput Target**: 10,000 requests/sec for API gateway.
- **Latency Target**: P99 response time < 250ms under peak load.
- **Sample k6 Script Configuration**:
  ```javascript
  export const options = {
    stages: [
      { duration: '2m', target: 500 },  // Ramp-up to 500 concurrent virtual users
      { duration: '10m', target: 500 }, // Hold load for 10 minutes
      { duration: '2m', target: 0 },    // Ramp-down
    ],
    thresholds: {
      http_req_duration: ['p(99)<250'],  // 99% of requests must complete below 250ms
    },
  };
  ```

---

## 5. Chaos Engineering (Chaos Mesh)
We use **Chaos Mesh** inside the EKS staging namespace to verify fault tolerance:
- **Pod Chaos**: Randomly terminates microservice pods (e.g. Chemical Service) to verify Kubernetes replicas recover without dropping client connections.
- **Network Chaos**: Injects 200ms latency and 2% packet loss between the microservices and MSK cluster to verify the Outbox Pattern retry queue functions correctly without losing events.

---

## 6. AI and Computer Vision Model Validation

### 6.1 AI Agent Evaluation (Ragas & Guardrails)
- **RAG Assessment**: Evaluate retrieval precision and generation faithfulness using the **Ragas** framework. A baseline score of > 0.85 is required for production.
- **Jailbreak Testing**: Simulated adversary agents attempt to trigger system prompt exposure or unauthorized safety rule overrides (tested daily).

### 6.2 Computer Vision Validation (mAP & Offline Buffers)
- **Model Metric**: Models must achieve a Mean Average Precision (mAP@0.5:0.95) score of **> 0.78** on our validation dataset before deployment to edge gateways.
- **Edge Buffering Test**: Edge gateways undergo physical network disconnect simulation to verify that video processing remains active and safety alerts queue locally in SQLite, syncing once connectivity is restored.
