package retigo

// Doer implementations issue commands to a Redis server.
type Doer interface {
	Do(cmd string, args ...interface{}) (interface{}, error)
}

// Get the value of a key.  If the key does not exist, the BulkString
// represents nil.  An error results if the value stored at the key
// is not a string.
func Get(conn Doer, key string) *BulkString {
	return NewBulkString(conn.Do("GET", key))
}

// Set the specified key to the value.
func Set(conn Doer, key string, val interface{}) *SimpleString {
	return NewSimpleString(conn.Do("SET", key, val))
}

// Exists reports how many of the specified keys exist.
func Exists(conn Doer, keys ...string) *Integer {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	return NewInteger(conn.Do("EXISTS", args...))
}

func SetEX(conn Doer, key string, value interface{}, seconds int) *SimpleString {
	return NewSimpleString(conn.Do("SET", key, value, "EX", seconds))
}

func SetPX(conn Doer, key string, value interface{}, milliseconds int) *SimpleString {
	return NewSimpleString(conn.Do("SET", key, value, "PX", milliseconds))
}

// SetNX sets the key to the specified value provided the key does not
// already exist.
func SetNX(conn Doer, key string, value interface{}) *SimpleString {
	return NewSimpleString(conn.Do("SET", key, value, "NX"))
}

// SetXX sets the key to the specified value provided the key already exists.
func SetXX(conn Doer, key string, value interface{}) *SimpleString {
	return NewSimpleString(conn.Do("SET", key, value, "XX"))
}

// HGet Returns the value associated with field in the hash stored at the key.
// If the field or hash stored at the key does not exist, the bulk string
// encodes a nil response.
func HGet(conn Doer, key string, field string) *BulkString {
	return NewBulkString(conn.Do("HGET", key, field))
}

// HSet sets the field in the hash stored at the key to value.
// If the key does not exist, a new key holding a hash is created.
// If the field already exists in the hash, it is overwritten.
// A value of 0 is returned if the hash already contains the field.
// A value of 1 is returned if the hash was created.
func HSet(conn Doer, key string, field string, value interface{}) *Integer {
	return NewInteger(conn.Do("HGET", key, field, value))
}
