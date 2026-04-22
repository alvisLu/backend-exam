package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/alvis/wallet_service/internal/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/accounts", h.createAccount)
	r.GET("/accounts/:id", h.getAccount)
}

func (h *Handler) createAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(httpx.BadRequest("invalid request body"))
		return
	}

	acc, err := h.svc.CreateAccount(c.Request.Context(), req.Name)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, toAccountResponse(acc))
}

func (h *Handler) getAccount(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(httpx.BadRequest("invalid account id"))
		return
	}

	acc, err := h.svc.GetAccount(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, toAccountResponse(acc))
}
