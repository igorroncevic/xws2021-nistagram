package kafka_util

func ConvertToPerformanceMessage(message map[string]interface{}) PerformanceMessage{
	converted := PerformanceMessage{}

	if _, ok := message["service"]; ok {
		value, _ := message["service"].(string)
		converted.Service = value
	}

	if _, ok := message["function"]; ok {
		value, _ := message["function"].(string)
		converted.Function = value
	}

	if _, ok := message["status"]; ok {
		value, _ := message["status"].(float64)
		converted.Status = int(value)
	}

	if _, ok := message["message"]; ok {
		value, _ := message["message"].(string)
		converted.Message = value
	}

	return converted
}


