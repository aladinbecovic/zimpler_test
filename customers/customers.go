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

func (tc *TopCustomers) SortTopCustomers() {
	for i := 0; i < len(*tc); i++ {
		min_idx := i
		for j := i + 1; j < len(*tc); j++ {
			if (*tc)[i].TotalSnacks < (*tc)[j].TotalSnacks {
				min_idx = j

				tmp := (*tc)[i]
				(*tc)[i] = (*tc)[min_idx]
				(*tc)[min_idx] = tmp
			}
		}
	}
}
