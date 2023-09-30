

# golang





# docker 

# docker启动MySQL
docker pull mysql/mysql-server  拉去最新mysql-server适配mac m1版本
1. 启动 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql/mysql-server
2. -v /Users/byte/Desktop/docker/mysql/data:/var/lib/mysql  -v /Users/byte/Desktop/docker/mysql/log:/logs   -v /Users/byte/Desktop/docker/mysql/conf:/etc/mysql/conf.d
3. 上述 -v 是Mysql数据映射到主机目录上 可以这么理解呗
2. 进入到mysql容器内部 docker exec -it mysql bash
   https://www.jianshu.com/p/eb3d9129d880  参考这个博客 基本上就是没有错误

# docker启动redis
1. 启动redis: docker run -p 6379:6379 --name redis -v /Users/bytedance/Desktop/whig/data/redis/data:/data -v /Users/bytedance/Desktop/whig/data/redis/redis.conf:/etc/redis/redis.conf -d redis redis-server /etc/redis/redis.conf
2. 参考博客 https://blog.csdn.net/u013868195/article/details/119778589



# -------------------==项目经历==------------------------



# 互娱研发-抖音 抖音全链路资产发现

# 项目概述

# 项目难点
1. 延迟问题：理论上十点采集的数据，11点应该识别出来，避免敏感字段透出太长时间。及时发现风险进行治理
   a. 前缀树优化 （children节点个数大于50的话，开启协程进行扫描 删除长时间未访问的数据、 前缀树进行压缩）
   b. 拉取overpass信息是通过中台服务，服务启动的时是缓存在本地。 --->缓存在中心化存储软件redis当中。
      同一批流量过来，如若版本号不一致 分布式锁更新缓存 其他流量进行等待  每个进程都需要进行拉取一遍 消耗太大
   c. bmq消费堆积问题。流量在百万QPS左右，这种mq消费堆积 一个group多个goroutine进行消费 采用ants协程池

# -------------------==面试==--------------------------



# bongcong 面试


华为：

1. golang数组作为参数，会影响外部吗
2. 进程、线程、协程之间通信方式
3. context 和 channel 关系和衍生
4. 子协程panic会影响主协程吗
5. http请求参数和响应参数、熟知状态码

算法：
1. 最大回文字符串

   知乎：
1. 令牌桶实现方式
2. redis_scene aof rdb实现方式
3. redis_scene zset底层数据结构 跳表时间复杂度

场景提：设计一个高并发的积分排行榜

算法： 回文链表

知乎二面：

1. 项目背景：链路层获取或者存储层获取 收集起来
   技术选型：bmq mysql redis abase 协程池、前缀树、令牌桶、通信协议、缓存结构（拉取overpass）
   redis、mysql 双写一致性  redis底层数据结构  AOF追加写机制(每写一条命令进行追加写到AOF文件当中) RDB快照机制(看频率 容易丢数据)
   redis分布式锁怎么实现的
   前缀树
   协程池
   分布式令牌桶  单机令牌桶



角色：维护注册任务、采集服务优化、要求实时性、10点采集数据11点透传到下游
不足：项目比较耦合、
1. CPU有尖刺状态 什么情况导致的 如何优化 如何进行解决 优化数据结构  LRU长度缩短点 利用多级缓存

场景提：
1. 并发高系统 go语言实现  缓存lru实现（实现上正常）  cpu每隔几分钟会有一个尖刺 分析原因  GC原因 分析原因
2. 设计签到系统 千万级QPS 同时运营需求 统计连续多少天签到 连续1、3、5签到 连续一个月签到 还需要历史签到能力做统计分析 半年内每个周五签到情况 redis只存当天的数据 -> mysql -> hive
   设计接口 设计存储 整条链路是怎么样
   千万级QPS数据库压力大了，怎么优化 setbit sing_status zhangsan+23829 1
   暴露签到接口（用户只关心今天已经签到） redis分布式存储 怎么实现的分布式存储（一致性hash）打到异步mq进行mysql存储
