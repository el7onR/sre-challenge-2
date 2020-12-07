package healthz

import (
	"backend/internal/config"

	"github.com/gin-gonic/gin"
)

func HealthzHandler(env *config.Env, r *gin.Engine) {
	r.GET("/healthz", healthzHandler(env))
}
