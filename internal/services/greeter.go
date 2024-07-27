package services

import (
	"fmt"
	v1 "gin-layout/api/helloworld/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateGreeter 创建Greeter
// @Tags Greeter
// @Summary 创建Greeter
// @Param json body v1.CreateGreeterRequest true "标签"
// @Success 200 {object} v1.CreateGreeterReply
// @Accept json
// @Router /greeter [POST]
func (s AppService) CreateGreeter(ctx *gin.Context) {
	var retErr error
	statusCode := http.StatusBadRequest
	defer func() {
		s.handleError(ctx, retErr, statusCode)
	}()

	var req v1.CreateGreeterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		retErr = fmt.Errorf("invalid request: %s", err.Error())
		return
	}
	if err := req.Validate(); err != nil {
		retErr = fmt.Errorf("vaildate fail: %s", err.Error())
		return
	}

	response, err := s.greeterRepo.Create(ctx, &req)
	if err != nil {
		retErr = fmt.Errorf("failed to create greeter: %s", err.Error())
		statusCode = http.StatusInternalServerError
		return
	}
	s.Response(ctx, http.StatusOK, response)
}
