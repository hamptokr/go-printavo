package printavo

import "time"

type Payment struct {
	Id              int       `json:"id,omitempty"`
	OrderId         int       `json:"order_id,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	Name            string    `json:"name,omitempty"`
	Amount          string    `json:"amount,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
