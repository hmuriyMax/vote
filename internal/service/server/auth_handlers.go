package server

import (
	"github.com/gin-gonic/gin"
	auth "github.com/hmuriyMax/vote/internal/pb/auth_service"
	"google.golang.org/grpc/status"
	"net/http"
)

func (s RestServer) login(ctx *gin.Context) {
	var req *auth.AuthRequest
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := s.authService.Auth(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": status.Code(err),
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  resp.GetStatus(),
		"token":   resp.GetToken().GetToken(),
		"expires": resp.GetToken().GetExpires(),
	})
}

func (s RestServer) register(ctx *gin.Context) {
	var req *auth.AuthRequest
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := s.authService.Register(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": status.Code(err),
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  resp.GetStatus(),
		"token":   resp.GetToken().GetToken(),
		"expires": resp.GetToken().GetExpires(),
	})
}
