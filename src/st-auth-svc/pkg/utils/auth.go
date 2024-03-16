package utils

import (
	"encoding/json"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

func censorSensitiveData(resp interface{}) interface{} {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		return resp
	}

	var respMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &respMap); err != nil {
		return resp
	}

	censorMap(respMap)

	return respMap
}

func censorMap(data map[string]interface{}) {
	for key, value := range data {
		switch key {
		case "token", "firstName", "lastName":
			data[key] = "******"
		case "email":
			if email, ok := value.(string); ok {
				data[key] = censorEmail(email)
			}
		default:
			if nestedMap, ok := value.(map[string]interface{}); ok {
				censorMap(nestedMap)
			}
		}
	}
}

func censorEmail(email string) string {
	parts := regexp.MustCompile(`^([^@]+)@(.+)$`).FindStringSubmatch(email)
	if len(parts) == 3 {
		return strings.Repeat("*", len(parts[1])) + "@" + parts[2]
	}
	return email
}

func LogResponse(logger *zap.Logger, action string, resp interface{}, statusCode int) {
	censoredResp := censorSensitiveData(resp)
	respJSON, err := json.Marshal(censoredResp)
	if err != nil {
		logger.Error("Failed to marshal response",
			zap.String("action", action),
			zap.Error(err),
		)
		return
	}

	level := GetStatusLevel(statusCode)
	switch level {
	case "ERROR":
		logger.Error(action+" response", zap.String("response", string(respJSON)), zap.Int("status", statusCode))
	case "WARNING":
		logger.Warn(action+" response", zap.String("response", string(respJSON)), zap.Int("status", statusCode))
	case "INFO":
		logger.Info(action+" response", zap.String("response", string(respJSON)), zap.Int("status", statusCode))
	default:
		logger.Debug(action+" response", zap.String("response", string(respJSON)), zap.Int("status", statusCode))
	}
}

func GetStatusLevel(statusCode int) string {
	switch {
	case statusCode >= 500:
		return "ERROR"
	case statusCode >= 400:
		return "WARNING"
	case statusCode >= 200:
		return "INFO"
	default:
		return "DEBUG"
	}
}

func ToInterfaceSlice(slice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, d := range slice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
