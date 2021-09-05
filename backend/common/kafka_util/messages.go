package kafka_util

func GetUserEventMessage(eventType UserEventType, success bool) string{
	switch eventType {
	case Login:
		if success { return "Successful login attempt." }
		return "Failed login attempt."
	case PasswordChange:
		if success { return "Successful password change." }
		return "Failed password change."
	case ProfileUpdate:
		if success { return "Successful profile update." }
		return "Failed profile update."
	default:
		return ""
	}
}