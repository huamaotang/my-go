# Goroutines 和Channels
## 并发程序
- 同时进行多个任务的程序

## Go语言中的并发程序可以用两种手段来实现
### CSP（communicating sequential process）顺序通信进程
- 一种现代的并发编程模型，在这种编程模型中，值会在不同的运行实例（goroutine）中传递，尽管大多数情况下仍然是被限制在单一实例中
- goroutine和channel支持CSP    

### 多线程共享内存
