# 大规模可扩展的web应用需要的特质：
1. 可扩展（对性能进行扩展）
	- 垂直扩展：提升单台设备性能  
	go支持并发编程,一个goweb应用只需要一个操作系统线程。
	- 水平扩展：增加pc数量来提高   
	go可通过多个Go web应用之上假设代理。go web会编译成静态二进制文件，方便统一部署。
2. 模块化
3. 可维护
4. 高性能

# web应用工作原理
## web应用定义
- 这个程序必须向发送命令请求的客户端返回HTML，客户端会向用户展示渲染后的HTML
- 这个程序向客户端传送数据时必须使用HTTP协议
- 若返回的是非HTML数据，那么这个程序就是为其他程序服务的`Web服务`

## HTTP

### 定义
- HTTP是一种无状态，由文本构成的请求-响应协议，使用的是CS的计算模型
- 客户端也被称为用户代理，服务器被称为Web服务器, HTTP客户端很多情况是Web浏览器。

### 一个典型的HTTP请求
#### 如下排列
1. 请求行[请求方法 统一资源标识符(URI) HTTP版本]
2. 零个或多个请求首部（header）
3. 一个空行;
4. 可选的报文主体（body）

#### example
```
GET xxx/xx.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(empty line)
```
### 请求方法
#### 目的
客户端想对资源进行的操作
#### 历程
|版本|方法|
|:----|:----|
|HTTP0.9|GET|
|1.0|POST和HEAD|
|1.1|PUT DELETE OPTIONS TRACE CONNECT|
#### 安全请求方法
- 如果一方法只求服务器返回信息。不会对服务器状态做出修改
- GET HEAD TRACE OPTIONS
#### 幂等请求方法




### 请求首部
#### 目的
- 记录了与请求本身和客户端相关的信息
- 用冒号分割纯文本键值对,最后用回车和换行结尾。

### HTTP响应
#### 组成
1. 一个状态行[状态码 原因短语]
2. 零个或任意数量响应首部
3. 一个空行
5. 一个可选报文主体
#### 响应状态码
|码|描述|
|:----|:----|
|1XX|情报状态码，S方告诉C方，已接到请求并做了处理|
|2XX|成功，已成功接收并处理|
|3XX|重定向|
|4XX|客户端错误,例404,S方无法找到C所想要资源|
|5XX|服务器错误,例500 Ineernal Server Error|

### URI
- 包含了URN（字符串表示资源名字）和URL（字符串表示资源位置）
- URI格式：`<方案名称>:<分层部分>[ ? <查询参数> ][ # <片段> ]`
- 方案：记录正在使用的方案，定义了其余部分的结构，本书使用的是HTTP方案
- 分层：包含资源的识别信息。// / /
- 例子：http://sausheong:password@www.example.com/docs/file?name=sausheong&location=singapore#summary
- URL规定：不能包含空格，？和#具有特殊用途，不能用于其他用途。
- 用URL编码对特殊符号转换。
  1. 获取该字符的ASCII字节值
  2. 转换成一个二位16进制数
  3. 在这个数字前加%
- 例：空格。字节值是32, 16进制是20，经过处理后就是%20

### HTTP/2 协议
- 对性能关注
- 二进制协议
- 完全多路复用，多个请求响应在同一时间用同一个连接。
- 首部压缩
- Go在使用HTTPS时会自动使用HTTP/2

### 总结

### 处理器

#### 功能
- 接收和处理客户端发来的请求
- 调用模板引擎，生成HTML，并把数据填充至将要传给客户端的响应报文中。

#### 性质
- MVC：模型视图控制器模式 Model-View-Controller
- 用MVC模式来讲，既是控制器，也是模型。
- 控制器应苗条， 只包含路由代码和HTTP报文的解包和打包逻辑。
- 模型应丰满。包含应用的逻辑和数据。

### 模板引擎

#### 功能
- HTTP响应报文传给客户端的HTML是模板转换而成的，模板引擎通过模板和数据生成最终的HTML

#### 分类
- 静态模板 
  - 夹杂占位符，引擎将数据替换占位符生成HTML
  - 少量或不含逻辑代码，又称为无逻辑模板。
- 动态模板
  - 包含占位符，一些编程语言结构。 
  - JSP ASP ERB