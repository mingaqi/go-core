[goroutine和并发相关](https://www.liwenzhou.com/posts/Go/14_concurrence/) <br>

[[典藏版]GMP模型](https://www.jianshu.com/p/fa696563c38a) <br>
[B站视频讲解](https://www.bilibili.com/video/BV1jT4y1j7Cd)

**CSP通过通信来共享内存，而非通过共享内存来通信**
### 并行和并发
并发: 同一时间段内执行多个任务(你和两个女朋友用微信聊天)
并行: 同一时刻执行多个任务(你和朋友都在用微信和女朋友聊天)

go语言通过`goroutine`实现并发, `goroutine`类似于线程,属于`用户态`线程,可以根据需要创建成千上万的`goroutine`并发工作,`goroutine`是有go运行时(runtime)调度, 而线程由操作系统调度完成.

go语言提供`channel`在多个`goroutine`中进行通信, `goroutine`和`channel`是go语言继承CSP(communicating Sequential Process)并发模式的重要实现基础

### goroutine

#### 可增长栈
`OS线程`(操作系统线程)一般有固定的栈内存(通常为2M, Java可以通过JVM参数设置不同大小), `goroutine`一开始只有2KB大小的栈内存, `goroutine`不是固定的它可以按需增大或缩小, goroutine栈大小可以大到1GB

#### GMP

GMP是golang运行时层面的实现, 是golang自己实现的一套调度系统, 区别于系统调度OS线程


- G(Goroutine) 就是goroutine, 为用户级的轻量级线程, 里面存放本goroutine信息之外, 还有与所在P的绑定信息.
- M(Machine) 对内核级线程的封装，M与内核线程一般是一一映射的关系, 一个goroutine最终要放到M上执行（真正干活的对象, golang限制10000个)
- P(Processor) 是一个抽象的概念，并不是真正的物理CPU ,管理着一组goroutine队列, P里面存放当前goroutine的运行的上下文环境(函数指针, 堆栈地址, 地址边界), P对自己管理的goroutine队列做一些调度(比如把占用CPU时间较长的goroutine暂停, 运行后续的goroutine),当自己的队列消费完了就去全局队列里取, 如果全局队列也消费完了会去其他P的队列里抢任务.

P与M一般也是一一对应: P管理着一组G挂载在M上运行. 当一个G长久的阻塞在一个M上时, runtime会新建一个M, 阻塞G所在的P会把其他的G挂载在新建的M上. 当旧的G阻塞加成或者r任务其已死掉时回收旧的M

P的个数是通过`runtime.GOMAXPROCS()`设定, 最大265.

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。


[ants协程池](https://taohuawu.club/high-performance-implementation-of-goroutine-pool)