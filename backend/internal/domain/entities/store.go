package entities

type NewStore struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

type ExistingStore struct {
	Id int64
}

type ExistingStoreWithAddress struct {
	Id      int64
	Address string
}
