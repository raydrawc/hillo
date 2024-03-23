package protal

import "go.etcd.io/etcd/clientv3"

// 可替换的配置中心
type EtcdSyncer struct {
	Addr   string // 监听地址
	Port   int64  // 监听端口
	User   string //
	Passwd string
	Params map[string]string
}

func(s *EtcdSyncer) New() {
	clientv3.New(clientv3.Config{
	})

}
