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
- go语句不能使用返回值
- `main函数`都是使用`goroutine`启动的,因此每个程序至少运行一个goroutine

## channel
- 允许你将值从一个goroutine发送到另一个goroutine
- 确保在 接收的goroutine 尝试使用该值之前，发送方已经发送该值。   
通过blocking--暂停当前goroutine所有进一步操作实现。  
例如：发送操作阻塞发送方，直到另一个gor在同一channel执行了接收操作。

