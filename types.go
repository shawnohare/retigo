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

type redigoResult struct {
	val interface{}
	err error
}

// Bool wraps the redigo response for a Redis function with boolean return type.
type Bool struct {
	redigo *redigoResult
}

// Integer wraps the redigo response for a Redis function with integer return type.
type Integer struct {
	redigo *redigoResult
}

// SimpleString wraps the redigo response for a Redis function with simple string return type.
type SimpleString struct {
	redigo *redigoResult
}

// BulkString wraps the redigo response for a Redis function with  bulk string return type.
type BulkString struct {
	redigo *redigoResult
}

// Array wraps the redigo response for a Redis function with  array string return type.
type Array struct {
	redigo *redigoResult
}

func (r *redigoResult) result() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.val, r.err
}

func newRedigoResult(val interface{}, err error) *redigoResult {
	return &redigoResult{val: val, err: err}
}

// Redigo returns the original redigo command response.
func (r *Integer) Redigo() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.redigo.result()
}

// Redigo returns the original redigo command response.
func (r *Bool) Redigo() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.redigo.result()
}

// Redigo returns the original redigo command response.
func (r *SimpleString) Redigo() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.redigo.result()
}

// Redigo returns the original redigo command response.
func (r *BulkString) Redigo() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.redigo.result()
}

// Redigo returns the original redigo command response.
func (r *Array) Redigo() (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	return r.redigo.result()
}

// Result returns a typed redigo command response.
func (r *Integer) Result() (int64, error) {
	return redis.Int64(r.redigo.result())
}

// Result returns a typed redigo command response.
func (r *Bool) Result() (bool, error) {
	return redis.Bool(r.redigo.result())
}

// Result returns a typed redigo command response.
func (r *SimpleString) Result() (string, error) {
	return redis.String(r.redigo.result())
}

// Result returns a typed redigo command response.
func (r *BulkString) Result() ([]byte, error) {
	return redis.Bytes(r.redigo.result())
}

// Result returns a typed redigo command response.
func (r *Array) Result() ([]interface{}, error) {
	return redis.Values(r.redigo.result())
}

func NewSimpleString(val interface{}, err error) *SimpleString {
	return &SimpleString{redigo: newRedigoResult(val, err)}
}

func NewBulkString(val interface{}, err error) *BulkString {
	return &BulkString{redigo: newRedigoResult(val, err)}
}

func NewInteger(val interface{}, err error) *Integer {
	return &Integer{redigo: newRedigoResult(val, err)}
}

func NewBool(val interface{}, err error) *Bool {
	return &Bool{redigo: newRedigoResult(val, err)}
}

func NewArray(val interface{}, err error) *Array {
	return &Array{redigo: newRedigoResult(val, err)}
}
