# pprof爬取和分析
profile:
	go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=60

# inuse_space：分析应用程序常驻内存的占用情况
heap:
	go tool pprof http://127.0.0.1:6060/debug/pprof/heap

# 每个goroutine的使用情况
goroutine:
	go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine

# 锁的装填
mutex:
	go tool pprof http://127.0.0.1:6061/debug/pprof/mutex

block:
	go tool pprof http://127.0.0.1:6061/debug/pprof/block

trace:
	go tool trace [文件]

#GODEBUG调试
go-debug:
 	GODEBUG=scheddetai=1,schedtrace=1000 go run main.go
# 跟踪gc信息
go-gc:
	GODEBUG=gctrace=1 go run main.go
# 安装gops
go-ps:
	go get github.com/google/gops

# 查看逃逸的变量
go-gc-escape:
	go build -gcflags '-m -l' main.go