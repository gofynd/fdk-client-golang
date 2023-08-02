package application

import (
	"errors"
	"regexp"

	"github.com/gofynd/fdk-client-golang/sdk/common"
)

// MongoIDRegExp is the `regex` for mongoDB ObehectID
const MongoIDRegExp = `^[0-9a-fA-F]{24}$`

// AppConfig provides configuration to a service client instance.
type AppConfig struct {
	ApplicationID    string
	ApplicationToken string
	Domain           string
	Options          *Options
}

// Options provides other configuration values
type Options struct {
}

//NewAppConfig provides application configuration
func NewAppConfig(applicationID, applicationToken, domain string, options *Options) (*AppConfig, error) {
	if applicationID == "" {
		return &AppConfig{}, common.NewFDKError(errors.New("No Application ID Present").Error())
	}
	if applicationToken == "" {
		return &AppConfig{}, common.NewFDKError(errors.New("No Application Token Present").Error())
	}
	r, err := regexp.Compile(MongoIDRegExp)
	if err != nil {
		return &AppConfig{}, common.NewFDKError(err.Error())
	}
	if !r.MatchString(applicationID) {
		return &AppConfig{}, common.NewFDKError(errors.New("Invalid Application ID. It should be Mongo ID").Error())
	}
	if len(applicationToken) < 5 {
		return &AppConfig{}, common.NewFDKError(errors.New("Invalid Application Token").Error())
	}
	return &AppConfig{applicationID, applicationToken, domain, options}, nil
}
