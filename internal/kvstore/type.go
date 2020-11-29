package kvstore

// Service provides operations on key-value store
type Service interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
