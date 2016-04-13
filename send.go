package retigo

// Sender implementations send commands to the connection's output buffer.
type Sender interface {
	Send(cmd string, args ...interface{}) error
}

// SendGet sends a Get command to the connection's output buffer.
func SendGet(conn Sender, key string) error {
	return conn.Send("GET", key)
}

// SendSet sends a Set command to the connection's output buffer.
func SendSet(conn Sender, key string, value interface{}) error {
	return conn.Send("SET", key, value)
}

// SendSetEX sends a SetEx command to the connection's output buffer.
func SendSetEX(conn Sender, key string, value interface{}, seconds int) error {
	return conn.Send("SET", key, value, "EX", seconds)
}

// SendSetPX sends a SetPX command to the connection's output buffer.
func SendSetPX(conn Sender, key string, value interface{}, milliseconds int) error {
	return conn.Send("SET", key, value, "PX", milliseconds)
}

// SendGet sends a Get command to the connection's output buffer.
func SendSetNX(conn Sender, key string, value interface{}) error {
	return conn.Send("SET", key, value, "NX")
}
