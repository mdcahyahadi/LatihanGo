package uploadphoto

import (
	"context"

	"github.com/gin-gonic/gin"
)

type IRepository interface {
	GetPhotos(ctx context.Context, limit int, searchBy string, keyword string) (resp []PhotoEntity, err error)
	StorePhoto(ctx context.Context, req PhotoEntity) (photoID int64, err error)
	UpdatePhoto(ctx context.Context, req PhotoEntity) (err error)
}

type IService interface {
	GetListPhoto(ctx *gin.Context, req GetListPhoto) (resp []PhotoEntity, err error)
	PostingPhoto(ctx *gin.Context, req SavePhoto) (resp interface{}, err error)
	EditPhoto(ctx *gin.Context, req EditPhoto) (resp interface{}, err error)
}
