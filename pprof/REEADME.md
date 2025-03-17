## bug分析pprof
1. 引入包 `_ "net/http/pprof"`
2. 使用命令分析
```bash
# pprof爬取和分析
	go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=60

# inuse_space：分析应用程序常驻内存的占用情况
	go tool pprof http://127.0.0.1:6060/debug/pprof/heap

# 每个goroutine的使用情况
	go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine

# 锁的装填
	go tool pprof http://127.0.0.1:6061/debug/pprof/mutex

block:
	go tool pprof http://127.0.0.1:6061/debug/pprof/block

trace:
	go tool trace [文件]

#GODEBUG调试
go-debug:
 	GODEBUG=scheddetai=1,schedtrace=1000 go run main.go

# 跟踪gc信息
	GODEBUG=gctrace=1 go run main.go

# 安装gops
	go get github.com/google/gops

# 查看逃逸的变量
	go build -gcflags '-m -l' main.go
```