3. 高并发排行榜  zset做排序 redis做分布式存储 在最后做归并排序 每日排行、周排行、月排行、历史排行  用户王者荣耀、好友排行、段位排行
4.
-----------------------------------------------------------------------------------------------------------------------
八股：
最多redis zset底层 redis持久化存储（redis怎么保证数据不丢失、写数据处理命令）
网络还行
操作系统也还行  进程、协程、线程
Mysql基本也不问八股

讲一下MVVC 怎么用
每条记录都有一个trx_id
事物开启生成快照 min_txr_id  活跃事物ID  max_trx_id+1（当前活跃事物找到最大值+1）
当前事物要查询记录  在这个事物里查询这条记录的事物ID < min_txr_id 这条记录是可见的
ID>max_trx_id 往前追溯 找到符合规则的记录 事物可见的记录(1.小于trx_id  2.在min和Max区间范围之内，但是不在活跃事物ID当中 存在有的事物执行完了)
更新   当前读
查询   select for update

操作系统：分配内存 golang、c++、c  malloc分配内存  熟悉内存组成
使用两种方式：阈值
大于128K 在mmap区进行分配内存 mmap文件映射区 (free函数之后 mmap可以主动释放、 释放之后交给内核)
小于128K brk函数在堆上进行分配  堆上不能主动释放（回收机制：堆上高地址内存没被释放、地址值内存地址也没被释放（不能给操作系统回收掉）  导致堆上有内存碎片 在申请可以直接使用 可以重复利用 ）

网络:
https:加密过程
状态码：
websocket tcp



# 拼多多 一面 23年9月3号

1. 项目介绍 团队职责 
2. 项目中认为什么地方困难点

MySQL：
1. 聚族索引和非聚族索引区别
2. MySQL select update两个操作实现要求操作同一条记录   用乐观锁和悲观锁两种方式进行实现
3. select ... for update 如果记录不存在是否进行上锁  
   要了解上的什么锁 要不还是不能畅谈

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
1. 两个链表，判断相同元素 （这里边应该是lc相交链表 但是好像还是不一样 题目没有问清楚 ） 解题思路：遍历+hash 询问是否压缩空间
2. 二叉树，根节点上的路径和  lc-路径总和II的一个变种题


# 拼多多 二面 23年9月10号
1. 项目介绍 将近半个小时 面试官不知道难点在哪 感觉没什么难点

算法
1. 最长递增子序列个数  【1，2，3，5，4，9】 【1，1，1，1，1】  dp动态数组没A出来 将近浪费半个点
2. 个人情况的先关介绍

职业发展
1. 为什么考虑换工作  目前工作没有挑战性 被问到什么工作具有挑战性 
2. 希望C端还是平台方向发展  本人还是更倾向于业务  业务创造价值被  更希望希望被广泛使用 是自己价值感和认同感得以提升
3. 上一次绩效
4. 考研成绩
5. 本科成绩





# 快手一面 23年9月12日
1. 项目介绍

mysql
1. select 一条语句发生了什么 认证->连接->缓存->词法分析->语法分析->存储引擎层
2. binlog日志了解吗 逻辑备份 主要用于数据库主从同步与增量复制
3. binlog日志  update100条语句 binlog日志写一条 还是写100条 取决于数据库复制的模式
              【1.row_level 每行记录修改都会被记录下来，优点数据一致性 缺点：日志量比较大 主从同步 从消费binlog日志可能导致延迟 】
              【2.statement_level】记录更改记录的SQL语句，而不是每条记录。但是缺点：主从同步缺少上下文信息 使用特定功能函数例如last_insert_id 会导致歧义
              【3.mix_level】集合1 2 两种方式

4. undolog日志 作用 
5. 更新一条语句都有哪些日志会更新  undolog binlog redolog 

redis
1. 数据结构 
2. 字符串  sds 底层动态数组 TLV组层 不同长度的字符串使用不同的数据结构
3. 字符串  bitmap底层
4. 字符串  set key value=数字的话 使用long类型来存储
5. list   底层ziplist 紧凑的数组 zlbytes  zllen  zltail 
6. set    底层hash
7. zset   跳表+hash
8. .....
9. redis 怎么定义大key  相对而已
10. redis 什么算作热点key 28原理 80%流量查询热点key 

