# redis

## 数据类型

### string
常用命令：`set、get、decr、incr、mget setnx setex`等<br>
string 类型是二进制安全的。意思是 redis 的 string 可以包含任何数据。比如jpg图片或者序列化的对象。


### hash
每个 hash 可以存储 2^32-1 键值对（40多亿）, 类似于Java 中 map
```shell script
hset key field1 value1 field2 value2....
hget key field1
> value1
```
常用命令：`hget、hset、hgetall`等


### list
Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。

list类型经常会被用于消息队列的服务，以完成多程序之间的消息交换。

常用命令：`lpush、rpush、lpop、rpop、lrange`等。

列表最多可存储 2^32-1 元素 (4294967295, 每个列表可存储40多亿)。


### set(无序)
Redis的Set是string类型的无序集合。和列表一样，在执行插入和删除和判断是否存在某元素时，效率是很高的。集合最大的优势在于可以进行`交集``并集``差集`操作。Set可包含的最大元素数量是4294967295。
集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。

### zset(有序)
Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。


## 集群

### 哨兵模式
>是基于主从模式做的一定变化，它能够为Redis提供了高可用性。在实际生产中，服务器难免不会遇到一些突发状况：服务器宕机，停电，硬件损坏等。这些情况一旦发生，其后果往往是不可估量的。而哨兵模式在一定程度上能够帮我们规避掉这些意外导致的灾难性后果。其实，哨兵模式的核心还是主从复制。只不过相对于主从模式在主节点宕机导致不可写的情况下，多了一个竞选机制——从所有的从节点竞选出新的主节点。竞选机制的实现，是依赖于在系统中启动一个sentinel进程。
#### 缺点:
1. 资源浪费, slave节点不提供服务只用于备份
2. 不能解决读写分离问题
#### 优点
部署简单

### cluster
>最小配置为6节点(3主三从) 其中主节点提供读写操作，从节点作为备用节点，不提供请求，只作为故障转移使用。Redis Cluster 采用虚拟槽分区，所有的键根据哈希函数映射到 0～16383 个整数槽内，每个节点负责维护一部分槽以及槽所印映射的键值数据。

#### 优点
无中心架构
高可用, 节点可扩展到1000, 可动态添加删除
高可用, 部分节点不可用时集群扔可用, 使用投票完成slave到master的提升
降低运维成本，提高系统的扩展性和可用性
#### 缺点
单机hot-key问题
只支持单机事物
只能使用数据库0
Redis Cluster 不建议使用 pipeline 和 multi-keys 操作，减少 max redirect 产生的场景