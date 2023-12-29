# go-actions-test
golang ci test


## 构建和运行

### docker容器运行当前容器
- 构建当前容器
```shell
./docker-build.sh
```

- 运行当前容器
```shell
docker run --name go-action-test -p 10101:10101 docker.io/akenisocean/go-action-test:v0.0.2
```
- 删除当前启动的容器
```shell
docker stop go-action-test && docker rm go-action-test
```

### helm charts k8s方式部署
- 部署
```shell
helm install go-actions-test charts/go-actions-test
```
- 卸载
```shell
helm uninstall go-actions-test
```
