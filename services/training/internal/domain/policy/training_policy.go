package policy

import (
	"errors"

	"prahari/services/training/internal/domain/training"
)

// CheckAttendanceRequirements checks if course requires attendance verifications.
func CheckAttendanceRequirements(t *training.Training, attendanceCount int) error {
	if t.StatusCode == "ASSESSMENT" && attendanceCount == 0 {
		return errors.New("course assessment review requires verified attendance records")
	}
	return nil
}
