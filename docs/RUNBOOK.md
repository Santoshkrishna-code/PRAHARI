# PRAHARI AWS SRE Operational Runbook

This runbook provides site reliability engineering procedures for operating and troubleshooting the **PRAHARI Monorepo & AI Platform** on AWS.

---

## 1. Routine Health & Observability Verification

### Check ECS Task Execution Status
```bash
aws ecs list-tasks --cluster prahari-hackathon-cluster
```

### Inspect Container Logs via CloudWatch
```bash
aws logs tail /ecs/prahari-hackathon --follow
```

### Verify Application Load Balancer Health
```bash
curl -I http://<ALB_DNS_NAME>/health
```

---

## 2. Emergency Incident Response Procedures

### Scenario A: High Latency or Timeout on AI Platform Task
1. Check CloudWatch CPU / Memory metrics:
   ```bash
   aws cloudwatch get-metric-statistics --namespace AWS/ECS --metric-name CPUUtilization ...
   ```
2. Restart ECS Service Tasks:
   ```bash
   aws ecs update-service --cluster prahari-hackathon-cluster --service prahari-ai-platform-service --force-new-deployment
   ```

### Scenario B: Database Connection Pool Exhaustion
1. Restart active PostgreSQL sessions or check RDS metrics:
   ```bash
   aws rds describe-db-instances --db-instance-identifier prahari-postgres-hackathon
   ```

### Scenario C: Flushing ElastiCache Redis Session Cache
```bash
redis-cli -h <REDIS_ENDPOINT> -p 6379 FLUSHALL
```
