package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DaffaJatmiko/smart-home-energy/model"
)

type AIModelConnector struct {
	Client *http.Client
}

func NewAIModelConnector() *AIModelConnector {
	return &AIModelConnector{
		Client: &http.Client{},
	}
}

func (c *AIModelConnector) ConnectAIModel(payload model.Inputs, token string) (model.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return model.Response{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq", bytes.NewBuffer(jsonData))
	if err != nil {
		return model.Response{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return model.Response{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return model.Response{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}

	var response model.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.Response{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
