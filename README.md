go-base 是针对whig的学习记录，主要分为如下章节

# benchmark 性能测试



## redis单机QPS



## mysql单机QPS



## tcp长连接



## websocket&&http





# design 系统设计



## feed流设计

```
# 设计feed流
朋友圈，微博，b 站抖音的推送，这些常用 app 的使用场景有一个共同的名词：<br>
是将 “N个发布者的信息单元” 通过 “关注关系” 传送给 “M个接收者”。

# 博客 https://blog.csdn.net/liyifan687/article/details/123090090

# https://zhuanlan.zhihu.com/p/259562762


# Feed流整体服务架构设计

# Feed流产品设计

1. 微信
2. 微博
3. 
```

## friends朋友圈设计

## rank排行榜设计

```
# 设计一个分布式排行榜

## redis本身实现通过zset来实现

## tair阿里开源KV键值对 多维度来进行rank
1. 使用C++来写的 好难受


3. 高并发排行榜  zset做排序 redis做分布式存储 在最后做归并排序 每日排行、周排行、月排行、历史排行  用户王者荣耀、好友排行、段位排行
```

## 长短URL设计

```
https://juejin.cn/post/6844904090602848270  相关博客设计的不错
```



## TCE部署平台



# distribute 分布式



## cap

CAP

CAP学术理论 
1. C ： 数据一致性 指数据能一起变化 能让数据整齐化一 : 写请求会导致变化 需要同步到各个节点 读请求来验证数据是否是一致性
2. A ： 可用性    它要求系统内的节点们接收到了无论是写请求还是读请求，都要能处理并给回响应结果 即便发生宕机了其他节点也可以正常返回
3. P ： 分区容错性 分布式系统中存在多个节点进行网络通信，节点与节点之间出现通信问题 数据版本不一致那么就会发生分区，但是即便发生了分区，分布式系统还要继续运行。这也是必须要选P的原因

CAP缺点

工程实践

MySQL是CP还是AP
既可以是CP也可以是AP，基于Mysql主从复制的配置
1. 完全同步 CP
2. 半同步   CP
3. 异步     AP

ZooKeeper是CP还是AP
CP
1. zooKeeper更倾向于CP   ZooKeeper 使用了一个基于 Paxos 算法的原子广播协议来实现数据的一致性
2. 当客户端向 ZooKeeper 发送写操作时，ZooKeeper 会将该操作转换为一个事务，并通过 Paxos 算法确保该事务在整个集群中的多数节点上达成一致。只有在大多数节点确认接受该事务后，ZooKeeper 才会将其应用到数据存储中，从而保证数据的一致性。
AP 
1. ZooKeeper 在一些特定的操作和场景下也可以提供一定程度的可用性和分区容忍性（AP），
2. 例如读取操作可以在不同节点之间进行负载均衡，以提高系统的读取性能和可用性。此外，ZooKeeper 还支持有序性和临时节点等功能，用于实现分布式锁、选举和事件通知等应用。

Redis是CP还是AP
AP 
1. redis还是更倾向于AP ，Redis 的主要目标是提供高性能、低延迟和高可用性的数据存储和缓存解决方案。
2. Redis主从复制来提高可用性
3. Redis并不保证强一致性 在主从复制过程中 数据复制是异步的，从服务器上可以存在一定数据同步延迟 
4. 故障切换可能会丢失数据
5. redis还提供sharding分片功能 可以将数据分散存储到以提高系统的扩展性和
  异步复制和数据延迟导致redis不能够保持数据一致性



## limit限流

### 单机限流

### 分布式限流



## 负载均衡

### 一致性hash



### nginx负载均衡算法



## lock锁



### redis 分布式锁

### etcd 分布式锁





# ds 数据结构



## cache缓存

1. lru
2. lru-k
3. lru-timing
4. lru-ttl
5. lru-rocksdb
6. lru-redis
7. lru-mysql

知识体系里边存在上述lru缓存设计

```
# cache设计
1. cache与cpu cacheline 运行体系绑定  优先page-cache读
2. 实现insert链路中的无锁  多线程更新情况下要保持cache数据结构更新的原子性 不可中断
3. 不影响性能的基础上对cache的内存控制  
4. 如何让频繁的内存分配和释放不是cache的性能瓶颈 key-value大小差异非常大，几个byte 到 几 M；频繁的插入和Erase）能够对Cache拿到的内存有一个高效的管理，在对分配器底层有足够了解的情况下 选择合适的分配器能够起到加速cache使用 的作用。
参考链接 https://blog.csdn.net/Z_Stand/article/details/118832473

# 个人设计
1. cpu cache line 目前大小为64byte  cache line是cpu与内存通信的最小单位
2. 每个core 有自己独享的L1 cache L2 cache  MESI协议保证了cache line的一致性
```



## map哈希

1. java hashmap
2. golang hmap
3. sort map 
4. 插入有序map
5. 插入key有序map
6. 并发map



## B树



## 跳表







# linux

CPU有尖刺状态 什么情况导致的 如何优化 如何进行解决 优化数据结构  LRU长度缩短点 利用多级缓存

1. 64位机器字长8字节 为啥
2. Linux下寻址范围2^48为啥

操作系统：分配内存 golang、c++、c  malloc分配内存  熟悉内存组成
使用两种方式：阈值
大于128K 在mmap区进行分配内存 mmap文件映射区 (free函数之后 mmap可以主动释放、 释放之后交给内核)
小于128K brk函数在堆上进行分配  堆上不能主动释放（回收机制：堆上高地址内存没被释放、地址值内存地址也没被释放（不能给操作系统回收掉）  导致堆上有内存碎片 在申请可以直接使用 可以重复利用 ）

## CPU负载过高如何排查





## CPU Cache 多级缓存



# docker

**docker启动Mysql**

docker pull mysql/mysql-server  拉去最新mysql-server适配mac m1版本

1. 启动 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql/mysql-server
2. -v /Users/byte/Desktop/docker/mysql/data:/var/lib/mysql  -v /Users/byte/Desktop/docker/mysql/log:/logs   -v /Users/byte/Desktop/docker/mysql/conf:/etc/mysql/conf.d
3. 上述 -v 是Mysql数据映射到主机目录上 可以这么理解呗
4. 进入到mysql容器内部 docker exec -it mysql bash
   https://www.jianshu.com/p/eb3d9129d880  参考这个博客 基本上就是没有错误

**docker启动redis**

1. 启动redis: docker run -p 6379:6379 --name redis -v /Users/bytedance/Desktop/whig/data/redis/data:/data -v /Users/bytedance/Desktop/whig/data/redis/redis.conf:/etc/redis/redis.conf -d redis redis-server /etc/redis/redis.conf
2. 参考博客 https://blog.csdn.net/u013868195/article/details/119778589


# 直播数字人
1. 直播间 推流拉流 
2. 虚拟人
3. tts
4. 弹幕、评论、场控  内容安全
5. 前端
