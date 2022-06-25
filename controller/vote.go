package controller

import (
	"awesomeProject/logic"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs)
		return
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
