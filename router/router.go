package router

import (
	"log/slog"

	"github.com/gaoozi/auc/api"
	"github.com/gin-gonic/gin"
)

func Serve() {
  gin.SetMode(gin.ReleaseMode)

  r := gin.Default()
  r.GET("ping", api.Ping)

  if err := r.Run(":8000"); err != nil {
    slog.Error("start server err:", err)
  }
}
