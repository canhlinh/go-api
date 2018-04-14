package stores

type MgoStore struct {
}

func NewMgoStore() Store {
	return &MgoStore{}
}
