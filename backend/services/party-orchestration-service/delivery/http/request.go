package http

type CreatePartyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SeatLimit   int64  `json:"seat_limit"`
}

type UpdatePartyRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	SeatLimit   *int64  `json:"seat_limit"`
}
