package gog_atu

type OrderPage struct {
	TotalPages
	Orders []Order `json:"orders"`
}

func (op *OrderPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, ord := range op.Orders {
		ord := ord
		idGetters = append(idGetters, &ord)
	}
	return idGetters
}
