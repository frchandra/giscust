package validations

type AgentListResponse struct {
	Data []Agent `json:"data"`
}

type Agent struct {
	CurrentCustomerCount int    `json:"current_customer_count"`
	Email                string `json:"email"`
	Id                   int    `json:"id"`
	IsAvailable          bool   `json:"is_available"`
	Name                 string `json:"name"`
}
