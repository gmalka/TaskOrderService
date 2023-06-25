package nosql


// ttl is the duration of the record existence in minutes
type NoSqlService interface {
	Get(username string, page int) ([]byte, error)
	Set(username string, page int, val []byte) error
	Delete(username string)
}