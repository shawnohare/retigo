// Package retigo provides typed Redis commands that can be used with
// (generalized) redigo connections.  The function response types follow
// the conventions of both Redis and redigo.
package retigo

import (
	"github.com/garyburd/redigo/redis"
)

type Conn interface {
	Do(cmd string, args ...interface{}) (interface{}, error)
	// Send(cmd string, args ...interface{}) error
	Close()
}

type Pool interface {
	Get() Conn
}

// PoolDoer implements a Doer interface from an underlying connection pool.
// Instances can be used with this package's typed redigo commands.
type PoolDoer struct {
	Pool Pool
}

// Do fetches a new connection from the underlying pool, calls its Do method,
// and closes the connection.
func (p PoolDoer) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

// Type represents a general Redis response type, as returned by redigo.
// All retigo types embed the Type structure and have access to its methods.
type Type struct {
	val interface{}
	err error
}

// Integer wraps the redigo response for a Redis function with integer return type.
type Integer struct {
	*Type
}

// SimpleString wraps the redigo response for a Redis function with simple string return type.
type SimpleString struct {
	*Type
}

// BulkString wraps the redigo response for a Redis function with  bulk string return type.
type BulkString struct {
	*Type
}

// Array wraps the redigo response for a Redis function with  array string return type.
type Array struct {
	*Type
}

// Redigo returns the original redigo command response the type represents.
// For example, conn.Do("GET", "key") is equivalent to
// Get(conn, "key").Redigo(), where conn is some redis.Conn instance.
func (t *Type) Redigo() (interface{}, error) {
	if t == nil {
		return nil, nil
	}
	return t.val, t.err
}

// Bool applies the redigo helper to the the retigo type.
// Bool is a helper that converts a command reply to a boolean. If err is not
// equal to nil, then Bool returns false, err. Otherwise Bool converts the
// reply to boolean as follows:
//
//  Reply type      Result
//  integer         value != 0, nil
//  bulk string     strconv.ParseBool(reply)
//  nil             false, ErrNil
//  other           false, error
func (t *Type) Bool() (bool, error) {
	return redis.Bool(t.Redigo())
}

// Bytes applies the redigo helper to the retigo type.
// Bytes is a helper that converts a command reply to a slice of bytes. If err
// is not equal to nil, then Bytes returns nil, err. Otherwise Bytes converts
// the reply to a slice of bytes as follows:
//
//  Reply type      Result
//  bulk string     reply, nil
//  simple string   []byte(reply), nil
//  nil             nil, ErrNil
//  other           nil, error
func (t *Type) Bytes() ([]byte, error) {
	return redis.Bytes(t.Redigo())
}

// Float64 applies the redigo helper to the retigo type.
// Float64 is a helper that converts a command reply to 64 bit float. If err is
// not equal to nil, then Float64 returns 0, err. Otherwise, Float64 converts
// the reply to an int as follows:
//
//  Reply type    Result
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func (t *Type) Float64() (float64, error) {
	return redis.Float64(t.Redigo())
}

// Int applies the redigo helper to the retigo type.
// Int is a helper that converts a command reply to an integer. If err is not
// equal to nil, then Int returns 0, err. Otherwise, Int converts the
// reply to an int as follows:
//
//  Reply type    Result
//  integer       int(reply), nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func (t *Type) Int() (int, error) {
	return redis.Int(t.Redigo())
}

// Int64 applies the redigo helper to the retigo type.
// Int64 is a helper that converts a command reply to 64 bit integer. If err is
// not equal to nil, then Int returns 0, err. Otherwise, Int64 converts the
// reply to an int64 as follows:
//
//  Reply type    Result
//  integer       reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func (t *Type) Int64() (int64, error) {
	return redis.Int64(t.Redigo())
}

// Uint64 applies the redigo helper to the retigo type.
// Uint64 is a helper that converts a command reply to 64 bit integer. If err is
// not equal to nil, then Int returns 0, err. Otherwise, Int64 converts the
// reply to an int64 as follows:
//
//  Reply type    Result
//  integer       reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func (t *Type) Uint64() (uint64, error) {
	return redis.Uint64(t.Redigo())
}

// String applies the redigo helper to the retigo type.
// String is a helper that converts a command reply to a string. If err is not
// equal to nil, then String returns "", err. Otherwise String converts the
// reply to a string as follows:
//
//  Reply type      Result
//  bulk string     string(reply), nil
//  simple string   reply, nil
//  nil             "",  ErrNil
//  other           "",  error
func (t *Type) String() (string, error) {
	return redis.String(t.Redigo())
}

// Result returns a typed redigo command response.
func (t *Integer) Result() (int64, error) {
	return t.Int64()
}

// Result returns a typed redigo command response.
func (t *SimpleString) Result() (string, error) {
	return t.String()
}

// Result returns a typed redigo command response.
func (t *BulkString) Result() ([]byte, error) {
	return t.Bytes()
}

// Result returns a typed redigo command response.
func (t *Array) Result() ([]interface{}, error) {
	return redis.Values(t.Redigo())
}

// Define array type specific helper methods.
func (t *Array) ByteSlices() ([][]byte, error) {
	return redis.ByteSlices(t.Redigo())
}

// Ints applies the redigo helper function to array responses.
func (t *Array) Ints() ([]int, error) {
	return redis.Ints(t.Redigo())
}

// Strings applies the redigo helper function to array responses.
func (t *Array) Strings() ([]string, error) {
	return redis.Strings(t.Redigo())
}

// StringMap applies the redigo helper function to array responses.
func (t *Array) StringMap() (map[string]string, error) {
	return redis.StringMap(t.Redigo())
}

// IntMap applies the redigo helper function to array responses.
func (t *Array) IntMap() (map[string]int, error) {
	return redis.IntMap(t.Redigo())
}

// Int64Map applies the redigo helper function to array responses.
func (t *Array) Int64Map() (map[string]int64, error) {
	return redis.Int64Map(t.Redigo())
}

func NewType(value interface{}, err error) *Type {
	return &Type{val: value, err: err}
}

func NewSimpleString(value interface{}, err error) *SimpleString {
	return &SimpleString{Type: NewType(value, err)}
}

func NewBulkString(value interface{}, err error) *BulkString {
	return &BulkString{Type: NewType(value, err)}
}

func NewInteger(value interface{}, err error) *Integer {
	return &Integer{Type: NewType(value, err)}
}

func NewArray(value interface{}, err error) *Array {
	return &Array{Type: NewType(value, err)}
}
