package main

import (
	"context"
	"encoding/json"

	"github.com/reubenmiller/go-c8y/pkg/c8y"
)

type Temperature struct {
	Temperature float64 `json:"temperature"`
	Id          string  `json:"id"`
}

func (t Temperature) Post(client *c8y.Client) error {
	opt := c8y.SimpleMeasurementOptions{
		SourceID:            t.Id,
		ValueFragmentType:   "c8y_TemperatureMeasurement",
		ValueFragmentSeries: "T",
		Value:               t.Temperature,
		Unit:                "C",
		Type:                "c8y_TemperatureMeasurement",
	}
	m, err := c8y.NewSimpleMeasurementRepresentation(opt)
	if err != nil {
		return err
	}
	_, _, err = client.Measurement.Create(context.Background(), *m)
	return err
}

//Unmarshal Payload to Measurement

func NewTemperature(payload []byte) (*Temperature, error) {
	var temperature Temperature
	err := json.Unmarshal(payload, &temperature)
	if err != nil {
		return nil, err
	}
	return &temperature, nil
}
