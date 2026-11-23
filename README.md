# 📦 PackerOps Community Edition
### Kubernetes GPU Bin-Packing Scheduler & Quiescent Idle Power Audit CLI

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-v1.28%2B-326ce5.svg)](https://kubernetes.io/)
[![Go Report Card](https://img.shields.io/badge/Go-1.11%2B-00ADD8.svg)](https://golang.org)
[![Website: packeropsai.com](https://img.shields.io/badge/Official_Website-packeropsai.com-0d9488.svg)](https://packeropsai.com)

**PackerOps Community Edition** is a vendor-neutral, cloud-native Kubernetes scheduling plugin and CLI audit tool designed to detect **quiescent idle electrical dissipation (400W–750W per node)** in high-density accelerator clusters (NVIDIA B200, H100, A100) and provide deterministic **Best-Fit Decreasing (BFD)** workload bin-packing.

---

## 🔍 The Quiescent Idle Power Paradox in GPU Clouds

In multi-tenant AI clusters, modern accelerator servers exhibit significant base electrical leakage even when idle. Due to High-Bandwidth Memory (HBM3/HBM3e) refresh cycles, PCIe Gen5/NVLink interconnect PHY leakage, and chassis cooling, a server with four NVIDIA H100 or B200 GPUs continuously draws **400W to 750W in quiescent base power**.

PackerOps provides:
1. **Quiescent Leakage Audit CLI (`packerops-audit`):** Quickly quantify monthly kilowatt-hour (kWh) and financial losses caused by idling accelerator nodes.
2. **Best-Fit Decreasing (BFD) Scheduler Plugin:** Replaces Kubernetes' default spread-out scheduling (`LeastRequestedPriority`) with deterministic bin-packing, maximizing allocation density and opening opportunities for physical node scale-to-zero.

---

## ⚡ Quick Start: 30-Second Cluster Waste Audit

Run the standalone zero-privilege audit tool to evaluate potential idle dissipation across your cluster:

```bash
# Clone the community repository
git clone https://github.com/packeropsai/packeropsai.git
cd packeropsai

# Execute the quiescent leakage audit
go run cmd/packerops-audit/main.go --cost-per-kwh=0.216
```

### Sample Audit Output:
```text
========================================================================================
  PackerOps Community Edition - GPU Cluster Quiescent Leakage Audit
========================================================================================
Node ID          GPU Model  GPUs       Idle Hours     Idle Power (W)  
----------------------------------------------------------------------------------------
gpu-node-01      B200       4          8.5            720.0           
gpu-node-02      B200       4          6.2            720.0           
gpu-node-03      H100       4          12.0           580.0           
gpu-node-04      H100       4          9.4            580.0           
gpu-node-05      A100       2          5.1            340.0           
----------------------------------------------------------------------------------------
TOTAL QUIESCENT DISSIPATION: 2.94 kW
MONTHLY ENERGY WASTE:        2,116.8 kWh
ESTIMATED MONTHLY FINANCIAL LOSS: $457.23 (@ $0.216/kWh)
ESTIMATED ANNUAL FINANCIAL LOSS:  $5,486.75
========================================================================================
```

---

## 🛠️ Deploying to Kubernetes (2 Quick Methods)

### Method A: One-Command Direct Install (Fastest, Zero Clone Required)
Deploy the PackerOps scheduler and RBAC controller into your cluster in 5 seconds:

```bash
kubectl apply -f https://raw.githubusercontent.com/packeropsai/packeropsai/main/deploy/manifests/install.yaml
```

Verify the scheduler pod is healthy:
```bash
kubectl get pods -n packerops-system
```

---

### Method B: Customizable Helm 3+ Installation
If you prefer managing configurations (replica counts, resource limits, leader election) with Helm:

```bash
# 1. Clone the repository
git clone https://github.com/packeropsai/packeropsai.git
cd packeropsai

# 2. Deploy using Helm into the 'packerops-system' namespace
helm install packerops ./deploy/helm/packerops-community \
  --namespace packerops-system \
  --create-namespace \
  --set scheduler.leaderElect=true
```

### Schedule Workloads via PackerOps:
Assign any Pod or Deployment to use the `packerops-scheduler`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: llm-inference-worker
  namespace: default
spec:
  replicas: 4
  template:
    spec:
      schedulerName: packerops-scheduler
      containers:
      - name: vllm-worker
        image: vllm/vllm-openai:latest
        resources:
          limits:
            nvidia.com/gpu: 2
            memory: "64Gi"
```

---

## 📊 Feature Matrix: Community vs. Enterprise

| Capability | PackerOps Community (This Repo) | PackerOpsAI Enterprise (Patent Protected) |
| :--- | :---: | :---: |
| **Core Algorithm** | Classical Best-Fit Decreasing (BFD) | **Two-Tier MUBGA Coupled Genetic Engine** |
| **Idle Power Detection** | Standalone CLI Audit | **Real-Time DCGM + Redfish Telemetry** |
| **Model Migration Control** | Kubernetes Eviction (Reactive) | **Lemma 1 Migration Break-Even Engine** |
| **Migration Thrashing Guard** | None | **Zero-Migration Stability Anchor (98% Churn Reduction)** |
| **Hardware Sleep States** | Manual Cordon / Drain | **Automated IPMI/Redfish ACPI S3 (<= 15W BMC Standby)** |
| **Roofline Batch Pacing** | Static Batch Size | **Dynamic Token SGR-LS (B in {16, 32, 64})** |
| **Substation Power Headroom** | None | **Dynamic Shaving (+25% More GPUs per Substation)** |

---

## 🌐 Enterprise Inquiries & Contact

For large-scale hyperscale or Neocloud deployments requiring the patent-protected **Two-Tier MUBGA Engine**, **Lemma 1 Migration Break-Even Engine**, and automated **ACPI S3 hardware sleep orchestration**, visit:

🔗 **[https://packeropsai.com](https://packeropsai.com)**

*(Please use the official contact form on the website for enterprise licensing, proof-of-concept cluster pilots, and partnership inquiries.)*

---

## 📄 License

This community edition is licensed under the **Apache License, Version 2.0**. See [LICENSE](LICENSE) for details.
