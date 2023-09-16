package tcp

/**
TCP 半连接队列：当服务端接受到客户端第一个syn请求连接包时，服务端为此类连接状态维护一个半连接队列 可以内核配置 tcp_max_syn_backlog 默认是2048

TCP 全连接队列：完成三次握手之后，在服务端 accept 之前, 连接会放到 全连接队列(accept queue). 网络配置somaxconn 默认是128个
全连接队列满之后展现行为：
tcp_abort_on_overflow=0 服务端丢弃客户端发来的ack包，等同于客户端并没有传ack宝，服务端并重传syn+ack（重传次数net.ipv4.tcp_synack_retries可进行配置 linux里边默认是5 ）
tcp_abort_on_overflow=1 服务端直接发送rst包进行关闭连接

https://user-images.githubusercontent.com/1747852/70062414-f0075b80-1620-11ea-8f42-47292f867b59.png 可看看图片了解半连接队列和全连接队列
*/
