# Operational Runbooks: Troubleshooting & Diagnostics

This guide outlines common operational failures in production environments and provides diagnostic resolution instructions.

---

## 1. Circuit Breaker Trips (Sustained Outages)

### Symptom
- Request logs show `circuitbreaker: request blocked because circuit is open` errors.

### Resolution
1. Inspect connection logs to check if target API or database servers respond.
2. Verify downstream latency metrics.
3. Check the configured circuit breaker cooldown interval to see when the circuit transitions to HALF-OPEN for automated re-tests.

---

## 2. Kafka Consumer Lag Spikes

### Symptom
- Metrics indicate rising lag on partition keys.

### Resolution
1. Verify partition leader balances.
2. Check if the read-commit loops log callback errors, preventing offset updates.
