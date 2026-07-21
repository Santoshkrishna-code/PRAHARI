# PRAHARI Platform: Emergency & Operational Runbooks

This document provides step-by-step resolution pathways for critical operational incidents on the PRAHARI Platform.

---

## Runbook 01: Production Outage (P1 Incident)
### Diagnosis
1. Global ingress alert triggered (HTTP 5xx rate > 5%).
2. Open the AWS CloudWatch Dashboard and verify the Application Load Balancer (ALB) health checks status.

### Actions
1. **Verify DNS Routing**: Check if Route 53 is successfully routing traffic to the primary ALB.
2. **Review Gateway Logs**: Connect to the EKS cluster and check Gateway service logs for routing failures:
   ```bash
   kubectl logs -n prahari-core -l app=gateway-service --tail=200
   ```
3. **Isolate Component**: If the failure is localized to a single service, scale the target Deployment to zero and re-scale to force pod regeneration.

---

## Runbook 02: Database Connection Failure (PostgreSQL)
### Diagnosis
- System alerts show `Database connection pool exhausted` or `Dial connection timeout` events.

### Actions
1. **Check Aurora Status**: Check the AWS Aurora PostgreSQL console. Verify if the cluster is performing a failover or minor version upgrade.
2. **Review Active Connections**: Run a query on PostgreSQL to identify and terminate idle sessions blocking the pool:
   ```sql
   SELECT pg_terminate_backend(pid) FROM pg_stat_activity 
   WHERE state = 'idle' AND state_change < now() - interval '5 minutes';
   ```
3. **Verify Security Groups**: Confirm VPC Security Groups permit TCP port 5432 communication between EKS nodes and RDS subnets.

---

## Runbook 03: Kafka Message Broker Failure
### Diagnosis
- Event processing lag metrics (`prahari_kafka_consumer_lag`) rising steadily.

### Actions
1. **Analyze Broker Health**: Check the MSK console to verify broker resource metrics. Check if broker CPU or memory utilization is exceeding 90%.
2. **Rebalance Partitions**: If a broker has crashed and restarted, execute the Kafka partition rebalancer tool to redistribute partitions across active brokers:
   ```bash
   kafka-reassign-partitions.sh --bootstrap-server $KAFKA_BROKERS --reassignment-json-file reassign.json --execute
   ```
3. **Inspect Consumer Group status**: Check if consumer pods are in a crash loop, preventing offset commits.

---

## Runbook 04: Redis Cache Failure
### Diagnosis
- Latency spikes on GET requests; database CPU load rising due to cache misses.

### Actions
1. **Check Replication Status**: Log in to AWS ElastiCache and verify primary-replica sync status.
2. **Eviction Audit**: Check if Redis has run out of memory. If memory usage is at 100%, update the eviction policy to `allkeys-lru` or increase the instance type size.
3. **Bypass Cache**: If Redis is completely offline, configure the application to fallback to direct database queries temporarily until the cluster is restored.

---

## Runbook 05: Kubernetes Cluster / Node Failure
### Diagnosis
- Pods remain in `Pending` state; nodes report status `NotReady`.

### Actions
1. **Check Karpenter Logs**: Review Karpenter controller logs to identify node provisioning failures:
   ```bash
   kubectl logs -n karpenter -l app.kubernetes.io/name=karpenter --tail=100
   ```
2. **Inspect EKS Node Events**: Run diagnostics on the failing nodes:
   ```bash
   kubectl describe node <node-name>
   ```
3. **Re-provision Nodes**: Force termination of the failing EC2 node; Karpenter will provision a fresh instance.

---

## Runbook 06: Deployment Failure
### Diagnosis
- ArgoCD sync failure; pods stuck in `CrashLoopBackOff` or `ImagePullBackOff`.

### Actions
1. **Inspect Pod Describe**: Identify the cause of the container crash:
   ```bash
   kubectl describe pod -n prahari-core <pod-name>
   ```
2. **Verify Credentials**: Check if the container registry configuration has expired, resulting in `ImagePullBackOff`.
3. **Rollback**: Rollback immediately by triggering an abort on the active Argo Rollout (see Runbook 11).

---

## Runbook 07: Security Incident (Credential / Token Leak)
### Diagnosis
- Secrets leak detected in repository commits, or abnormal API requests from unrecognized external IPs.

### Actions
1. **Revoke Exposed Tokens**: If JWT signature verification keys are leaked, rotate the JWKS key ring inside AWS Cognito.
2. **Rotate Secrets**: Force rotate the affected database passwords or API keys in Secrets Manager.
3. **Invalidate Sessions**: Force close active user sessions in Cognito to force re-authentication across all clients.

---

## Runbook 08: AI Agent Service Failure (Timeout / Model Errors)
### Diagnosis
- Users report `AI Assistant Unavailable` errors; application traces show API timeouts on LLM endpoints.

### Actions
1. **Check LLM Providers**: Check status pages of the external LLM APIs (OpenAI, Anthropic) for outages.
2. **Verify Bedrock Standby**: Check if fallback logic has successfully routed requests to Amazon Bedrock.
3. **Clear Vector Indexes**: If search queries fail, verify if the pgvector index has become corrupt. Re-index if needed:
   ```sql
   REINDEX INDEX idx_sds_embeddings_hnsw;
   ```

---

## Runbook 09: Computer Vision Processing Failure (RTSP Stream Drop)
### Diagnosis
- CV alerts cease; dashboard shows `Camera Stream Disconnected` statuses.

### Actions
1. **Verify Ping**: Ping the IP cameras from the local Edge Gateway.
2. **Check GStreamer / DeepStream Pipeline**: Restart the local frame processing container to re-initialize the RTSP buffer:
   ```bash
   docker restart prahari-cv-inference
   ```
3. **Inspect Frame Drop logs**: Verify if network congestion is causing packet drop rates > 15%, causing decoding pipeline crashes.

---

## Runbook 10: Digital Twin Sync Failure
### Diagnosis
- Physical plant state does not reflect in 3D visualization; synchronization lag metrics rising.

### Actions
1. **Check WebSocket Connection**: Inspect logs on the gateway service to verify if the client-to-twin WebSocket connection is dropping.
2. **Check Time Sync**: Ensure system clocks on edge cameras, IoT sensors, and cloud systems are synchronized via NTP. Time drift > 1s causes event dropouts.
3. **Purge Cache**: Reset the twin cache key inside Redis to force a complete database-to-twin synchronization cycle.

---

## Runbook 11: Rollback Procedure (Argo Rollouts)
### Diagnosis
- Canary deployment fails telemetry validation; error rates rise.

### Actions
1. **Command Rollback**: Run the CLI rollback command to immediately terminate the deployment rollout:
   ```bash
   kubectl argo rollouts undo -n prahari-core deployment/chemical-service
   ```
2. **Verify Sync**: Confirm that ArgoCD has returned the deployment state to the previous stable git commit tag.

---

## Runbook 12: Disaster Recovery (DR) Region Promotion
### Diagnosis
- Primary AWS region (`ap-south-1`) suffers an outage.

### Actions
1. **Promote DB**: Log in to AWS Aurora in the secondary region (`ap-northeast-1`) and promote the reader cluster to active writer status.
2. **Update DNS**: Update Route 53 DNS routing parameters to divert all traffic to the secondary region ALB.
3. **Sync EKS Workloads**: Instruct ArgoCD to synchronize application manifests to the EKS cluster in the secondary region.
