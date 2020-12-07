package users

import (
	"backend/internal/config"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

const (
	errorFetchingData     = "Error while fetching data %s from database, details %+v"
	errorMarshallingJSON  = "Error while marshalling or unmarshalling json %+v"
	userRequestedNotFound = "User requested not found %+v"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//FindAllUsersHandler retrieves all users from database
func FindAllUsersHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		env.Log.Info("Getting users")

		users, err := List(env)
		if len(*users) == 0 {
			env.Log.Infof(userRequestedNotFound, err)
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			env.Log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, &users)
	}
}

//FindByIDUserHandler retrieves an specific user from database
func FindByIDUserHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		env.Log.Info("Getting user by id")

		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			env.Log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err)
			return
		}

		u, err := ByID(env, &id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			env.Log.Infof(userRequestedNotFound, err.Error())
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			env.Log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, &u)
	}
}

//CreateUserHandler persists a new user on database
func CreateUserHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		env.Log.Info("Creating user")
		u := User{}
		err := c.BindJSON(&u)
		if err != nil {
			env.Log.Errorf("Error while reading body %+v", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = validate.Struct(&u)
		if err != nil {
			env.Log.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		created, err := New(env, &u)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			env.Log.Infof(userRequestedNotFound, err)
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			env.Log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, created)
	}
}

//DeleteUserHandler deletes an user from database
func DeleteUserHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		env.Log.Info("Deleting user")
		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			env.Log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = Delete(env, &id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			env.Log.Infof(userRequestedNotFound, err)
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			env.Log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Status(http.StatusOK)
	}
}

//UpdateUserHandler updates an user from database
func UpdateUserHandler(env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		env.Log.Info("Updating user")

		id, err := strconv.ParseInt(c.Param("id"), 10, 0)
		if err != nil {
			env.Log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err)
			return
		}

		u := User{}
		err = c.BindJSON(&u)
		if err != nil {
			env.Log.Errorf("Error while reading body %+v", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		updated, err := Update(env, &u, &id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			env.Log.Infof(userRequestedNotFound, err)
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			env.Log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, updated)
	}
}
