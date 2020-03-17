# plog
golang 基于zap库封装的log包

#### 安装：
```go get github.com/jageros/plog```
#### 使用：

```go
//初始化, 不调用此函数可默认初始化为“debug”模式
plog.InitPlogConfig("debug")
//打印日志
plog.Infof("xxxxxx%d", 404)
```