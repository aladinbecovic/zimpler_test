package customers

type TopCustomers struct {
	Customers []*Customer
}

type Customer struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int32  `json:"totalSnacks"`
}
