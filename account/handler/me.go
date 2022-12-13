package handler

import (
	"log"
	"net/http"

	"github.com/Nuriddin-Olimjon/memrizr/account/model"
	"github.com/Nuriddin-Olimjon/memrizr/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Me(ctx *gin.Context) {
	// A *model.User will eventually be added to context in middleware

	user, exists := ctx.Get("user")

	// This shouldn't happen, as our middleware ought to throw an error.
	// This is an extra safety measure
	// We'll extract this logic later as it will be common to all handler
	// methods which require a valid user
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", ctx)
		err := apperrors.NewInternal()
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	uid := user.(*model.User).UID

	// gin.Context satisfies go's context.Context interface
	u, err := h.UserService.Get(ctx, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		ctx.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
