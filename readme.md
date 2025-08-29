# example 微服务
# 1. 服务注册 
+ etcd
```shell
etcdctl --user=root:root get --prefix order.rpc
etcdctl --user=root:root get --prefix user.rpc
```
## 安装etcd docker
```shell
 docker pull bitnami/etcd:latest
# ETCD_ROOT_PASSWORD=root
#ALLOW_NONE_AUTHENTICATION=yes
docker run -d --name etcd \
  -p 2379:2379 \
  -p 2380:2380 \
  -e ALLOW_NONE_AUTHENTICATION=yes \
  -e ETCD_ROOT_PASSWORD=root \
  bitnami/etcd:latest 
```

## 服务启动
```shell
go run example/order/main.go
go run example/user/main.go
```
## 关于服务
+ 目前准备写一个简单的购物车支付订单的验证demo
+ 设计的思路不会很复杂，主要验证一下分布式事务、和etcd的使用、验证一下订单的拆单

## grpc相关
+ 生成rpc文件
```shell
goctl rpc protoc {name}.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
```