package main

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"log"
	"time"
)

func init() {
	// Register boltdb store to libkv
	boltdb.Register()
}

func main() {
	client := "./local-kv.db" // ./ appears to be necessary

	// Initialize a new store
	kv, err := libkv.NewStore(
		store.BOLTDB, // or "boltdb"
		[]string{client},
		&store.Config{
			Bucket:            "libnetwork",
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatalf("Cannot create store: %v", err)
	}
	//kv.Delete("docker/network/v1.0/endpoint_count/2436da6c65d238fb667e2bdc4d10d3f77bdced46efae15fb0f518750b87ea664/")
	//kv.Delete("docker/network/v1.0/network/2436da6c65d238fb667e2bdc4d10d3f77bdced46efae15fb0f518750b87ea664/")
	pair, err := kv.List("docker/network")
	for _, p := range pair {
		println("key:", string(p.Key))
		println("value:", string(p.Value))
	}
	//kv.Put("docker/network/v1.0/network/b090008c23828642b21bbb6c36de4ab271651b7bbbda2f11dac05d637492b56d/",[]byte("{\"addrSpace\":\"LocalDefault\",\"attachable\":false,\"configFrom\":\"\",\"configOnly\":false,\"created\":\"2020-07-01T11:37:32.177614796+08:00\",\"enableIPv6\":false,\"generic\":{\"com.docker.network.enable_ipv6\":false,\"com.docker.network.generic\":{\"parent\":\"enp94s0f1\"}},\"id\":\"b090008c23828642b21bbb6c36de4ab271651b7bbbda2f11dac05d637492b56d\",\"inDelete\":false,\"ingress\":false,\"internal\":false,\"ipamOptions\":{},\"ipamType\":\"default\",\"ipamV4Config\":\"[{\\\"PreferredPool\\\":\\\"192.168.2.0/24\\\",\\\"SubPool\\\":\\\"\\\",\\\"Gateway\\\":\\\"192.168.2.1\\\",\\\"AuxAddresses\\\":{\\\"local1\\\":\\\"192.168.2.2\\\"}}]\",\"ipamV4Info\":\"[{\\\"IPAMData\\\":\\\"{\\\\\\\"AddressSpace\\\\\\\":\\\\\\\"LocalDefault\\\\\\\",\\\\\\\"AuxAddresses\\\\\\\":{\\\\\\\"local1\\\\\\\":\\\\\\\"192.168.2.2/24\\\\\\\"},\\\\\\\"Gateway\\\\\\\":\\\\\\\"192.168.2.1/24\\\\\\\",\\\\\\\"Pool\\\\\\\":\\\\\\\"192.168.2.0/24\\\\\\\"}\\\",\\\"PoolID\\\":\\\"LocalDefault/192.168.2.0/24\\\"}]\",\"labels\":{},\"loadBalancerIP\":\"\",\"loadBalancerMode\":\"NAT\",\"name\":\"host\",\"networkType\":\"macvlan\",\"persist\":true,\"postIPv6\":false,\"scope\":\"local\"}"),&store.WriteOptions{IsDir:false,TTL: 10})
	//p, _ :=kv.Get("docker/network/v1.0/network/b090008c23828642b21bbb6c36de4ab271651b7bbbda2f11dac05d637492b56d/")
	//println("key:", string(p.Key))
	//println("value:", string(p.Value))

	kv.Close()

}
