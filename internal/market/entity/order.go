package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	Pedido        string
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transactions
}
