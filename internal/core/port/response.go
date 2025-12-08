package port

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FieldResponse struct {
	ID           uint   `json:"ID"`
	Name         string `json:"name"`
	PricePerHour int    `json:"price_per_hour"`
	Location     string `json:"location"`
}
