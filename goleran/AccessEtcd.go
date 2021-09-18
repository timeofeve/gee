package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.151.3.131:23790", "10.151.3.131:23791", "10.151.3.131:23792"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// KV接口的实现
	kv := clientv3.NewKV(cli)

	// put 值
	// putResp, err := kv.Put(context.TODO(), "/test/key1", "hello etcd!")
	// fmt.Printf("PutResponse: %v, err: %v", putResp, err)

	// kv.Put(context.TODO(), "/test/key2", "Hello World!")
	// // 再写一个同前缀的干扰项
	// kv.Put(context.TODO(), "/testspam", "spam")

	// // get
	// getResp, err := kv.Get(context.TODO(), "/test/key1")
	// fmt.Printf("getResp: %v \n", getResp)

	// // range get
	// rangeResp, err := kv.Get(context.TODO(), "/test/", clientv3.WithPrefix())
	// fmt.Printf("rangeResp: %v \n", rangeResp)

	// put
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// _, err = cli.Put(ctx, "lmh", "lmh")
	// cancel()
	// if err != nil {
	// 	fmt.Printf("put to etcd failed, err:%v\n", err)
	// 	return
	// }
	// get
	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// resp, err := cli.Get(ctx, "lmh")
	// cancel()
	// if err != nil {
	// 	fmt.Printf("get from etcd failed, err:%v\n", err)
	// 	return
	// }
	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	// }

	// Lease
	lease := clientv3.NewLease(cli)
	// 设置 ttl为10s租约
	grantResp, err := lease.Grant(context.TODO(), 10)

	kv.Put(context.TODO(), "/test/vanish", "vanish in 10s", clientv3.WithLease(grantResp.ID))

	keepResp, err := lease.KeepAliveOnce(context.TODO(), grantResp.ID)
	fmt.Printf("rangeResp: %v \n", keepResp)

}
