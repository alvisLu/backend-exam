package transfers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alvis/wallet_service/internal/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/transfer", h.transfer)
}

func (h *Handler) transfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(httpx.BadRequest("invalid request body"))
		return
	}

	result, err := h.svc.Transfer(c.Request.Context(), req.FromID, req.ToID, req.Amount)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, toTransferResponse(result))
}
