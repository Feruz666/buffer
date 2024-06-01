package handler

import (
	"context"
	"github.com/Feruz666/buffer/internal/store"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	timeFormat = "2024-05-01"
)

type BufferDataRequest struct {
	PeriodStart         string `form:"period_start" binding:"required"`
	PeriodEnd           string `form:"period_end" binding:"required"`
	PeriodKey           string `form:"period_key" binding:"required"`
	IndicatorToMoId     string `form:"indicator_to_mo_id" binding:"required"`
	IndicatorToMoFactId string `form:"indicator_to_mo_fact_id" binding:"required"`
	Value               string `form:"value" binding:"required"`
	FactTime            string `form:"fact_time" binding:"required"`
	IsPlan              string `form:"is_plan" binding:"required"`
	AuthUserId          string `form:"auth_user_id" binding:"required"`
	Comment             string `form:"comment" binding:"required"`
}

func (b *BuffHandler) ProcessData(c *gin.Context) {

	bufferRequest := &BufferDataRequest{}
	if err := c.Bind(bufferRequest); err != nil {
		log.Println("Binding error, err:", err)
		return
	}

	if err := b.s.SaveData(
		context.Background(),
		&store.BufferData{
			PeriodStart:         bufferRequest.PeriodStart,
			PeriodEnd:           bufferRequest.PeriodEnd,
			PeriodKey:           bufferRequest.PeriodKey,
			IndicatorToMoId:     bufferRequest.IndicatorToMoId,
			IndicatorToMoFactId: bufferRequest.IndicatorToMoFactId,
			Value:               bufferRequest.Value,
			FactTime:            bufferRequest.FactTime,
			IsPlan:              bufferRequest.IsPlan,
			AuthUserId:          bufferRequest.AuthUserId,
			Comment:             bufferRequest.Comment,
		},
	); err != nil {
		log.Println(err)
		return
	}

}
