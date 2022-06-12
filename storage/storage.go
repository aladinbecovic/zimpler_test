package storage

type Storage interface {
	SaveData() error
	LoadData() error
	Close() error
}
