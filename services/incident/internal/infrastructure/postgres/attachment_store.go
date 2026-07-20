package postgres

import (
	"context"
	"database/sql"
	"fmt"

	attachmentDomain "prahari/services/incident/internal/domain/attachment"
)

// AttachmentStore implements the attachment persistence adapter against PostgreSQL.
type AttachmentStore struct {
	db *sql.DB
}

// NewAttachmentStore constructs an AttachmentStore.
func NewAttachmentStore(db *sql.DB) *AttachmentStore {
	return &AttachmentStore{db: db}
}

// Create persists attachment metadata.
func (s *AttachmentStore) Create(ctx context.Context, a *attachmentDomain.Attachment) error {
	query := `INSERT INTO incident_attachments (id, incident_id, file_name, file_size, content_type, storage_path, uploaded_by, uploaded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.IncidentID, a.FileName, a.FileSize, a.ContentType, a.StoragePath, a.UploadedBy, a.UploadedAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert attachment: %w", err)
	}
	return nil
}

// FindByIncidentID retrieves all attachments for an incident.
func (s *AttachmentStore) FindByIncidentID(ctx context.Context, incidentID string) ([]*attachmentDomain.Attachment, error) {
	query := `SELECT id, incident_id, file_name, file_size, content_type, storage_path, uploaded_by, uploaded_at
		FROM incident_attachments WHERE incident_id = $1 ORDER BY uploaded_at DESC`
	rows, err := s.db.QueryContext(ctx, query, incidentID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list attachments: %w", err)
	}
	defer rows.Close()

	var attachments []*attachmentDomain.Attachment
	for rows.Next() {
		a := &attachmentDomain.Attachment{}
		if err := rows.Scan(&a.ID, &a.IncidentID, &a.FileName, &a.FileSize, &a.ContentType, &a.StoragePath, &a.UploadedBy, &a.UploadedAt); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan attachment: %w", err)
		}
		attachments = append(attachments, a)
	}
	return attachments, nil
}
