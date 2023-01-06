package validations

type Settlement struct {
	Customer   Customer   `json:"customer"`
	ResolvedBy ResolvedBy `json:"resolved_by"`
	Service    Service    `json:"service"`
}

type Customer struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type ResolvedBy struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type Service struct {
	Id     int    `json:"id"`
	RoomId string `json:"room_id"`
}
