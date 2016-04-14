package retigo

// import "github.com/garyburd/redigo/redis"

// Conn extends the redis.Conn interface to add typed commands.
// TODO: finish adding methods.
// type Conn struct {
// 	redis.Conn
// }

// // Pool defines a Get
// type Pool struct {
// 	RedigoPool *redis.Pool
// }

// func (c *Conn) Get(key string) *BulkString {
// 	return Get(c, key)
// }

// func NewConn(conn redis.Conn) *Conn {
// 	return &Conn{Conn: conn}
// }

// func NewPool(pool *redis.Pool) *Pool {
// 	return &Pool{RedigoPool: pool}
// }

// func (p *Pool) Get() *Conn {
// 	return NewConn(p.RedigoPool.Get())
// }

// // UpgradeConn adds typed methods to a redis.Conn interface.  This is an
// // alias for NewConn.
// func UpgradeConn(conn redis.Conn) *Conn {
// 	return NewConn(conn)
// }
