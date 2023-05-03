package service

import (
	"encoding/json"
	"os"
)

type Data struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Payload struct {
	Data []Data
}

// raw returns the raw data from the JSON file.
func raw() ([]Data, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw data into a Payload struct.
	var payload Payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		return nil, err
	}
	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetById(idx int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}

	// If the index is out of range, return an empty slice.
	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}
	return data[idx], nil
}
