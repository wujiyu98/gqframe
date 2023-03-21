package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wujiyu98/gqframe/model"
)

type Front struct {
}

func NewFront() *Front {
	return &Front{}
}
func (s Front) Index(ctx *gin.Context) (rep model.IndexRep, err error) {

	return
}
