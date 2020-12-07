package main

import (
	"context"

	"backend/internal/config"
	"backend/internal/healthz"
	"backend/internal/users"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BACKEND")
	viper.SetDefault("DB_NAME", "database")
	viper.SetDefault("DB_ADDRESS", "localhost:3306")
	viper.SetDefault("DB_PASSWORD", "root_pass")
	viper.SetDefault("DB_USERNAME", "root")

	db := config.Database(context.Background(), &config.ConnectionInfo{
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Address:  viper.GetString("DB_ADDRESS"),
		DBName:   viper.GetString("DB_NAME"),
	})

	log := config.Log()

	env := config.Env{
		DB:  db,
		Log: log,
	}

	users.Migrate(&env)

	router := gin.Default()
	router.Use(ginprom.PromMiddleware(nil))
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	users.UsersHandler(&env, router)
	healthz.HealthzHandler(&env, router)

	log.Info("Starting backend")
	err := router.Run(":8080")
	if err != nil {
		env.Log.Fatal(err)
	}
}
