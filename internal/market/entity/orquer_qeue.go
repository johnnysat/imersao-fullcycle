package entity

type OrderQeue struct {
	Orders []*Order
}

// Método Less
func (oq OrderQeue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

//Método Swap

//Método Len

// Método Push

// Método Pop
