package common

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdClient 封装了连接逻辑，返回一个标准的 Etcd 客户端
// 这样 Master 和 Worker 都可以复用这段代码
func NewEtcdClient(endpoints []string) (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}
	return clientv3.New(config)
}
