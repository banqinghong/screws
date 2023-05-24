package betcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient() (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EndPoint},
		DialTimeout: time.Duration(5) * time.Second,
		Username:    UserName,
		Password:    Password,
	})
	return client, err
}

func GetKVWithPrefix(key string) {
	cli, err := NewEtcdClient()
	if err != nil {
		fmt.Printf("new client error: %s\n", err)
		return
	}
	kvs, err := cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("query kvs with prefix error: %s\n", err)
		return
	}

	for _, kv := range kvs.Kvs {
		fmt.Println("key: ", string(kv.Key))
		fmt.Println("value: ", string(kv.Value))
	}

	// jsonStr, _ := json.Marshal(kvs.Kvs)
	// fmt.Println("result: ", string(jsonStr))
}

func SetKV(key, value string) {
	cli, err := NewEtcdClient()
	if err != nil {
		fmt.Printf("new client error: %s\n", err)
		return
	}
	kvs, err := cli.Put(context.Background(), key, value)
	if err != nil {
		fmt.Printf("put kvs error: %s\n", err)
		return
	}

	fmt.Println("successful: ", kvs.PrevKv.String())
	// jsonStr, _ := json.Marshal(kvs.Kvs)
	// fmt.Println("result: ", string(jsonStr))
}
