package v1

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	_ = api.Group("/v1")
	{
	}
}
