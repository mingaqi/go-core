trace and pprof

https://zhuanlan.zhihu.com/p/141640004
[Golang 大杀器之跟踪剖析 trace](https://www.jianshu.com/p/81b6c0df66d1)


[GC](https://zhuanlan.zhihu.com/p/74853110)


[escape-逃逸分析](https://driverzhang.github.io/post/golang内存分配逃逸分析/)
1. 栈中创建的对象作为返回值逃逸
2. 闭包所引用的对象太乙
3. 动态类型逃逸, fmt.Println(a …interface{}) 编译期间很难确定其参数的具体类型，也能产生逃逸。
4. 栈空间不足逃逸 例如创建长度过长的切片
5. 指针逃逸 , 函数内使用new()创建对象, 并返回指针变量 