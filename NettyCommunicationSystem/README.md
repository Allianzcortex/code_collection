# NettyCommunicationSystem
一个基于 Netty 的聊天程序(包含客户端与服务器端)

#### 设计思想

Netty 是一个 **异步**(Asynchronous) 和 **事件驱动**(event-driven) 的高并发服务器开发框架。


#### 开发实例

在 [《Netty In Action》](https://www.manning.com/books/netty-in-action) 书的第 12 章实现了一个用 WebSocket 进行 server 和 client 端进行聊天的程序。

主要是为了熟悉 Netty 的 Pipeline/ChannelHandler 以及服务器端编程。基于不同的思路进行了实现：

- 用一个 static 的 HashMap 来存储每次连接时所创建的 `<远程端口，ChannelHandlerCortex>`,并在用户选择聊天后根据对应的 ctx 来发送信息 [√]

- 用 redis 来存储，方便实现 expire 登陆有效期等功能 [×]

- 基于 ChannelLogger 进行日志记录 [×]

![first_demo](/first_demo.png)

