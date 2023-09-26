package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Environment struct {
	BaseUrl           string `json:"baseUrl"`
	BootstrapPassword string `json:"bootstrapPassword"`
	BootstrapUser     string `json:"bootstrapUser"`
	BootstrapTenant   string `json:"bootstrapTenant"`
	Tenant            string `json:"tenant"`
	User              string `json:"user"`
	Password          string `json:"password"`
}

func (e Environment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := e.Encode(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (e Environment) Encode(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func NewEnvironment() *Environment {
	return &Environment{
		BaseUrl:           BaseURL(),
		BootstrapPassword: BootstrapPassword(),
		BootstrapUser:     BootstrapUser(),
		BootstrapTenant:   BootstrapTenant(),
		User:              ApplicationUser(),
		Password:          ApplicationPassword(),
		Tenant:            ApplicationTenant(),
	}
}

func ApplicationTenant() string {
	return os.Getenv("C8Y_TENANT")
}

func BaseURL() string {
	return os.Getenv("C8Y_BASEURL")
}

func BootstrapUser() string {
	return os.Getenv("C8Y_BOOTSTRAP_USER")
}

func BootstrapPassword() string {
	return os.Getenv("C8Y_BOOTSTRAP_PASSWORD")
}

func ApplicationUser() string {
	return os.Getenv("C8Y_USER")
}

func ApplicationPassword() string {
	return os.Getenv("C8Y_PASSWORD")
}

func BootstrapTenant() string {
	return os.Getenv("C8Y_BOOTSTRAP_TENANT")
}
