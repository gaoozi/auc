package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DebugType  int
	ApiCode    int
)

const (
	Success     ApiCode = 0
	BodyBindErr ApiCode = 10001
	ParamErr    ApiCode = 10002
	RegisterErr ApiCode = 10003
	LoginErr    ApiCode = 10004
	LogoutErr   ApiCode = 10005
	GetErr      ApiCode = 10006
	UpdateErr   ApiCode = 10007
  HashErr     ApiCode = 10008
	Unknown     ApiCode = 20000
)

type ApiResponse struct {
	Code    ApiCode
	Data    interface{}
	Message string
	Ctx     *gin.Context `json:"-"`
}

func (resp *ApiResponse) Error(code ApiCode, message string) {
  resp.Code = code
  resp.Message = message
  resp.Ctx.JSON(http.StatusInternalServerError, resp)
}

func (resp *ApiResponse) Success() {
  resp.Code = Success
  resp.Message = "success"
  resp.Ctx.JSON(http.StatusOK, resp)
}

func (resp *ApiResponse) WithData(data interface{}) {
  resp.Code = Success
  resp.Message = "success"
  resp.Data = data
  resp.Ctx.JSON(http.StatusOK, resp)
}
