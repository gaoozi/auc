package router

import (
	"log/slog"
	"strconv"

	"github.com/gaoozi/auc/api"
	"github.com/gaoozi/auc/config"
	"github.com/gin-gonic/gin"
)

func Serve() {
  gin.SetMode(gin.ReleaseMode)

  r := gin.Default()
  r.GET("ping", api.Ping)

  conf := config.GetConfig()
  if err := r.Run(":" + strconv.Itoa(conf.Server.Port)); err != nil {
    slog.Error("start server err:", err)
  }
}
