package users

import "backend/internal/config"

func Migrate(env *config.Env) {
	env.DB.AutoMigrate(User{})
}
