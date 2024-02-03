package main

import (
	"log/slog"
	"os"

	"github.com/gaoozi/auc/config"
	"github.com/gaoozi/auc/router"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    slog.Error("Error loading .env file")
    os.Exit(1)
  }

  config.GetConfig()
  router.Serve()
}
