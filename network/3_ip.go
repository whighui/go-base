package network

/**
IP（Internet Protocol）：IP是互联网上最为重要的网络层协议之一。它定义了数据包的格式和结构，并负责数据包的寻址、路由和转发。IP协议是基于IP地址的，它为每个连接到互联网的设备分配唯一的IP地址。

ICMP（Internet Control Message Protocol）：ICMP是一种用于在IP网络中传递控制消息的协议。它主要用于网络诊断、错误报告和网络管理。常见的ICMP消息包括ping请求和回应、路由器通告等。

ARP（Address Resolution Protocol）：ARP用于将IP地址转换为对应的物理（MAC）地址。当主机要向另一个主机发送数据时，它需要知道目标主机的MAC地址，而ARP协议就负责解析IP地址和MAC地址之间的对应关系。

RARP（Reverse Address Resolution Protocol）：RARP是与ARP相反的协议，用于将物理地址（MAC地址）转换为对应的IP地址。它主要用于某些特定场景下，如在无磁盘工作站上，根据物理地址获取对应的IP地址。

IPsec（IP Security）：IPsec是一组安全扩展协议，用于在IP网络中提供数据的机密性、完整性和认证功能。它通过加密和认证机制，保护IP数据包的安全性。
*/

/**
网络层 IP 数据报，会限制传输大小；
当IP报文标识 允许分片时，若数据包大小超过MTU、则会进行分片；
当IP报文中的不分段标志置为1时，则不允许IP分段。此时数据包大小若超过MTU，则IP层会将该数据包丢弃，并发送一个ICMP差错报文给源主机。
IP分片和重组都发生在网络层，都由网络层IP协议完成。
IP首部中包含的数据为分片和重新组装提供了足够的信息。
*/
