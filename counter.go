package main

import (
	sm "dashboard-sanitizer/model"
	"fmt"
	"github.com/clarketm/json"
)

type Stats struct {
	Count   int            `json:"count"`
	Details map[string]int `json:"details"`
}

type StatsCounter struct {
	Total     int   `json:"total"`
	Processed int   `json:"processed"`
	Skipped   Stats `json:"skipped"`
}

func (c *StatsCounter) RegisterSkipped(do sm.DashboardObject) {
	if do.Type == "" {
		b, _ := json.Marshal(do)
		fmt.Printf("%s", b)
		return
	}
	c.Skipped.Count++
	objectTypeCount := c.Skipped.Details[do.Type]
	if objectTypeCount == 0 {
		c.Skipped.Details[do.Type] = 1
	} else {
		c.Skipped.Details[do.Type] = objectTypeCount + 1
	}
	c.Total++
}
func (c *StatsCounter) RegisterProcessed(do sm.DashboardObject) {
	c.Processed++
	c.Total++
}

func NewStatusCount() *StatsCounter {
	return &StatsCounter{
		Skipped: Stats{
			Count:   0,
			Details: map[string]int{},
		},
	}
}

func (c *StatsCounter) PrintStats() {
	by, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s", by)
}
