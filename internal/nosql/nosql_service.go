package nosql

type NoSqlService interface {
	Get(key string, ttl int) (string, error)
	Set(key, value string, ttl int) error
	Delete(key string)
}