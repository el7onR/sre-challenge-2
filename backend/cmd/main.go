package main

import (
	"context"
	"net/http"

	"backend/internal/config"
	"backend/internal/healthz"
	"backend/internal/metric"
	"backend/internal/users"

	"github.com/gin-gonic/gin"
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

	users.UsersHandler(&env, router)
	healthz.HealthzHandler(&env, router)
	metric.MetricsHandler(&env, router)

	log.Info("Starting backend")

	http.ListenAndServe(":8080", router)

}
