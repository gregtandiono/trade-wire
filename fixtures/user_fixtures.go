package fixtures

// UserFixtures generates a set list of users and returns an object of user-case scenarios
// will be used to mock registration requests
func UserFixtures() map[string]map[string]string {
	return map[string]map[string]string{
		"validUserSignup": {
			"name":     "George Handsometon",
			"username": "ghandsometon",
			"type":     "admin",
			"password": "password123",
		},
		"validEmployeeSignup": {
			"name":     "Gregory Tandiono",
			"username": "gtandiono",
			"type":     "employee",
			"password": "password123456",
		},
		"invalidUserSignup": {
			"name":     "George Notsohandsometon",
			"username": "gnothandsometon",
			"type":     "",
			"password": "",
		},
		"validUserLogin": {
			"username": "ghandsometon",
			"password": "password123",
		},
		"invalidUserLogin": {
			"username": "ghandsometon",
			"password": "passwo",
		},
	}
}
