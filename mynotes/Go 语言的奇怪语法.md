Go 是充满了 Unix 哲学和 KISS 原则的语言。很多的设计确实是“简陋”，在 KISS 原则里简单大于一切，所以看起来是简陋了点。

- https://github.com/wngtk/LearningProgress/blob/main/%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80%E7%9A%84%E5%A5%87%E6%80%AA%E8%AF%AD%E6%B3%95.md

## range

- https://go.dev/wiki/Range 

Range Clause 用来迭代数组、切片、字符串、map、通道。

正如 defer 是显式的 RAII，range 就是显式的 for each 循环迭代，相当于显示调用迭代器。

怎么读？

for all values in channel ch

for x := range ch {
    fmt.Println(x)
}

for all index and value in array arr

for idx, val := range {
    println(idx, val)
}


Go 语言比较简单

Go 语言的 for range 循环依赖编译器的实现，自定义的类型没有办法使用 range。Go 语言的语法支持竟然就是编译器的一个特殊支持。range 左边可以接收一个值，也可以接收两个值。这看起来就很不一致，运算符重载是不能根据返回值来重载的，也就没法根据赋值左边的数量来判断调用哪个函数来获取值。

Go 语言直接给 range “开后门”。针对 range 根据赋值的左边的变量数量来选择编译出什么样的代码。同样的情况还有 map 的取值，可以有一个变量接收 ok，也可以没有。

我认为要批评的是，Go 语言的 range 迭代不是在语言机制之上使用编程的方式（统一的方法 or 库）来实现，而是直接在编译器层面写死。

C++ 的 for range 循环是迭代器的语法糖。其他编程语言的 for 循环迭代一个数据结构都是使用迭代器的编程方式在标准库层面实现，编译器解析语法后还是调用的迭代器对应的方法/函数。而 Go 语言不是，所以我们自定义的数据类型无法使用 for range。

我不知道为什么 Go 语言要这样设计，但是 Go 语言这样设计的地方可能还有很多。range 我看就是想着有一个语句或者关键字能够方便迭代一个数组或者 map，然后就简单实现了一下。没有太多的深思熟虑，也不考虑用户自定义类型能不能用 range。

[ch4 README.md](../ch4/README.md) 读了一下

---

for range 的语法和 for 本身的语法很不一致。for 本身承担了 while。

Go 语言的一些机制的实现是通过 Go 编译器开后门实现的，没有一个统一的规定。Go 语言在工程上能够成功简直是奇迹。Go 的优点是编译速度快，容易部署。或许是因为这两点还没有其他的语言更快，更容易部署，所以 Go 语言获得了成功吧。

同一种语法的写法太多了，

- Go 自己会决定变量放在堆还是栈
- Go 自己会决定 range 的后面是一个变量还是两个变量的情况如何处理。
- Go 语言的 for 循环是传统的 for 循环，也是 while 循环也是 for each 循环。
- Go 语言的 range 和 map 的左边可以是一个 receiver 也可以是两个。

这些 Go 的特性不是通过语言机制来实现，而是 Go 编译器的特殊支持。

Rust 的错误检查是使用枚举 Option，Rust 使用 Option 就永远要检查 None。

本质上都是通过返回值来检查是否有错，Rust 不想 match/if let 就要指定使用某一个方法，Go 是直接根据语法在编译的时候就知道你应该是调用哪个方法。Go 不像 C++ 内部 STL 的方法都是暴露出来的，Go 语言内部的方法是

## 分号和代码格式

Go 语言的代码格式确实是一般。我不是很喜欢。有些我认为要加空格的地方自动格式化工具却不加。例如结构体初始化。

分号确实不用写，但是分号是自动加在行尾，所以你还是不能把 `.` 放到新起一行。所以代码看起来就有一点奇怪。