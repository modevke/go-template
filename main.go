package main

import (
	"fmt"
	"log"
	rn "math/rand"
	"os"
	"time"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/modevke/go-template/docs"
	"github.com/modevke/go-template/db"
	"github.com/modevke/go-template/infrastructure/types"
	"github.com/modevke/go-template/interfaces"
)

func init() {
	rn.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// @title Golang template
// @version 1.0.0
// @description Golang REST API template.
func main() {

	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// LOAD AND VALIDATE ENV VARIABLES
	var host string
	var envVars types.EnvironmentVariables

	if envErr := godotenv.Load(); envErr != nil {
		err = fmt.Errorf(".env file is missing")
		return
	}

	envVars.Port = os.Getenv("API_PORT")
	envVars.Host = os.Getenv("API_HOST")
	envVars.ProjectName = os.Getenv("PROJECT_NAME")
	envVars.Scheme = os.Getenv("API_SCHEME")
	envVars.Environment = os.Getenv("GO_ENV")
	envVars.DbUser = os.Getenv("DB_USER")
	envVars.DbPassword = os.Getenv("DB_PASSWORD")
	envVars.DbPort = os.Getenv("DB_PORT")
	envVars.DbHost = os.Getenv("DB_HOST")
	envVars.DbName = os.Getenv("DB_NAME")

	if verr := envVars.Validate(); verr != nil {
		err = verr
		return
	}

	// SWAGGER VARIABLES
	host = envVars.Host + ":" + envVars.Port
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = []string{envVars.Scheme}
	docs.SwaggerInfo.BasePath = "/api/v1"

	// DATABASE CONNECTION
	services, dberr := db.NewRepositories(envVars.DbUser, envVars.DbPassword, envVars.DbPort, envVars.DbHost, envVars.DbName)
	if dberr != nil {
		err = dberr
		return
	}
	defer services.Close()

	// ROUTER
	router := interfaces.Routing(&envVars)
	app := http.Server{
		Handler: router,
		Addr:    host,

		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	if serr := app.ListenAndServe(); serr != http.ErrServerClosed {
		err = fmt.Errorf("Unable to start server: %v", serr)
		return
	}
}
