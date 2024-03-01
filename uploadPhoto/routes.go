package uploadphoto

import "github.com/gin-gonic/gin"

// Routes struct contains http delivery
type Routes struct {
	delivery *Delivery
}

// NewRoutes to init http delivery with registered routes
func NewRoutes(delivery *Delivery) *Routes {
	return &Routes{delivery: delivery}
}

// RegisterRoutes to register all rcMapper routes
func (routes Routes) RegisterRoutes(r *gin.Engine) {
	route := r.Group("/v1")
	{
		route.GET("/photo/", routes.delivery.GetPhotos)
		route.POST("/photo/", routes.delivery.SavePhoto)
		route.PUT("/photo/", routes.delivery.EditPhoto)
	}
}
