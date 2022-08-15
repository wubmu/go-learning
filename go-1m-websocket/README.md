## go-1m-websocket

```shell
go run client.go 10
```

### 太多的open files
- 每一个socket代表一个文件描述
- OS需要内存去管理每一个open file
- 内存是有限的
- 可以通过`ulimit`更改最大open files数量