kafka
1. kafka架构
2. 幂等性原理 
3. 如何保证不重复消费  假如正在消费，但是突然节点宕机，消费成功但是offset没有+1 在该场景下如何保证不重复消费（生产者不生产重复消息，消费者在业务逻辑上做消费处理，可以加一个唯一标识 例如产品ID 入库判断）
4. 如何保证消息投递的可靠性

算法：
1. 合并两个有序链表  循环写出来 面试官再让用递归来写



# 快手二面 23年9.19号

1. 项目背景  难点  收货
2. 长短链系统设计  高并发、稳定性、扩展性、易用性等几个方面来设计 存储、编码和解码
   1. 商品ID+短链设计  比较菜   没有打好
3. redis集群部署方式  clauster codis 中心化存储  比较菜没答上来

算法

1. 12345输出一万两千三百四十五  文本格式调通了



# 迅雷一面

1. 项目介绍 认为项目成长的地方 以及 项目中难点
2. 为什么跳糟

golang
1. 并发常见的数据结构
2. context中什么功能最重要
3. mutex和atomic原子包实现并发安全区别
4. golang gc触发时刻 gc调度时刻
5. map 什么类型可以当做key  map中struct结构体可以当做key吗（没有回答上来，主要是hash函数由谁来实现的）
6. map 为什么并发不安全  不安全在哪
7. map rehash扩容机制 （分读请求和写请求）写请求写新的 读请求如果命中顺带写 读请求两个都命中那么以新的为准（因为新的可能写）
8. mutex实现原理 （简单说了一下就是用数组来实现，通过不同的标志位来判断是否进行上锁，标志位的更改利用原子能力来实现，并且有进行过优化）


mysql
1. mysql主动使用过事物没
2. mysql索引使用Innodb 建立索引一些注意事项
3. mysql 线上发生过死锁没
4. mysql 为什么不建议varchar较长类型建立索引 （更新、删除、查找耗时除外） 建立二级索引B+树 每个节点占据内存空间较大，如果超出16KB会进行分页，IO查询比较高
5. 联合索引 index(A,B,C) 
6. mysql死锁场景发生过没

redis
1. 常见数据结构
2. 说了下底层实现
3. 还有什么数据结构

算法：
给定二叉树任意节点， 求二叉树中序遍历的下一个节点 （没有做上来 二叉树这里边还是比较懵啊）主要是算法没回答上来

反问
1. 技术leader当组长 而不是管理型leader当组长
2. 项目变换频繁吗 或者OKR制定中途是否更改呢 
    个人感觉 迅雷面试官给人感觉不错 整体上引导面试者来说

#  23年9月27号 百度地图 一面与二面

## 一面 

**八股+设计**

1. 64位操作系统字长为啥是8字节 32位操作系统是几字节

2. golang slice 讲一下 OMAKE->OSLICE->runtime.makeslice结构体形式->扩容规则

   1. 扩容规则
   2. 讲扩容之后申请内存需要进行内存对齐（被面试官问住 操作系统为啥内存对齐，64位为啥是8字节）

3. linux 进程内存大概长什么样子->栈、堆、数据段、代码段。我说golang中堆可以分配2^48范围的堆内存。面试官问我为什么是48  **还是不会** 我说好像跟地址总线相关 但是不知道具体怎么来的

4. 查看CPU以及内存负载使用什么命令 top

   1. 在问top有什么具体参数吗？不知道转移话题 mac电脑top一下都出来进程、内存、CPU负载等信息
   2. 面试官说Linux下做过吗，莫有

