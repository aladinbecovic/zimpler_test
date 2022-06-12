package storage

import "github.com/aladinbecovic/zimpler_test/customers"

type Storage interface {
	SaveData(customers *customers.TopCustomers) error
	LoadData(topCustomers *customers.TopCustomers) error
}
