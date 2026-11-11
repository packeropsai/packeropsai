package main

import (
	"flag"

	"github.com/packeropsai/packerops-community/pkg/audit"
)

func main() {
	costPerKWh := flag.Float64("cost-per-kwh", 0.216, "Electricity rate in USD/kWh")
	flag.Parse()

	nodes := audit.SampleSyntheticCluster()
	result := audit.CalculateClusterWaste(nodes, *costPerKWh)
	audit.PrintReport(nodes, result, *costPerKWh)
}
