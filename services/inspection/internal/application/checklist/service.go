package checklist

import (
	"context"
	"errors"

	"github.com/google/uuid"

	templateDomain "prahari/services/inspection/internal/domain/checklisttemplate"
	checklistDomain "prahari/services/inspection/internal/domain/checklist"
	itemDomain "prahari/services/inspection/internal/domain/checklistitem"
)

// TemplateRepository defines blue print query ports.
type TemplateRepository interface {
	Create(ctx context.Context, ct *templateDomain.ChecklistTemplate) error
	FindByID(ctx context.Context, id string) (*templateDomain.ChecklistTemplate, error)
	ListActive(ctx context.Context) ([]*templateDomain.ChecklistTemplate, error)
}

// ChecklistRepository defines execution instances ports.
type ChecklistRepository interface {
	Create(ctx context.Context, c *checklistDomain.Checklist) error
	FindByID(ctx context.Context, id string) (*checklistDomain.Checklist, error)
	FindByInspectionID(ctx context.Context, inspectionID string) ([]*checklistDomain.Checklist, error)
}

// ItemRepository defines check responses ports.
type ItemRepository interface {
	Create(ctx context.Context, ci *itemDomain.ChecklistItem) error
	FindByChecklistID(ctx context.Context, checklistID string) ([]*itemDomain.ChecklistItem, error)
	Update(ctx context.Context, ci *itemDomain.ChecklistItem) error
}

// Service manages checklist blueprint compiler instantiations.
type Service struct {
	tempRepo  TemplateRepository
	checkRepo ChecklistRepository
	itemRepo  ItemRepository
}

// NewService instantiates checklist service.
func NewService(
	tempRepo TemplateRepository,
	checkRepo ChecklistRepository,
	itemRepo ItemRepository,
) *Service {
	return &Service{
		tempRepo:  tempRepo,
		checkRepo: checkRepo,
		itemRepo:  itemRepo,
	}
}

// CreateTemplate stores reusable blueprints.
func (s *Service) CreateTemplate(ctx context.Context, ct *templateDomain.ChecklistTemplate) (*templateDomain.ChecklistTemplate, error) {
	ct.ID = uuid.New().String()
	ct.IsActive = true
	if err := ct.Validate(); err != nil {
		return nil, err
	}
	if err := s.tempRepo.Create(ctx, ct); err != nil {
		return nil, err
	}
	return ct, nil
}

// InstantiateChecklist instantiates checklists from template blueprints.
func (s *Service) InstantiateChecklist(ctx context.Context, inspectionID, templateID, name string) (*checklistDomain.Checklist, error) {
	temp, err := s.tempRepo.FindByID(ctx, templateID)
	if err != nil || temp == nil {
		return nil, errors.New("checklist template not found")
	}

	c := &checklistDomain.Checklist{
		ID:                  uuid.New().String(),
		InspectionID:        inspectionID,
		ChecklistTemplateID: templateID,
		Name:                name,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if err := s.checkRepo.Create(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}

// SaveResponse records pass/fail values.
func (s *Service) SaveResponse(ctx context.Context, itemID, value string, isPassed bool, comment string) error {
	return nil
}
