package responseutil

func Rest(status int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

type ControllerMeta struct {
	Status  int
	Error   error
	Message string
}
