## 函数
```go
func functionname(parametername type) returntype {  
    // 函数体（具体实现的功能）
}
```

### 多返回值
```go
// 多返回值函数
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
```

### 命名返回值
```go
// 命名返回值
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
```

### 空白符(匿名变量)
以多返回值函数为例
```go
a,_ = rectProps(1,2)
```

### init()
1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
2 每个包可以拥有多个init函数
3 包的每个源文件也可以拥有多个init函数
4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
5 init函数`不能被其他函数调用`，而是`在main函数执行之前`，自动被调用
6 init函数无参数无返回值
