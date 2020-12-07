package healthz

import (
	"backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthzHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlDB, err := env.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		d := Dependency{databaseStats: sqlDB.Stats()}
		c.JSON(http.StatusOK, gin.H{"databaseStatus": d.databaseStats})

	}
}
