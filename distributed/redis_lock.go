package distributed

//普通锁机制  CAS+自旋  阻塞/唤醒机制

//分布式锁 适用于多机器实例下 对共享资源的保护

/**
分布式分成两类： 主动轮训分布式锁 redis mysql  类似于 CAS+自旋
               watch etcd                  类似于阻塞/唤醒
               回调机制 zookper
*/
