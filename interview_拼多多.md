# 一面 23年9月3号
1. 项目介绍 团队职责 
2. 项目中认为什么地方困难点

MySQL：
1. 聚族索引和非聚族索引区别
2. MySQL select update两个操作实现要求操作同一条记录   用乐观锁和悲观锁两种方式进行实现
3. select ... for update 如果记录不存在是否进行上锁

Kafka:
1. 如何保证消息投递的可靠性  ack+retry
2. broken主节点宕机了 怎么办

Redis:
1. 数据结构 balabala 沉默了五秒钟
2. 疯狂输出sds底层数据结构 length长度 3bit 1byte 2byte 4byte 8byte字节的数据结构 并讲述packet禁止内存对齐 面试官不在在redis上停留 还的是疯狂输出比较得劲

java：
1. java和golang区别
2. spring aop 解释一下，框架是怎么实现的 (参考gorm钩子函数)

other:
1. 乐观锁和悲观锁区别 

算法：
1. 两个链表，判断相同元素
2. 二叉树，根节点上的路径和
