package scheduler

import (
	"fmt"
	"sort"
	"sync"
)

// NodeCapacity captures allocatable resources for candidate host
type NodeCapacity struct {
	Name            string
	TotalVRAMGB     float64
	AllocatedVRAMGB float64
	TotalGPUs       int
	AllocatedGPUs   int
}

// RemainingVRAM returns unallocated VRAM in GB
func (n *NodeCapacity) RemainingVRAM() float64 {
	return n.TotalVRAMGB - n.AllocatedVRAMGB
}

// WorkloadRequest defines incoming container demand
type WorkloadRequest struct {
	PodName      string
	RequiredVRAM float64
	RequiredGPUs int
}

// BestFitScheduler executes Best-Fit Decreasing packing
type BestFitScheduler struct {
	mu sync.RWMutex
}

// NewBestFitScheduler initializes a new scheduler instance
func NewBestFitScheduler() *BestFitScheduler {
	return &BestFitScheduler{}
}

// ScoreNode evaluates tightest fit placement
func (s *BestFitScheduler) ScoreNode(w WorkloadRequest, n NodeCapacity) (int64, error) {
	if n.RemainingVRAM() < w.RequiredVRAM || (n.TotalGPUs-n.AllocatedGPUs) < w.RequiredGPUs {
		return 0, fmt.Errorf("insufficient capacity")
	}

	remAfter := n.RemainingVRAM() - w.RequiredVRAM
	score := int64(100.0 * (1.0 - (remAfter / n.TotalVRAMGB)))
	if score < 1 {
		score = 1
	}
	if score > 100 {
		score = 100
	}
	return score, nil
}

// SelectNode finds the highest scoring node
func (s *BestFitScheduler) SelectNode(w WorkloadRequest, nodes []NodeCapacity) (*NodeCapacity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	type scored struct {
		node  *NodeCapacity
		score int64
	}

	candidates := make([]scored, 0)
	for i := range nodes {
		sc, err := s.ScoreNode(w, nodes[i])
		if err == nil {
			candidates = append(candidates, scored{node: &nodes[i], score: sc})
		}
	}

	if len(candidates) == 0 {
		return nil, fmt.Errorf("no candidate nodes satisfy resource constraints")
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	return candidates[0].node, nil
}
