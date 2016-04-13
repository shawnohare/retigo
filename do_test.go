package retigo

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("GET", "key").Expect("value")

	v, err := Get(conn, "key").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, string(v), "value")
}

func TestGetKeyWithNonStringValue(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("GET", "key").Expect(redis.Error("key has non-string value"))

	v, err := Get(conn, "key").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.Error(t, err)
	assert.Nil(t, v)
}

func TestExistsKeyExists(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("EXISTS", "key").Expect(int64(1))

	v, err := Exists(conn, "key").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v)
}

func TestExistsKeyDoesNotExist(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("EXISTS", "key").Expect(int64(0))

	v, err := Exists(conn, "key").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.NoError(t, err)
	assert.Equal(t, int64(0), v)
}

func TestExistsMultiple(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("EXISTS", "key1", "key2", "key3").Expect(int64(2))

	v, err := Exists(conn, "key1", "key2", "key3").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.NoError(t, err)
	assert.Equal(t, int64(2), v)
}

func TestSet(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("SET", "key", "value").Expect("OK")

	v, err := Set(conn, "key", "value").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.NoError(t, err)
	assert.Equal(t, v, "OK")
}

func TestGetNoKey(t *testing.T) {
	conn := redigomock.NewConn()
	cmd := conn.Command("GET", "key").Expect(nil)

	v, err := Get(conn, "key").Result()
	if conn.Stats(cmd) != 1 {
		t.Fatal("Command was not called.")
	}
	assert.Error(t, err) // there should be a "got nil" error
	assert.Nil(t, v)
}
