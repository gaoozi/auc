package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gaoozi/auc/db"
	"github.com/gaoozi/auc/model"
	"github.com/gaoozi/auc/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(ctx *gin.Context) {
  ctx.String(http.StatusOK, "pong")
}

func Register(ctx *gin.Context) {
  resp := ApiResponse{Ctx: ctx}
  
  req := model.RegisterRequest{}
  if err := ctx.ShouldBindJSON(&req); err != nil {
    slog.Error("Register request json err:", err)
    resp.Error(BodyBindErr, err.Error())
    return
  }

  if req.Password != req.CheckPassword {
    slog.Error("Two passwords must be the same")
    resp.Error(RegisterErr, "Two password must be the same")
    return
  }

  result := db.GetDb().First(&model.User{}, "Username = ?", req.Username)
  if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
    slog.Error("Username is already in use")
    resp.Error(ParamErr, "Username is already in use")
    return
  }

  passwordHash, err := util.GeneratePasswordHash(req.Password)
  if err != nil {
    slog.Error("Hash password failed:", err)
    resp.Error(ParamErr, "Hash password failed")
    return
  }

  user := &model.User {
    Username: req.Username,
    Password: string(passwordHash),
  }

  result = db.GetDb().Create(user)
  if result.Error != nil {
    slog.Error("Create user failed:", result.Error)
    resp.Error(UpdateErr, "Create user failed")
    return
  }

  resp.WithData(user.ID)
}
