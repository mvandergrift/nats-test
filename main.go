package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

}

func connectToNats() (*nats.Conn, error) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func createKVStream(nc *nats.Conn, store string) (jetstream.KeyValue, error) {
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	kv, err := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: store,
		// Description: fmt.Sprintf("Finite Store for %s", store),
		// TTL:         c.v.GetDuration("NATS.JobResponseKVTTLMinutes") * time.Minute,
		// MaxBytes:    c.v.GetInt64("NATS.JobResponseKVMaxGB") * 1024 * 1024 * 1024,
		// Replicas:    c.v.GetInt("NATS.JobResponseKVReplicas"),
	})
	if err != nil {
		return nil, err
	}
	return kv, nil
}

func fillKV(kv jetstream.KeyValue, size int) error {
	ctx := context.Background()

	value := make([]byte, 1024*1024)
	rand.Read(value)

	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key-%d", i)
		_, err := kv.Put(ctx, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func getKVRange(kv jetstream.KeyValue, size int, runs int) error {
	ctx := context.Background()

	for i := 0; i < runs; i++ {
		key := fmt.Sprintf("key-%d", rand.Intn(size))
		_, err := kv.Get(ctx, key)
		if err != nil {
			return err
		}
	}

	return nil
}
