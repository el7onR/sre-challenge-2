package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"backend/internal/durable"
	"backend/internal/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func users(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Getting users")

	users := []model.User{}

	result := db.Find(&users)

	if result.Error != nil {
		log.Error("No users found")
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	output, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func userById(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Getting user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	u := model.User{}
	result := db.Take(&u, &id)
	if result.Error != nil {
		log.WithError(result.Error)
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	output, err := json.Marshal(&u)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Creating user")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var u model.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := db.FirstOrCreate(&u)
	if result.Error != nil {
		log.WithError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Created user '%d'.", &u.ID)

	output, err := json.Marshal(&u)
	if err != nil {
		log.WithError(err).Error(output)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	log.Info("Deleting user")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := model.User{}
	result := db.Delete(&u, &id)
	if result.Error != nil {
		log.WithError(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		log.Info(result.RowsAffected)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	log.Infof("Deleted user '%d'.", id)

}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("BACKEND")
	viper.SetDefault("DB_NAME", "database")
	viper.SetDefault("DB_ADDRESS", "localhost:3306")
	viper.SetDefault("DB_PASSWORD", "root_pass")
	viper.SetDefault("DB_USERNAME", "root")

	db, err = durable.OpenDatabaseClient(context.Background(), &durable.ConnectionInfo{
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Address:  viper.GetString("DB_ADDRESS"),
		DBName:   viper.GetString("DB_NAME"),
	})

	if err != nil {
		log.Fatal("Error while opening connection to the database")
		os.Exit(1)
	}

	db.AutoMigrate(model.User{})

	router := mux.NewRouter()
	router.HandleFunc("/users", users).Methods("GET")
	router.HandleFunc("/healthz", healthzHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userById).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", deleteUser).Methods("DELETE")

	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting backend")

	http.ListenAndServe(":8080", router)

}
