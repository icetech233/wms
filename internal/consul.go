package internal

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func readConsul() []byte {

	consul_config := &api.Config{
		Address: "127.0.0.1:8500",
		Scheme:  "http",
	}
	// Get a new client
	client, err := api.NewClient(consul_config)
	if err != nil {
		panic(err)
	}
	// Get a handle to the KV API
	kv := client.KV()
	// (*api.KVPair, *api.QueryMeta, error)
	pair, q, err := kv.Get("wms", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("请求时间:", q.RequestTime)
	//fmt.Printf("KV: %v \n**\n%s\n", pair.Key, pair.Value)
	return pair.Value
}
