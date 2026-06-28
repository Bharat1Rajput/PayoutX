package handler

import (
	"net/http"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PayoutHandler struct {
	svc *service.PayoutService
}

func NewPayoutHandler(svc *service.PayoutService) *PayoutHandler {
	return &PayoutHandler{svc: svc}
}

func (h *PayoutHandler) CreatePayout(c *gin.Context) {

	var req model.CreatePayoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	res, err := h.svc.CreatePayout(c.Request.Context(), req)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		res,
	)
}

func (h *PayoutHandler) HandleBankWebhook(
	c *gin.Context,
) {
	var req model.BankWebhookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.svc.UpdatePayoutStatus(
		c.Request.Context(),
		req.PayoutID,
		req.Status,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "webhook processed",
	})
}

func (h *PayoutHandler) UpdatePayoutStatus(
	c *gin.Context,
) {

	payoutID := c.Param("id")

	var req model.UpdateStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.svc.UpdatePayoutStatus(
		c.Request.Context(),
		payoutID,
		req.Status,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "status updated",
	})
}

func (h *PayoutHandler) UpdateBankReference(
	c *gin.Context,
) {

	payoutID := c.Param("id")

	var req model.UpdateBankReferenceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.svc.UpdateBankReference(
		c.Request.Context(),
		payoutID,
		req.BankReference,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "bank reference updated",
		},
	)
}
