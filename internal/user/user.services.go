package user

func GetAllUsersServices(payload any) map[string]interface{} {
	user := map[string]interface{}{
		"id":   1,
		"name": "Jobayar",
	}
	return user
}
func GetUsersByIdServices(payload string) string {
	
	return payload
}