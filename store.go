package postmi

type Store interface {
	Connect(string) error

	Save(*Post) error
	Get(int64) (*Post, error)
	Delete(int64) error

	GetSlice(int64, int64) ([]*Post, error)
}
