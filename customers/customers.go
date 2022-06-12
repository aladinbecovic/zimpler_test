package customers

type TopCustomers []*Customer

type Customer struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

func New() *TopCustomers {
	return &TopCustomers{}
}
