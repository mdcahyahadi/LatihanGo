package uploadphoto

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetPhotos(ctx context.Context, limit int, searchBy string, keyword string) (resp []PhotoEntity, err error) {
	db := repo.db.WithContext(ctx).
		Table("photosx").
		Select("ID, TITLE, CAPTION, URL, USERID, CREATEDAT, UPDATEDAT").
		Where(searchBy+" = ?", keyword).
		Order("ID DESC")
	if limit > 0 {
		db = db.Limit(limit)
	}
	err = db.Find(&resp).Error
	return
}

func (repo *Repository) StorePhoto(ctx context.Context, req PhotoEntity) (photoID int64, err error) {
	db := repo.db.WithContext(ctx)
	result := db.Create(&req)
	if result.Error != nil {
		err = result.Error
		return
	}
	photoID = req.ID
	return
}

func (repo *Repository) UpdatePhoto(ctx context.Context, req PhotoEntity) (err error) {
	db := repo.db.WithContext(ctx)
	err = db.Model(PhotoEntity{}).
		Where("ID = ?", req.ID).
		Updates(&req).Error
	return
}
