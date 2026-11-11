package audit

import (
	"fmt"
	"time"
)

// GPUNodeSample represents basic hardware profile of a cluster host
type GPUNodeSample struct {
	NodeID       string  `json:"node_id"`
	GPUModel     string  `json:"gpu_model"`
	GPUCount     int     `json:"gpu_count"`
	IdleWatts    float64 `json:"idle_watts"`
	IdleDuration float64 `json:"idle_duration_hours"`
}

// AuditResult summarizes quiescent electrical waste
type AuditResult struct {
	TotalQuiescentKW float64 `json:"total_quiescent_kw"`
	MonthlyLossKWh   float64 `json:"monthly_loss_kwh"`
	MonthlyCostUSD   float64 `json:"monthly_cost_usd"`
	AnnualCostUSD    float64 `json:"annual_cost_usd"`
}

// CalculateClusterWaste evaluates electrical dissipation of unutilized accelerator hosts
func CalculateClusterWaste(nodes []GPUNodeSample, costPerKWh float64) AuditResult {
	var totalWatts float64
	for _, n := range nodes {
		totalWatts += n.IdleWatts
	}

	totalKW := totalWatts / 1000.0
	monthlyKWh := totalKW * 24.0 * 30.0
	monthlyUSD := monthlyKWh * costPerKWh
	annualUSD := monthlyUSD * 12.0

	return AuditResult{
		TotalQuiescentKW: totalKW,
		MonthlyLossKWh:   monthlyKWh,
		MonthlyCostUSD:   monthlyUSD,
		AnnualCostUSD:    annualUSD,
	}
}

// SampleSyntheticCluster returns representative sample nodes for demonstration
func SampleSyntheticCluster() []GPUNodeSample {
	return []GPUNodeSample{
		{NodeID: "gpu-node-01", GPUModel: "B200", GPUCount: 4, IdleWatts: 720.0, IdleDuration: 8.5},
		{NodeID: "gpu-node-02", GPUModel: "B200", GPUCount: 4, IdleWatts: 720.0, IdleDuration: 6.2},
		{NodeID: "gpu-node-03", GPUModel: "H100", GPUCount: 4, IdleWatts: 580.0, IdleDuration: 12.0},
		{NodeID: "gpu-node-04", GPUModel: "H100", GPUCount: 4, IdleWatts: 580.0, IdleDuration: 9.4},
		{NodeID: "gpu-node-05", GPUModel: "A100", GPUCount: 2, IdleWatts: 340.0, IdleDuration: 5.1},
	}
}

// PrintReport outputs audit summary to console
func PrintReport(nodes []GPUNodeSample, res AuditResult, costPerKWh float64) {
	fmt.Println("========================================================================================")
	fmt.Println("  PackerOps Community Edition - GPU Cluster Quiescent Leakage Audit")
	fmt.Printf("  Report Generated: %s\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Println("========================================================================================")
	fmt.Printf("%-16s %-10s %-10s %-14s %-16s\n", "Node ID", "GPU Model", "GPUs", "Idle Hours", "Idle Power (W)")
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, n := range nodes {
		fmt.Printf("%-16s %-10s %-10d %-14.1f %-16.1f\n", n.NodeID, n.GPUModel, n.GPUCount, n.IdleDuration, n.IdleWatts)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("TOTAL QUIESCENT DISSIPATION: %.2f kW\n", res.TotalQuiescentKW)
	fmt.Printf("MONTHLY ENERGY WASTE:        %.1f kWh\n", res.MonthlyLossKWh)
	fmt.Printf("ESTIMATED MONTHLY FINANCIAL LOSS: $%.2f (@ $%.3f/kWh)\n", res.MonthlyCostUSD, costPerKWh)
	fmt.Printf("ESTIMATED ANNUAL FINANCIAL LOSS:  $%.2f\n", res.AnnualCostUSD)
	fmt.Println("========================================================================================")
}
