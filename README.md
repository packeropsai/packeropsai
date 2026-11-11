# PackerOps Community Edition
### Open-Source Best-Fit Kubernetes GPU Bin-Packing Scheduler & Idle Audit Tool

PackerOps Community Edition is an open-source Kubernetes Scheduling Framework plugin and cluster audit CLI designed to detect quiescent electrical dissipation in accelerated datacenters and provide deterministic Best-Fit Decreasing (BFD) workload bin-packing.

---

## ⚡ Quick Start: Audit Cluster Idle Leakage

```bash
# Run standalone audit CLI
go run cmd/packerops-audit/main.go --cost-per-kwh=0.216
```

## 📦 Features (Community vs. Enterprise)

| Feature | PackerOps Community | PackerOpsAI Enterprise |
| :--- | :---: | :---: |
| **Algorithm** | Best-Fit Decreasing (BFD) | Two-Tier MUBGA Coupled Genetic Engine |
| **Idle Leakage Audit** | CLI Report | Real-time DCGM + BMC Telemetry |
| **Migration Management**| Reactive | Lemma 1 Break-Even Engine |
| **Hardware Sleep** | Manual | Automated Out-of-band ACPI S3 (<= 15W) |
| **SLA Latency Guard** | Basic | Roofline P90 ITL / TTFT Pacing |

For enterprise inquiries, visit [packeropsai.com](https://packeropsai.com).

## 📄 License
Licensed under the Apache License, Version 2.0.
