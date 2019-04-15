package restapi

import (
	"database/postgres"
	"sync"
)

var wg sync.WaitGroup

// CreateTicketOption creates tickets and return the created ones for the given id
func CreateTicketOption(ticketOption *TicketOptions) (*TicketOptions, error) {
	db, err := postgres.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO ticket_options (name, "desc", allocation,created_at,updated_at) VALUES ( $1, $2, $3,NOW(),NOW()) RETURNING id`
	var id string
	err = db.QueryRow(sqlStatement, ticketOption.Name, ticketOption.Description, ticketOption.Allocation).Scan(&id)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(`SELECT id, name, "desc", allocation  FROM ticket_options where id = $1`, id).Scan(&ticketOption.ID, &ticketOption.Name, &ticketOption.Description, &ticketOption.Allocation)
	if err != nil {
		return nil, err
	}

	return ticketOption, nil
}

// GetTicketOptionByID return the tickets for the given id
func GetTicketOptionByID(id string) (TicketOptions, error) {
	db, err := postgres.Connect()
	if err != nil {
		return TicketOptions{}, err
	}
	defer db.Close()
	var ticketOption TicketOptions

	wg.Add(1)

	go func(id string) {
		defer wg.Done()

		db.QueryRow(`SELECT id, name, "desc", allocation  FROM ticket_options where id = $1`, id).Scan(&ticketOption.ID, &ticketOption.Name, &ticketOption.Description, &ticketOption.Allocation)

	}(id)

	// wait until all conections are open
	wg.Wait()

	return ticketOption, nil
}
