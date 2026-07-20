package postgres

import (
	"context"
	"database/sql"

	attachmentDomain "prahari/services/inspection/internal/domain/attachment"
)

// AttachmentStore implements storage files referencing.
type AttachmentStore struct {
	db *sql.DB
}

// NewAttachmentStore instantiates AttachmentStore.
func NewAttachmentStore(db *sql.DB) *AttachmentStore {
	return &AttachmentStore{db: db}
}

// Create stores metadata.
func (s *AttachmentStore) Create(ctx context.Context, a *attachmentDomain.Attachment) error {
	query := `INSERT INTO inspection_attachments (id, inspection_id, file_name, file_size, content_type, storage_path, uploaded_by, uploaded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.InspectionID, a.FileName, a.FileSize, a.ContentType, a.StoragePath, a.UploadedBy, a.UploadedAt)
	return err
}

// FindByInspectionID returns uploads metadata lists.
func (s *AttachmentStore) FindByInspectionID(ctx context.Context, inspectionID string) ([]*attachmentDomain.Attachment, error) {
	query := `SELECT id, inspection_id, file_name, file_size, content_type, storage_path, uploaded_by, uploaded_at
		FROM inspection_attachments WHERE inspection_id = $1 ORDER BY uploaded_at DESC`
	rows, err := s.db.QueryContext(ctx, query, inspectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []*attachmentDomain.Attachment
	for rows.Next() {
		a := &attachmentDomain.Attachment{}
		err = rows.Scan(&a.ID, &a.InspectionID, &a.FileName, &a.FileSize, &a.ContentType, &a.StoragePath, &a.UploadedBy, &a.UploadedAt)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, a)
	}
	return attachments, nil
}
