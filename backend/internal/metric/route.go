package metric

import (
	"backend/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(env *config.Env, r *gin.Engine) {
	r.GET("/metrics", gin.WrapF(promhttp.Handler().ServeHTTP))
}
