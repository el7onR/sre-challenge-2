package users

import (
	"backend/internal/config"

	"github.com/gin-gonic/gin"
)

func UsersHandler(env *config.Env, r *gin.Engine) {

	r.GET("/users", FindAllUsersHandler(env))
	r.GET("/users/:id", FindByIDUserHandler(env))
	r.POST("/users", CreateUserHandler(env))
	r.DELETE("/users/:id", DeleteUserHandler(env))
	r.PUT("/users/:id", UpdateUserHandler(env))
}