5. 服务中限流如何做的？

   1. 分为分布式限流和本地限流 限流是限流自己的服务 而不是接入的外部服务，面试官比较诡异

   2. 本地令牌桶算法有什么缺点吗？没答上来好尴尬

   3. 问知道其他限流器吗？滑动窗口、计数器等限流，讲述了一下上述限流器的缺点。

   4. 分布式限流项目中如何做的，redis incr计数器 滑动窗口形式来做

   5. 本地令牌桶如何做的呢？使用开源的google_limit

   6. 有自己实现过吗？手写一个本地令牌桶算法

      ```go
      
      // 百度一面面试官被闻到过单机限流令牌桶如何实现
      type Bucket struct {
      	maxLimit int32
      	token    int32
      	psm      string
      	lock     *sync.Mutex 
      }
      
      func (b *Bucket) allow(psm string) bool {
      	b.lock.Lock()
      	defer b.lock.Unlock() //这么做就是好像有点简单呗
      	if b.token > 0 {
      		b.token--
      		return true
      	}
      	return false
      }
      
      func (b *Bucket) timer() {
      	tick := time.Tick(1 * time.Second)
      
      	for {
      		select {
      		case <-tick:
      			//重新将令牌桶初始化为最大值
      			b.lock.Lock()
      			defer b.lock.Unlock()
      			b.token = b.maxLimit
      		default:
      		}
      	}
      }
      
      
      ```

   7. 面试官追加必须要进行上锁吗使用Mutex这种关键字 我说可以使用atomic.SwapAndCompareAddInt32这个方法 让token值进行+1

   8. 你上述这个场景应该是针对method方法的，假如要是不同用户有不同的限流，假如A用户限流20，但是B用户是付费用户，那么限流就是60。这种场景下的限流又是要该如何实现呢，根据面试官提供的场景改成`map[int][int` 结构体来进行不同用户限流的划分。

      ```go
      
      // 百度一面面试官被闻到过单机限流令牌桶如何实现
      type Bucket struct {
      	maxLimit int32
      	token    int32
      	maxLimit map[int][int]  //这里边根据上述场景改成map形式
      	lock     *sync.Mutex 
      }
      
      func (b *Bucket) allow(psm string) bool {
      	b.lock.Lock()
      	defer b.lock.Unlock() //这么做就是好像有点简单呗
      	if b.token > 0 {
      		b.token--
      		return true
      	}
      	return false
      }
      
      func (b *Bucket) timer() {
      	tick := time.Tick(1 * time.Second)
      
      	for {
      		select {
      		case <-tick:
      			//重新将令牌桶初始化为最大值
      			b.lock.Lock()
      			defer b.lock.Unlock()
      			b.token = b.maxLimit
      		default:
      		}
      	}
      }
      
      
      
      
      ```

   9. 上述场景被问到是否可以进行优化，我说可以根据user_id来进行划分，将map分成多个map，在timer中分成多个timer来进行令牌token的续放。

6. 分布式令牌桶限流设计一个，是利用redis key value=hash来应对上述场景，但是好像造成大Key问题。最后说不会
7. 我说是否可以使用zset跳表这种形式来做，面试官问我跳表是什么，画了一下示意图
8. B树是什么（这块紧张了，关键不知道根节点下 可以固定多少个孩子，怎么确定树的高度）导致总体B树没有讲好。

多次引导面试官来问Go语言，面试官就是不问，很难受。面试官主要是C++语言，面试官说我问你C++你也不会呀。

**延伸**

1. 做过相关性能测试吗 
   1. 我说现在项目就是在重构当中，讲了一下提升的性能点。面试官说不是，抛开项目来讲是否做过其他性能测试。例如在48G机器上，长连接websocket 能支持多大QPS
   2. 我说没有做过，但是我知道go语言上性能优化的提升，每个版本做了哪些性能优化。哎 面试官还是不问Go 崩溃。
2. 项目现在主要几个人在做 我说两个人。
3. 最后终于来了一个Go方面知识 让我讲一讲slice  一顿输出 到 CPU内存对齐 为啥字长是8字节



## 二面 

先自我介绍一下

项目：

