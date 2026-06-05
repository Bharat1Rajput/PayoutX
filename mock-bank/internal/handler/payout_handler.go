package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/model"
	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/service"
)

type PayoutHandler struct {
	service *service.PayoutService
}

func NewPayoutHandler(
	service *service.PayoutService,
) *PayoutHandler {
	return &PayoutHandler{
		service: service,
	}
}

func (h *PayoutHandler) CreatePayout(
	c *gin.Context,
) {

	var req model.CreatePayoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.service.CreatePayout(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}