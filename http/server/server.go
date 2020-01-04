package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//v1
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write([]byte("httpserver v1"))
//	})
//	http.HandleFunc("/bye", sayBye)
//	log.Println("Starting v1 server ...")
//	log.Fatal(http.ListenAndServe(":1210", nil))
//}
//
//func sayBye(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("bye bye ,this is v1 httpServer"))
//}

// v2
//func main() {
//	mux := http.NewServeMux()
//	mux.Handle("/", &myHandler{})
//	mux.HandleFunc("/bye", sayBye)
//
//	log.Println("Starting v2 httpserver")
//	log.Fatal(http.ListenAndServe(":1210", mux))
//}
//
//type myHandler struct{}
//
//func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("this is version 2"))
//}
//func sayBye(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("bye bye ,this is v2 httpServer"))
//}

// v3
//func main() {
//	mux := http.NewServeMux()
//	mux.Handle("/", &myHandler{})
//	mux.HandleFunc("/bye", sayBye)
//
//	server := &http.Server{
//		Addr:         ":1210",
//		WriteTimeout: time.Second * 3, //设置3秒的写超时
//		Handler:      mux,
//	}
//	log.Println("Starting v3 httpserver")
//	log.Fatal(server.ListenAndServe())
//}
//
//type myHandler struct{}
//
//func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	_, _ = w.Write([]byte("this is version 3"))
//}
//
//func sayBye(w http.ResponseWriter, r *http.Request) {
//	// 睡眠4秒  上面配置了3秒写超时，所以访问 “/bye“路由会出现没有响应的现象
//	//time.Sleep(4 * time.Second)
//	_, _ = w.Write([]byte("bye bye ,this is v3 httpServer"))
//}

//在go1.8中新增了一个新特性，利用Shutdown(ctx context.Context) 优雅地关闭http服务。
//文档中描述:
//Shutdown 将无中断的关闭正在活跃的连接，然后平滑的停止服务。处理流程如下：
//
//首先关闭所有的监听;
//然后关闭所有的空闲连接;
//然后无限期等待连接处理完毕转为空闲，并关闭;
//如果提供了 带有超时的Context，将在服务关闭前返回 Context的超时错误;

// 主动关闭服务器
var server *http.Server

func main() {
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server = &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting v3 httpserver")
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("this is version 3"))
}

// 关闭http
func sayBye(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("bye bye ,shutdown the server")) // 没有输出
	err := server.Shutdown(nil)
	if err != nil {
		log.Println("shutdown the server err")
	}
}
