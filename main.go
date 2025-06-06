package main

import (
	"log/slog"
	"os"

	"github.com/Belxz777/backgo/common/db"
	"github.com/Belxz777/backgo/logic/admin"
	"github.com/Belxz777/backgo/logic/auth/register"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("DATABASE_PORT").(string)
	host := viper.Get("DATABASE_HOST").(string)
	dbUrl := viper.Get("DATABASE_URL").(string)
	dbName := viper.Get("DATABASE_NAME").(string)
	dbUser := viper.Get("DATABASE_USER").(string)
	dbPass := viper.Get("DATABASE_PASSWORD").(string)

	r := gin.Default()
	h := db.Init(host, port, dbUser, dbPass, dbName)
	register.RegisterRoutes(r, h)
	admin.RegisterRoutes(r, h)
	setupLogger(viper.Get("APP_ENV").(string))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	r.Run(":8888")
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "production":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
