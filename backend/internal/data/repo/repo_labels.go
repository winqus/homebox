package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/label"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

type LabelRepository struct {
	db *ent.Client
}
type (
	LabelCreate struct {
		Name        string `json:"name" validate:"required,min=1,max=255"`
		Description string `json:"description" validate:"max=255"`
		Color       string `json:"color"`
	}

	LabelUpdate struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name" validate:"required,min=1,max=255"`
		Description string    `json:"description" validate:"max=255"`
		Color       string    `json:"color"`
	}

	LabelSummary struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	LabelOut struct {
		LabelSummary
		Items []ItemSummary `json:"items"`
	}
)

func mapLabelSummary(label *ent.Label) LabelSummary {
	return LabelSummary{
		ID:          label.ID,
		Name:        label.Name,
		Description: label.Description,
		CreatedAt:   label.CreatedAt,
		UpdatedAt:   label.UpdatedAt,
	}
}

var (
	mapLabelOutErr = mapTErrFunc(mapLabelOut)
	mapLabelsOut   = mapTEachErrFunc(mapLabelSummary)
)

func mapLabelOut(label *ent.Label) LabelOut {
	return LabelOut{
		LabelSummary: mapLabelSummary(label),
		Items:        mapEach(label.Edges.Items, mapItemSummary),
	}
}

func (r *LabelRepository) getOne(ctx context.Context, where ...predicate.Label) (LabelOut, error) {
	return mapLabelOutErr(r.db.Label.Query().
		Where(where...).
		WithGroup().
		WithItems(func(iq *ent.ItemQuery) {
			iq.Where(item.Archived(false))
		}).
		Only(ctx),
	)
}

func (r *LabelRepository) GetOne(ctx context.Context, ID uuid.UUID) (LabelOut, error) {
	return r.getOne(ctx, label.ID(ID))
}

func (r *LabelRepository) GetOneByGroup(ctx context.Context, gid, ld uuid.UUID) (LabelOut, error) {
	return r.getOne(ctx, label.ID(ld), label.HasGroupWith(group.ID(gid)))
}

func (r *LabelRepository) GetAll(ctx context.Context, groupId uuid.UUID) ([]LabelSummary, error) {
	return mapLabelsOut(r.db.Label.Query().
		Where(label.HasGroupWith(group.ID(groupId))).
		Order(ent.Asc(label.FieldName)).
		WithGroup().
		All(ctx),
	)
}

func (r *LabelRepository) Create(ctx context.Context, groupdId uuid.UUID, data LabelCreate) (LabelOut, error) {
	label, err := r.db.Label.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetColor(data.Color).
		SetGroupID(groupdId).
		Save(ctx)
	if err != nil {
		return LabelOut{}, err
	}

	label.Edges.Group = &ent.Group{ID: groupdId} // bootstrap group ID
	return mapLabelOut(label), err
}

func (r *LabelRepository) update(ctx context.Context, data LabelUpdate, where ...predicate.Label) (int, error) {
	if len(where) == 0 {
		panic("empty where not supported empty")
	}

	return r.db.Label.Update().
		Where(where...).
		SetName(data.Name).
		SetDescription(data.Description).
		SetColor(data.Color).
		Save(ctx)
}

func (r *LabelRepository) Update(ctx context.Context, data LabelUpdate) (LabelOut, error) {
	_, err := r.update(ctx, data, label.ID(data.ID))
	if err != nil {
		return LabelOut{}, err
	}

	return r.GetOne(ctx, data.ID)
}

func (r *LabelRepository) UpdateByGroup(ctx context.Context, GID uuid.UUID, data LabelUpdate) (LabelOut, error) {
	_, err := r.update(ctx, data, label.ID(data.ID), label.HasGroupWith(group.ID(GID)))
	if err != nil {
		return LabelOut{}, err
	}

	return r.GetOne(ctx, data.ID)
}

func (r *LabelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Label.DeleteOneID(id).Exec(ctx)
}

func (r *LabelRepository) DeleteByGroup(ctx context.Context, gid, id uuid.UUID) error {
	_, err := r.db.Label.Delete().
		Where(
			label.ID(id),
			label.HasGroupWith(group.ID(gid)),
		).Exec(ctx)

	return err
}
