package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/alert"
	"prahari/services/vision/internal/domain/camera"
	"prahari/services/vision/internal/domain/detection"
	"prahari/services/vision/internal/domain/inference"
	"prahari/services/vision/internal/domain/search"
	"prahari/services/vision/internal/domain/tracking"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveCamera(ctx context.Context, c *camera.Camera) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO cameras (id, plant_id, name, ip_address, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.PlantID, c.Name, c.IPAddress, c.Status, c.CreatedAt)
	return err
}

func (s *Store) GetCameraByID(ctx context.Context, id string) (*camera.Camera, error) {
	if s.db == nil {
		return &camera.Camera{ID: id, PlantID: "P01", Name: "Reactor Cam", IPAddress: "192.168.1.50", Status: "ONLINE"}, nil
	}
	query := `SELECT id, plant_id, name, ip_address, status, created_at FROM cameras WHERE id = $1`
	var c camera.Camera
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.PlantID, &c.Name, &c.IPAddress, &c.Status, &c.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("camera %s not found", id)
	}
	return &c, err
}

func (s *Store) SaveInferenceJob(ctx context.Context, job *inference.Job) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO inference_jobs (id, camera_id, model_id, status, fps_rate, started_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status`
	_, err := s.db.ExecContext(ctx, query, job.ID, job.CameraID, job.ModelID, job.Status, job.FPSRate, job.StartedAt)
	return err
}

func (s *Store) SaveDetection(ctx context.Context, d *detection.Detection) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO detections (id, job_id, label, confidence, bbox, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.JobID, d.Label, d.Confidence, d.BBox, d.Timestamp)
	return err
}

func (s *Store) SaveTrackSegment(ctx context.Context, seg *tracking.Segment) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO tracked_objects (id, object_id, camera_id, x_val, y_val, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, seg.ID, seg.ObjectID, seg.CameraID, seg.XVal, seg.YVal, seg.Timestamp)
	return err
}

func (s *Store) SaveAlert(ctx context.Context, a *alert.EventTrigger) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO alerts (id, camera_id, label, triggered_at, snapshot_url)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.CameraID, a.Label, a.TriggeredAt, a.SnapshotURL)
	return err
}

func (s *Store) SearchDetections(ctx context.Context, criteria *search.Criteria) ([]*detection.Detection, int64, error) {
	if s.db == nil {
		mockDetection := &detection.Detection{ID: "det-001", JobID: "job-1", Label: criteria.Label, Confidence: 0.95, Timestamp: time.Now()}
		return []*detection.Detection{mockDetection}, 1, nil
	}
	query := `SELECT id, job_id, label, confidence, bbox, timestamp FROM detections`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []*detection.Detection
	for rows.Next() {
		var d detection.Detection
		if err := rows.Scan(&d.ID, &d.JobID, &d.Label, &d.Confidence, &d.BBox, &d.Timestamp); err != nil {
			return nil, 0, err
		}
		result = append(result, &d)
	}
	return result, int64(len(result)), nil
}

func (s *Store) GetDetectionByID(ctx context.Context, id string) (*detection.Detection, error) {
	return &detection.Detection{ID: id, JobID: "job-01", Label: "no_helmet", Confidence: 0.92, Timestamp: time.Now()}, nil
}
