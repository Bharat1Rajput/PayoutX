package handler

import (
	"net/http"

	"github.com/Bharat1Rajput/payout-service/internal/model"
	"github.com/Bharat1Rajput/payout-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PayoutHandler struct {
	svc *service.PayoutService
}

func NewPayoutHandler(svc *service.PayoutService) *PayoutHandler{
	return &PayoutHandler{svc : svc}
}


func (h *PayoutHandler)CreatePayout (c *gin.Context){

	var req model.CreatePayoutRequest
	if err := c.ShouldBindJSON(&req); err!= nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return 
	}

	res ,err := h.svc.CreatePayout(c.Request.Context(),req)

	if err!= nil {
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