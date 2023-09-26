package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"github.com/reubenmiller/go-c8y/pkg/c8y"

	"github.com/scfx/goneric-mqtt/handlers"
)

func main() {
	logger := log.New(os.Stdout, "Mapping: ", log.LstdFlags)

	//Create handlers
	hh := handlers.NewHealth()
	eh := handlers.NewEnvironment()
	ch := handlers.NewCounter()

	user := handlers.ApplicationTenant() + "/" + handlers.ApplicationUser()
	password := handlers.ApplicationPassword()
	url := "hackathon1.dev.c8y.io"
	topic := "nojava"
	clientId := "mapping-service-local"
	broker := fmt.Sprintf("mqtts://%s:9883", url)
	logger.Printf("Starting mapping service")

	restClient := c8y.NewClient(nil, "https://"+url, handlers.ApplicationTenant(), handlers.ApplicationUser(), password, true)

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId).SetUsername(user).SetPassword(password)

	logger.Printf("Connect to broker %s and clientId %s with user: %s and password: %s", broker, clientId, user, password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Printf("Error connecting to broker %s: %s", broker, token.Error())
	}
	logger.Printf("Connected to broker %s", broker)
	if token := client.Subscribe(topic, 0, onTemperatureMessage(logger, restClient, ch)); token.Wait() && token.Error() != nil {
		logger.Printf("Error subscribing to topic %s: %s", topic, token.Error())
	}

	r := mux.NewRouter()
	r.Handle("/health", hh).Methods("GET")
	r.Handle("/env", eh).Methods("GET")
	r.Handle("/counter", ch).Methods("GET")
	http.ListenAndServe(":80", r)
}

func onTemperatureMessage(log *log.Logger, restClient *c8y.Client, counter *handlers.Counter) func(mqtt.Client, mqtt.Message) {
	return func(client mqtt.Client, message mqtt.Message) {
		log.Printf("Received message on topic %s: %s", message.Topic(), message.Payload())
		temperature, err := NewTemperature(message.Payload())
		if err != nil {
			log.Printf("Error parsing payload: %s", err)
			return
		}
		err = temperature.Post(restClient)
		if err != nil {
			log.Printf("Error posting temperature: %s", err)
			return
		}
		log.Printf("Temperature: %f", temperature.Temperature)
		counter.AddOne()
	}
}
