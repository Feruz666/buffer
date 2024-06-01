package handler

import (
	"github.com/Feruz666/buffer/internal/store"
	"github.com/gin-gonic/gin"
)

type BuffHandler struct {
	s *store.Store
}

func New(s *store.Store) *BuffHandler {
	return &BuffHandler{s: s}
}

func (b *BuffHandler) RouterFunc() *gin.Engine {
	r := gin.New()

	r.POST("/save", b.ProcessData)

	return r
}
