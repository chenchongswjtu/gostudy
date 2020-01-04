package main

// golang中的空结构体 channel := make(chan struct{})
// 特点
// 省内存，尤其在事件通信的时候。
// struct零值就是本身，读取close的channel返回零值
// 使用场景
// 首先事件通知，可以通过写入 通知其他协程，但是只能通知一个。
// 和close进行配合，通知所有相关协程。

//channel := make(chan struct{})
//go func() {
//	// ... do something
//	channel <- struct{}{}
//}()
//fmt.Println(<-channel)

//type Server struct {
//	serverStopChan chan struct{}
//	stopWg         sync.WaitGroup
//}
//func (s *Server) Stop() {
//	if s.serverStopChan == nil {
//		panic("gorpc.Server: server must be started before stopping it")
//	}
//	close(s.serverStopChan)
//	s.stopWg.Wait()
//	s.serverStopChan = nil
//}
//func serverHandler(s *Server){
//	for {
//		select {
//		case <-s.serverStopChan:
//			return
//		default:
//			// .. do something
//		}
//	}
//}

func main() {

}
