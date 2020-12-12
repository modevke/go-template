package types

import(
	"github.com/go-ozzo/ozzo-validation/v4"
)

type EnvironmentVariables struct {
	Port        string
	Host        string
	Scheme      string
	ProjectName string
	Environment string

	DbUser		string
	DbPassword	string
	DbPort		string
	DbHost		string
	DbName		string
}

func (e EnvironmentVariables) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.Port, validation.Required),
		validation.Field(&e.Host, validation.Required),
		validation.Field(&e.Scheme, validation.Required),
		validation.Field(&e.ProjectName, validation.Required),
		validation.Field(&e.Environment, validation.Required),
		validation.Field(&e.DbUser, validation.Required),
		validation.Field(&e.DbPassword, validation.Required),
		validation.Field(&e.DbPort, validation.Required),
		validation.Field(&e.DbHost, validation.Required),
		validation.Field(&e.DbName, validation.Required),
	)
}