1. 问我解析流量多大QPS 我说在80万左右
2. 服务怎么抗住这么大的QPS的呢？我说本地缓存+中心化缓存。本地缓存存储idl信息
3. 面试官问：本地缓存是每个机器上都存储一份吗，针对不同Psm 我说是的
4. 面试官说：这样子很浪费内存空间啊，我说是，目前在进行优化使用中心化redis缓存来解决。
5. 面试追问：假如就是用本地缓存来解决这个问题如何进行解决呢。给个场景就是不同用户代表psm，尽量避免地域请求
   1. 我说我们目前是1000台左右的实例，如何按照现在方式的话，就是每个实例上存储不同范围的psm缓存信息，可以理解为user_id来进行划分1000台实例上。
   2. 面试官追问：但是这样做跟你目前的系统有什么差异？我说应该加一个Proxy代理层来实现，代理层决定那个user_id经过hash去访问那个实例信息。
   3. 面试过：说可以的，但是你引入了一个代理层又会带来什么影响？我说网络延迟以及个实例内部可能访问不均衡。
   4. 面试官追问：那么如何解决每个实例访问不均衡呢？我说分布式一致性hash算法 crc32这种，可以避免访问布局
   5. 面试官又追问：即便分布式hash也会导致分配不均，又如何处理呢？我说那就重代理层来决定，假如多长时间访问一个实例进行计数，如果超过一定阈值，请求别的实例。
   6. 面试官追问：还有别的解决方案了嘛？我说暂时想不到 
   7. 面试官追问：了解mysql 分片原理吗，直到Msyql是如何进行分片的呢？我说不了解
      1. 依稀记得mysql分片 好像也是利用一致性hash来做的呢。

算法：

问我刷题了，我说刷了一点。面试官说那咱们不按照lc来刷题，按照一个场景吧，比如地图

```go
			|
.—————.—————. 
      |
```

点定义：node-->表存 node_id [link1,link2,link3...]  link代表边 

边定义   link---->表存 link_id start_node_id end_node_id。

上述信息都是存储在表里边，问我，将这些点、边都加载到内存中应该使用什么数据结构。我说查询场景是什么，面试官说可能存在点、边 需要展示当前节点的周围路段信息。

--------------------

思考了一分钟：我说使用二维数组来存储

```go
node1   node2   node3   node4   node5
link1   link4   ..      ..      ..
link2   link5   ..      ..
link3
```

但是这种方案就是被面试官否决了，因为什么呢？主要根据点来查到边，边里有点，关键点边里边的点如何在方面查询。

面试官想要O(n)或者O(logn)来快速查询。这种思路下就是没有想出来比较更好的。

面试官直接给出数据结构形式，一个node节点+边代表一个路况节点，其实以一个节点来延伸展示出来的就是一个数状结构。

```go
type Node struct{
  node_id int
  link_id int
  next []*Node
}
```

将表里边信息全部装在到内存中  面试官不要让我去向递归退出条件什么的 

```go

var nodeLinkInfo map[int32][]int32  //key代表点 value代表边数组
var linkNodeInfo map[int32][2]int32 //key代表边，value[0]是边的起点ID value[1]是边的终点ID

type Node struct {
	nodeId int32   //nodeID代表地图中的一个节点
	linkId int32   //linkID代表某个节点上的一条边
	next   []*Node //上述点和边组成的路段的临近路况信息
}

// 这个是代表司机就是正在这个  A--司机---B 司机在当前边上和起点A 想要展示地图1000范围内的全貌
// 假如传进来的root节点是已知rootID和linkID
func generate(root *Node) *Node {
	if root == nil || root.nodeId == 0 {
		return nil
	}
	root.next = make([]*Node, 0)
	linkList := nodeLinkInfo[root.nodeId]
	if len(linkList) == 0 {
		return nil
	}
	for _, linkID := range linkList {
		nodeArray := linkNodeInfo[linkID]
		if nodeArray[1] == 0 { //也即就是这条边就是没有终点的呗
			return nil
		}
		for _, nodeID := range nodeArray {
			node := &Node{
				nodeId: nodeID,
				linkId: linkID,
			}
			root.next = append(root.next, node)

			generate(node) //这里边面试官告诉我先
		}
	}
	return root
}

```



