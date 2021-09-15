package types

// Amount of money in minimum units (diram, kopeyka, cent...)
type Money int64

// Category of payment
type Category string

// Status of payments
type Status string

// Statuses of payments
const (
	StatusOK         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Structure of payment
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
