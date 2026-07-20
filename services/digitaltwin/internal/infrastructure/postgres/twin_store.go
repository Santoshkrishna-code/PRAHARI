package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/digitaltwin/internal/domain/overlay"
	"prahari/services/digitaltwin/internal/domain/playback"
	"prahari/services/digitaltwin/internal/domain/search"
	"prahari/services/digitaltwin/internal/domain/simulation"
	"prahari/services/digitaltwin/internal/domain/state"
	"prahari/services/digitaltwin/internal/domain/twin"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveTwin(ctx context.Context, t *twin.DigitalTwin) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO digital_twins (id, plant_id, name, status, version, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, version = EXCLUDED.version, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.PlantID, t.Name, t.Status, t.Version, t.CreatedAt, t.UpdatedAt)
	return err
}

func (s *Store) GetTwinByID(ctx context.Context, id string) (*twin.DigitalTwin, error) {
	if s.db == nil {
		return &twin.DigitalTwin{ID: id, PlantID: "P01", Name: "Distillation Unit Twin", Status: "ACTIVE", Version: 1, CreatedAt: time.Now()}, nil
	}
	query := `SELECT id, plant_id, name, status, version, created_at, updated_at FROM digital_twins WHERE id = $1`
	var t twin.DigitalTwin
	err := s.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.PlantID, &t.Name, &t.Status, &t.Version, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("digital twin %s not found", id)
	}
	return &t, err
}

func (s *Store) SaveLiveState(ctx context.Context, ls *state.LiveState) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO live_states (id, twin_id, equipment_id, value, quality, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET value = EXCLUDED.value, quality = EXCLUDED.quality, timestamp = EXCLUDED.timestamp`
	_, err := s.db.ExecContext(ctx, query, ls.ID, ls.TwinID, ls.EquipmentID, ls.Value, ls.Quality, ls.Timestamp)
	return err
}

func (s *Store) SaveScenario(ctx context.Context, sc *simulation.Scenario) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO simulations (id, twin_id, name, status, parameters, result_data, started_at, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, result_data = EXCLUDED.result_data, completed_at = EXCLUDED.completed_at`
	_, err := s.db.ExecContext(ctx, query, sc.ID, sc.TwinID, sc.Name, sc.Status, sc.Parameters, sc.ResultData, sc.StartedAt, sc.CompletedAt)
	return err
}

func (s *Store) SavePlaybackSession(ctx context.Context, ps *playback.Session) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO playback_sessions (id, twin_id, start_time, end_time, speed, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status`
	_, err := s.db.ExecContext(ctx, query, ps.ID, ps.TwinID, ps.StartTime, ps.EndTime, ps.Speed, ps.Status, ps.CreatedAt)
	return err
}

func (s *Store) SaveOverlay(ctx context.Context, o *overlay.Layer) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO overlays (id, twin_id, layer_type, source_id, label, metadata, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, o.ID, o.TwinID, o.LayerType, o.SourceID, o.Label, o.Metadata, o.Timestamp)
	return err
}

func (s *Store) GetOverlaysByTwin(ctx context.Context, twinID string) ([]*overlay.Layer, error) {
	if s.db == nil {
		return []*overlay.Layer{
			{ID: "lay-1", TwinID: twinID, LayerType: "ALARM", SourceID: "evt-01", Label: "High Temperature Valve", Metadata: `{"alarm_level": "CRITICAL"}`, Timestamp: time.Now()},
		}, nil
	}
	query := `SELECT id, twin_id, layer_type, source_id, label, metadata, timestamp FROM overlays WHERE twin_id = $1`
	rows, err := s.db.QueryContext(ctx, query, twinID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*overlay.Layer
	for rows.Next() {
		var o overlay.Layer
		if err := rows.Scan(&o.ID, &o.TwinID, &o.LayerType, &o.SourceID, &o.Label, &o.Metadata, &o.Timestamp); err != nil {
			return nil, err
		}
		result = append(result, &o)
	}
	return result, nil
}

func (s *Store) SearchTwins(ctx context.Context, criteria *search.Criteria) ([]*twin.DigitalTwin, int64, error) {
	if s.db == nil {
		mockTwin := &twin.DigitalTwin{ID: "twin-001", PlantID: criteria.PlantID, Name: "Crude Unit Twin", Status: "ACTIVE", Version: 1, CreatedAt: time.Now()}
		return []*twin.DigitalTwin{mockTwin}, 1, nil
	}
	query := `SELECT id, plant_id, name, status, version, created_at, updated_at FROM digital_twins`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []*twin.DigitalTwin
	for rows.Next() {
		var t twin.DigitalTwin
		if err := rows.Scan(&t.ID, &t.PlantID, &t.Name, &t.Status, &t.Version, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, 0, err
		}
		result = append(result, &t)
	}
	return result, int64(len(result)), nil
}
