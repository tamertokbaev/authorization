package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Handler struct {
	healthcheckFn func() error
	baseUrl       string
}

func NewHandlerDelivery(
	healthcheckFn func() error,
	baseUrl string,
) *Handler {
	return &Handler{
		baseUrl:       baseUrl,
		healthcheckFn: healthcheckFn,
	}
}

func (h *Handler) Init() (*gin.Engine, error) {
	app := gin.New()
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	app.GET("/readiness", func(c *gin.Context) {
		if err := h.healthcheckFn(); err != nil {
			c.JSON(http.StatusServiceUnavailable, map[string]string{"message": err.Error()})
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"message": "ok"})
		}
	})
	app.GET("/liveness", func(c *gin.Context) {
		if err := h.healthcheckFn(); err != nil {
			c.JSON(http.StatusServiceUnavailable, map[string]string{"message": err.Error()})
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"message": "ok"})
		}
	})
	h.initAPI(app)
	return app, nil
}

func (h *Handler) initAPI(router *gin.Engine) {
	//baseUrl := router.Group(h.baseUrl)

	router.GET(h.baseUrl+"-docs/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//api := baseUrl.Group("/api")
}
