package response

// Response header for data
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
