package errors

//StatusMsgError represents the status and message of the error response
type StatusMsgError struct {
	Status  int
	Message string
}

//ErrorResponse represents all errors response from the API
var ErrorResponse = map[string]StatusMsgError{
	"notAuthorized": {
		Status:  401,
		Message: "Unauthorized user",
	},
	"categoryAlreadyExists": {
		Status:  409,
		Message: "Category already exists",
	},
	"categoryNotFound": {
		Status:  400,
		Message: `"categoryIds" not found`,
	},
	"missingFields": {
		Status:  400,
		Message: "Some required fields are missing",
	},
	"invalidFields": {
		Status:  400,
		Message: "Invalid fields",
	},
	"emailIsRequired": {
		Status:  400,
		Message: `"email" must be a valid email`,
	},
	"invalidPassword": {
		Status:  400,
		Message: `"password" length must be at least 6 characters long`,
	},
	"passwordIsRequired": {
		Status:  400,
		Message: `"password" is required`,
	},
	"tokenNotFound": {
		Status:  401,
		Message: "Token not found",
	},
	"invalidToken": {
		Status:  401,
		Message: "Expired or invalid token",
	},
}
