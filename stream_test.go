package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToNats(t *testing.T) {
	_, err := connectToNats()
	if err != nil {
		t.Errorf("connectToNats() failed: %v", err)
	}
}

func TestCreateKVStream(t *testing.T) {
	nc, err := connectToNats()
	assert.NoError(t, err)
	assert.NotNil(t, nc)

	kv, err := createKVStream(nc, "test")
	assert.NoError(t, err)
	assert.NotNil(t, kv)
}

func TestFillKV(t *testing.T) {
	nc, err := connectToNats()
	assert.NoError(t, err)
	assert.NotNil(t, nc)

	kv, err := createKVStream(nc, "test")
	assert.NoError(t, err)
	assert.NotNil(t, kv)

	err = fillKV(kv, 100000)
	assert.NoError(t, err)
}

func TestGetKVRange(t *testing.T) {
	RANGE := 10000
	RUNS := 10000

	nc, err := connectToNats()
	assert.NoError(t, err)
	assert.NotNil(t, nc)
	kv, err := createKVStream(nc, "test")

	assert.NoError(t, err)
	assert.NotNil(t, kv)

	t.Run("Fill KV", func(t *testing.T) {
		err = fillKV(kv, RANGE)
		assert.NoError(t, err)
	})

	t.Run("Get KV Range", func(t *testing.T) {
		err = getKVRange(kv, RANGE, RUNS)
		assert.NoError(t, err)
	})
}
