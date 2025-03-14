package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var datas []string

func init() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}
func main() {
	test1()

}

func test2() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := range 100 {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}
	http.ListenAndServe(":6061", nil)
}

func test1() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-program-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	http.ListenAndServe(":6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
