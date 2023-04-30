package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	desc "github.com/hmuriyMax/vote/internal/pb/vote_service"
	"net/http"
)

func (s RestServer) vote(c *gin.Context) {
	var req desc.VoteRequest

	// Демаршализация запроса
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating vote request")
	}

	// Использует сервис голосования, чтобы засчитать голос
	resp, err := s.voteService.Vote(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}
