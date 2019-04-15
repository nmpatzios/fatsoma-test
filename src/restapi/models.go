package restapi

// TicketOptions holds the data for tickets options
type TicketOptions struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Allocation  int    `json:"allocation"`
}

// PurchaseTickets holds the data for purchasing tickets
type PurchaseTickets struct {
	Quantity int    `json:"quantity"`
	UserID   string `json:"user_id"`
}
