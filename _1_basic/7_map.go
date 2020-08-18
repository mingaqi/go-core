package _1basic

// map实现原理:https://www.jianshu.com/p/aa0d4808cbb8
// 线程安全map:https://www.jianshu.com/p/fc25c44f57f8

//map底层是hash表, 然后扩容缩容会发生rehash, 顺序会发生变化, 但是没有扩缩容的时候,顺序是有保证的, 但是golang为了让程序员不依赖这种不可靠的保证,就干脆遍历的时候加入随机数,然后不管什么时候遍历,顺序都是不保证的
