package nosql


// ttl is the duration of the record existence in minutes
type NoSqlService interface {
	Get(key string) (string, error)
	Set(key, value string, ttl int) error
	Delete(key string)
}