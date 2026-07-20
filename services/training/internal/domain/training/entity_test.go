package training

import (
	"testing"
)

func TestTrainingValidate(t *testing.T) {
	tests := []struct {
		name    string
		tr      Training
		wantErr bool
	}{
		{
			name: "Valid Training",
			tr: Training{
				Title:        "Confined Space Safety Training",
				CourseID:     "course-123",
				DepartmentID: "dept-456",
			},
			wantErr: false,
		},
		{
			name: "Missing Title",
			tr: Training{
				CourseID:     "course-123",
				DepartmentID: "dept-456",
			},
			wantErr: true,
		},
		{
			name: "Title Too Long",
			tr: Training{
				Title:        string(make([]byte, 201)),
				CourseID:     "course-123",
				DepartmentID: "dept-456",
			},
			wantErr: true,
		},
		{
			name: "Missing CourseID",
			tr: Training{
				Title:        "Confined Space Safety Training",
				DepartmentID: "dept-456",
			},
			wantErr: true,
		},
		{
			name: "Missing DepartmentID",
			tr: Training{
				Title:    "Confined Space Safety Training",
				CourseID: "course-123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tr.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Training.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
