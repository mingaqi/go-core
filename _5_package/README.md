## 包的概念
>包用于组织 Go 源代码，提供了更好的可重用性与可读性
### src 
是源码目录，包含一个main函数，main所在的go文件开头必须是**package main**
此时，main函数才算是入口，一个程序只能包含一个main函数

### 新建包
新建 rectangle 包, 并且新建文件rectangle.go, 文件起始位置**package function**

#### 导出名字（Exported Names）
可以注意到 包中的两个func方法名都是大写字母开头。在 Go 中这具有特殊意义。在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。在这里，我们需要在 main 包中访问 Area 和 Diagonal 函数，因此会将它们的首字母大写。
则该程序运行时，编译器会抛出错误 **geometry.go:11: cannot refer to unexported name rectangle.area**

### init函数
> 所有的包都可以包含一个init函数，并且不应该有任何返回值和参数，在代码中不能显式的调用。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。
```go
func init(){
}
```
包的初始化顺序
1. 首先初始化包级别的变量
2. 调用init函数

如果一个包导入了另一个包，会先初始化被导入的包。
尽管一个包可能会被导入多次，但是它只会被初始化一次。
为了理解 init 函数，我们接下来对程序做了一些修改。

运行main

首先初始化被导入的包。因此，首先初始化了 rectangle 包。
接着初始化了包级别的变量 rectLen 和 rectWidth。
调用 init 函数。
最后调用 main 函数。

### 空白标识符

导入包却不使用，在go中是非法的，编译器是会报错，其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加。
```go
package main

import (  
    "rectangle" 
)

var _ = rectangle.Area // 错误屏蔽器

func main() {

}
```
var _ = rectangle.Area 这一行屏蔽了错误。我们应该了解这些错误屏蔽器（Error Silencer）的动态，在程序开发结束时就移除它们，包括那些还没有使用过的包。由此建议在 import 语句下面的包级别范围中写上错误屏蔽器。

```go
package main 

import (
    _ "geometry/rectangle" 
)
func main() {

}
```
有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。例如，我们或许需要确保调用了 rectangle 包的 init 函数，而不需要在代码中使用它。这种情况也可以使用空白标识符