### go的writer和reader

> go的writer和reader接口设计遵循了Unix的输入和输出, 一个程序的输出可以是另一个程序的输出

```go
var 
(
    Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

```
这三种标准输出都是*file, 而file恰恰实现了`IO.writer和IO.reader`接口, 所以它同时具备了输入和输出功能.


Go标准库的io包也是基于Unix这种输入和输出的理念，大部分的接口都是扩展了`io.Writer`和`io.Reader`，大部分的类型也都选择地实现了`io.Writer`和`io.Reader`这两个接口，然后把数据的输入和输出，抽象为流的读写。所以只要实现了这两个接口，都可以使用流的读写功能。


io.Writer和io.Reader两个接口的`高度抽象`，让我们不用再面向具体的业务，我们只关注，是读还是写。只要我们定义的方法函数可以接收这两个接口作为参数，那么我们就可以进行流的读写，而不用关心如何读、写到哪里去，这也是面向接口编程的好处。

```go
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
    Write(p []byte) (n int, err error)
}

```
这是Wirter接口的定义，它只有一个Write方法。它接受一个byte的切片，返回两个值，n表示写入的字节数、err表示写入时发生的错误。

```go

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
type Reader interface {
    Read(p []byte) (n int, err error)
}

```

这是io.Reader接口定义，也只有一个Read方法。这个方法接受一个byte的切片，并返回两个值，一个是读入的字节数，一个是err错误。

[摘自CSDN](https://blog.csdn.net/weixin_34205076/article/details/93042605)