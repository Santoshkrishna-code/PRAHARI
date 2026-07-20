package policy

import (
	"testing"
	"time"

	"prahari/services/shift/internal/domain/checklist"
)

func TestCanInitiateHandover(t *testing.T) {
	itemsAllCompleted := []checklist.Item{
		{ID: "c1", Completed: true},
		{ID: "c2", Completed: true},
	}
	itemsSomeIncomplete := []checklist.Item{
		{ID: "c1", Completed: true},
		{ID: "c2", Completed: false},
	}

	if !CanInitiateHandover(itemsAllCompleted) {
		t.Errorf("Expected CanInitiateHandover to return true when all items completed")
	}
	if CanInitiateHandover(itemsSomeIncomplete) {
		t.Errorf("Expected CanInitiateHandover to return false when some items incomplete")
	}
}

func TestCalculateShiftDuration(t *testing.T) {
	start := time.Now()
	end := start.Add(10 * time.Hour)

	duration := CalculateShiftDuration(start, end)
	if duration != 10.0 {
		t.Errorf("Expected duration to be 10.0, got %.2f", duration)
	}
}

func TestIsOvertimeRequired(t *testing.T) {
	if !IsOvertimeRequired(9.5) {
		t.Errorf("Expected IsOvertimeRequired to return true for 9.5 hours")
	}
	if IsOvertimeRequired(7.5) {
		t.Errorf("Expected IsOvertimeRequired to return false for 7.5 hours")
	}
}
