package handlers

// swagger:route GET /users user EmptyRequest
// Get all users.
//
// responses:
//   200: UsersGetAllResponse
//   400: StatusBadRequest

// swagger:parameters EmptyRequest
type EmptyRequest struct {
}

// swagger:response UsersGetAllResponse
type UsersGetAllResponse struct {
	// list of all users
	//
	// in: body
	Body []string `json:"users"`
}

// ErrorReply replys an error of API calling.
//
// swagger:response
type ErrorReply struct {
	// in: body
	Message string `json:"message"`
}