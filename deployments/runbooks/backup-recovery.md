# Incident Response Runbook — PostgreSQL Backup & Disaster Recovery

## Incident Summary

Database connection dropouts or transactional corruption causing loss of service configurations.

---

## 1. On-Call Checklist

1. Verify target RDS DB instance status using AWS CLI:
   ```bash
   aws rds describe-db-instances --db-instance-identifier prahari-postgres-production
   ```
2. Check metrics graphs in Grafana monitoring panel: **Active Database Backends**.
3. Ping service health checks endpoints: `/health` and `/ready`.

---

## 2. Recovery Actions (Point-in-Time Recovery)

In the event of database corruption or deletion, restore data from AWS RDS snapshot copies:

1. Identify the target snapshot reference tag:
   ```bash
   aws rds describe-db-snapshots --db-instance-identifier prahari-postgres-production --query "DBSnapshots[*].[DBSnapshotIdentifier,SnapshotCreateTime]"
   ```
2. Trigger RDS instance restoration:
   ```bash
   aws rds restore-db-instance-from-db-snapshot \
     --db-instance-identifier prahari-postgres-restored \
     --db-snapshot-identifier prahari-postgres-snapshot-target \
     --db-subnet-group-name prahari-db-subnet-group-production
   ```
3. Update EKS service configurations to route database endpoints traffic to `prahari-postgres-restored`.
