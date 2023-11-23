package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	// Get a new client
	client, err := api.NewClient(&api.Config{
		Address: "http://localhost:8500",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "my-svr", Value: []byte("1000"), Flags: 32}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("my-svr", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

func TestHttpReg(t *testing.T) {
	svrID := HttpReg("mysvr-rpc", "localhost", 3456)
	log.Printf("服务注册成功：%v\n", svrID)
	for true {
		// 制定一个定时器，模拟20s后摘除服务
		t := time.NewTicker(20 * time.Second)
		select {
		case tt := <-t.C:
			HttpUnReg(svrID)
			log.Printf("服务卸载，tt:%v\n", tt)
			return
		}
	}
}
