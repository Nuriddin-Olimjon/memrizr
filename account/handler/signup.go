package handler

import (
	"log"

	"github.com/Nuriddin-Olimjon/memrizr/account/model"
	"github.com/Nuriddin-Olimjon/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	var req signupReq
	if !bindData(c, &req) {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)

	if err != nil {
		log.Printf("Failed to signup user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
}
