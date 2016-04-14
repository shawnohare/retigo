// Package retigo provides typed Redis commands that can be used with
// (generalized) redigo connections.  The function response types follow
// the conventions of both Redis and redigo.
//
// The provided functions return structures representing Redis command return
// types.  These retigo types are essentially thin wrappers around
// redigo command responses.  Their Result methods return the appropriate
// Go types as per the redigo Redis to Go type conversion specifications.
// For example, `BulkString.Result` returns ([]byte, error)
//
// Basic Usage
//
// Basic typed Redigo commands are currently exposed as top-level functions.
// For instance:
//
//  var conn redis.Conn
//  // ... conn is initialized
//  retigo.Set(conn, "key1", "value1")
//  var value []byte
//  value, err := retigo.Get(conn, "key1").Result() // string(value) == "value1" is true
//  // is equivalent to
//  value, err := redis.Bytes(conn.Do("GET", "key1"))
//
// Returning the original redigo response
//
// The original redigo response is returned by the retigo type's Redigo method,
// as follows:
//
//  val, err := retigo.Get(conn, "key1")
//  // is equivalent to
//  val, err := conn.Do("GET", "key1")
//
// Redigo helper functions
//
// All retigo types (the values returned by the package's functions) have
// methods mirroring the basic redigo type assertion helper functions. The
// `Array` type has access to the array response specific helper functions.
// Moreover, since retigo return values respect the Redis to Go type conversions present
// in redigo, redigo's type assertion helper functions can be chained with
// retigo types.  For example,
//
//  i, err := retigo.Get(conn, "key").Int()
//  // is equivalent to
//  i, err := redis.Int(retigo.Get(conn, "key").Redigo())
//  // is equivalent to
//  i, err := redis.Int(conn.Do("GET", "key"))
package retigo
