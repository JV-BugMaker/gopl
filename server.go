package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
// 并发对资源加锁 竞态条件
var mu sync.Mutex
var count int
func main() {
	http.Handle("/",handler)
	http.Handle("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	//加锁
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}

func counter(w http.ResponseWriter,w *http.Request){
	//加锁操作
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n",count)
	mu.Unlock()
}

