package dataTransformer

import "encoding/json"

type HealthDataDetails struct {
	ResourceType string
	Patient      string
}

type DataTransformer struct{}

func (dt DataTransformer) Extract(inJsonData string) HealthDataDetails {
	var data map[string]interface{}
	json.Unmarshal([]byte(inJsonData), &data)

	resourceType := data["resourceType"].(string)
	patient := data["patient"].(map[string]interface{})
	patientReference := patient["reference"].(string)

	return HealthDataDetails{ResourceType: resourceType, Patient: patientReference}
}
