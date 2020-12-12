package interfaces

import (

	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/modevke/go-template/infrastructure/types"
)

func Routing(envVars *types.EnvironmentVariables) http.Handler{
	r := mux.NewRouter()

	// ROUTER GROUPS
	// api := r.PathPrefix("/api/v1").Subrouter()

	//SWAGGER
	swaggerURL := "/api/documentation/"
	r.PathPrefix(swaggerURL).Handler(httpSwagger.Handler(
		httpSwagger.URL(envVars.Scheme+"://"+envVars.Host+":"+envVars.Port+swaggerURL+"doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	return r
}