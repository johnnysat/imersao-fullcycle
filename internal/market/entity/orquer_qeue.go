package entity

type OrderQeue struct {
	Orders []*Order
}

// Método Less
func (oq OrderQeue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Método Swap
func (oq OrderQeue) Swap(i, j int) bool {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

//Método Len

// Método Push

// Método Pop
