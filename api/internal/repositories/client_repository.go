package repositories

import (
	"context"

	"myapp/api/internal/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (d *clientRepository) CreateClient(ctx context.Context, entity *entities.Client) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Create(entity).
		Error
}

func (d *clientRepository) UpdateClient(ctx context.Context, entity *entities.Client) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Save(entity).
		Error
}

func (d *clientRepository) DeleteClient(ctx context.Context, entity *entities.Client) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Delete(entity).
		Error
}

func (d *clientRepository) FindClientByID(ctx context.Context, id int) (*entities.Client, error) {
	var entity *entities.Client

	err := d.db.
		Preload(clause.Associations).
		Where("id = ?", id).
		Limit(1).
		Find(&entity).Error

	return entity, err
}

func (d *clientRepository) FindClientByName(ctx context.Context, name string) ([]*entities.Client, error) {
	var entity []*entities.Client

	err := d.db.
		Preload(clause.Associations).
		Where("name = ?", name).
		Limit(100).
		Find(&entity).Error

	return entity, err
}

func (d *clientRepository) FindClientByPartnerID(ctx context.Context, partnerID int, name string) ([]*entities.Client, error) {
	var entity []*entities.Client

	err := d.db.
		Preload(clause.Associations).
		Where("partner_id = ?", partnerID).
		Where("name = ?", name).
		Limit(1).
		Find(&entity).Error

	return entity, err
}

func (d *clientRepository) ListClient(ctx context.Context) ([]*entities.Client, error) {
	//TODO impl pagination
	var entities []*entities.Client

	err := d.db.
		Preload(clause.Associations).
		Limit(100).
		Order("created_at desc").
		Find(&entities).Error

	if err != nil {
		return nil, err
	}

	return entities, nil
}
