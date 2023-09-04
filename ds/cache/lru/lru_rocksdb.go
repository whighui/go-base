package lru

//基于c++实现的rocksdb实现的分布式kv存储 参考其中lru缓存设计理念

//1.上锁操作 看看锁的粒度是否可以减少
//2.分shard 能力 类似于java 并发map
//3.内存控制管理能力 优先low_pool删除 其次high_pool删除
