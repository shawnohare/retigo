# retigo

Package retigo provides typed Redis commands that can be used with
(generalized) redigo connections.  Each Redis command has a redigo Do and
Send version.

The provided functions return structures representing Redis command return
types.  The Go types these structures return via their `Result` methods
respect the conventions in the redigo package.  For example,
`BulkString.Result` returns `([]byte, error)`

## Example (Do)

Suppose a `conn` is a redigo `redis.Conn` instance connected to some Redis
instance that stores the value "value1" in the key "key1".  
```go
var value []byte
value, err := retigo.Get("key1").Result() // string(value) == "value1" is true
```
is equivalent to
```go
var value []byte
value, err := redis.Bytes(redis.Do("GET", "key1")) 
```



