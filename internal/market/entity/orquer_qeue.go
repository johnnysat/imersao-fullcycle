package entity

type OrderQeue struct {
	Orders []*Order
}

// Método Less
func (oq *OrderQeue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Método Swap
func (oq *OrderQeue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Método Len
func (oq *OrderQeue) Len() int {
	return len(oq.Orders)
}

// Método Push
func (oq *OrderQeue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Método Pop
