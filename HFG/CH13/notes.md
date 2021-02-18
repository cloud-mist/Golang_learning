# goroutine & channel

## 定义
- goroutine可以让程序同时处理几个不同的任务，可以使用channel协调它们的工作
- channel允许goroutine互相发送数据并同步

## 什么时候该用
- 做的这个例子，调用三个网站，寻找body长度，但是只能一个接着一个做，
- 等待用户输入
- 等待文件读取等等

## 如何使用
- 函数或方法前面加一个`go`
- `main函数`都是使用`goroutine`启动的,因此每个程序至少运行一个goroutine

