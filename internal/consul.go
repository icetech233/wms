package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func readConsul() []byte {
	t1 := time.Now()
	resp, err := http.Get("http://127.0.0.1:8500/v1/kv/wms")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var entries []*KVPair
	// json.Unmarshal()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&entries)
	if err != nil {
		panic(err)
	}
	if len(entries) == 0 {
		panic("entries len is 0")
	}
	fmt.Printf("请求时间:%d微秒", time.Since(t1).Microseconds())
	return entries[0].Value
}

// KVPair
type KVPair struct {
	// Key is the name of the key. It is also part of the URL path when accessed via the API.
	Key string

	// CreateIndex holds the index corresponding the creation of this KVPair. This is a read-only field.
	CreateIndex uint64

	// ModifyIndex is used for the Check-And-Set operations and can also be fed
	// back into the WaitIndex of the QueryOptions in order to perform blocking
	// queries.
	ModifyIndex uint64

	// LockIndex holds the index corresponding to a lock on this key, if any. This
	// is a read-only field.
	LockIndex uint64

	// Flags are any user-defined flags on the key. It is up to the implementer to check these values, since Consul does not treat them specially.
	Flags uint64

	// Value is the value for the key. This can be any value, but it will be base64 encoded upon transport.
	Value []byte

	// Session is a string representing the ID of the session. Any other interactions with this key over the same session must specify the same session ID.
	Session string
}
