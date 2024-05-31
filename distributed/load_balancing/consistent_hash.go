package load_balancing

//https://mp.weixin.qq.com/s?__biz=MzkxMjQzMjA0OQ==&mid=2247484687&idx=1&sn=36befe944baf5a8314f7dc575c21248a 小徐线程编程世界主要来实现一致性hash算法

/**
主要来实现一致性hash算法 相比较于正常hash取模进行运算 来看看一致性hash算法有什么优点呗。 首先就是区分出有状态服务和无状态服务

哈希环：
1. 起点位置0 与终点位置2^32-1 重合
2. 环上每个刻度对应 （0，2^32-1）之间的数值
3. hash环是一个首位相连接的环
4. 扩容和缩容下 如何进行处理比较关键
5. 误差拟合（当节点数量较少时、可能造成数量分配不均衡、应该引入虚拟节点增大样本来减少误差）
6. 每个节点可能权重值不一样，能者多劳 所以在分配虚拟节点数量也不应该一直 带权负载均衡

-----------------------------------------------------------------------------------------------------------------------
一致性hash算法使用于有状态服务并且根据key进行hash（key需要有较好的连续分布、如果分布不均衡也容易造成负载不均衡问题） 如果针对无状态服务不考虑负载策略（轮训算法、负载均衡算法）
一致性hash算法最大的优势，在节点数量变更时，只需要承担局部小范围数据局部迁移
一致性hash算法可以通过引入虚拟节点来提高负载均衡程度，并能很好支持带权分治的诉求  一致性hash算法就是使用mysql分片功能被
*/