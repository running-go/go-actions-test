# go-actions-test
golang ci test


## 构建和运行


### 运行当前容器
```shell
docker run --name go-action-test -p 10101:10101 docker.io/akenisocean/go-action-test:v0.0.1
```
删除当前启动的容器
```shell
docker stop go-action-test && docker rm go-action-test
```