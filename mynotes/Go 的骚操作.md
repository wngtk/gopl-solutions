# Go 的骚操作

Go 语言有 nil，但是没有 NULL，也没有 null_ptr。**nil 是一个特殊的标识符，是预先已经声明好的一个变量**。nil 不代表地址为 0，和 Rust 的 None 一样，是一个标识符。Rust 对 None 的处理方式，None 是一个枚举，None 是一个值，None 是一个变量。因为 nil 并不代表地址为 0，所以我认为 Go 语言还是实现了消除空指针。

Go 没有 Rust 那么强大的枚举，要手动检查 `if err != nil {}`，但是 Rust 也是要 `if let Some(x) = y { }`。从两种语言在错误处理的方式来看，Go 和 Rust 代码的风格会有很大的差别。在Go中，错误处理有一套独特的编码风格。成功时的逻辑代码不应放在else语句块中，而应直接放在函数体中。Go 语言的 if 块内是错误处理，而 Rust 把成功的逻辑放在 `if let Some(x) = y { }` 块内。两种语言都是通过返回值的机制来处理错误，本质上都差不多。所以我认为 Go 语言的错误处理比之 Rust 没有什么缺点，也是不错的，只是表现形式不同。See also: [EOF](../ch5/README.md)

Go 语言只要你接收了 `err` 的值，你就必须处理错误。除非你在 `err` 的地方用的是 `_` 作为占位符，表示你并不关心这个值。

Rust 还有一个东西叫 unit 类型，Rust 的 unit 类型是 `()`，这个类型专门表达传统 C 语言里的 `void`，因为 Rust 是基于表达式的编程语言，所有的东西都是表达式，if 也是表达式，函数块代码块也是表达式，什么都没有返回的表达式的值就是 `()`，类型是 unit，unit 类型就一个值 `()`。`()` 是一个值，不是表示一对括号。

Lua 的 nil 也是一个值，但是 Lua 的 nil 是一个类型，nil 类型只有一个值 nil。

Python 的 None 是一个变量，NoneType 类型的变量。对于动态类型的语言，NoneType 类型的变量可以赋值给任意类型这很合理。静态类型的语言像 Go 就只能赋值给特定的类型 nil。

C++/Python 的运算符都会对应一个函数，C++ 可以重载运算符，Python 有运算符对应的魔术方法可以覆盖。Rust 的运算符都对应了一个 Trait，通过实现 Trait 可以做到运算符重载。而 Go 语言的 `comparable` 似乎只是一个特殊的标记，只是告诉 Go 语言的编译器编译时进行检查。

我的一点小心得：`[]Type`，`map[Type]Type` 可以看作是特殊标识符的结构体，给他们赋值 nil，实际上给他们的某一个成员赋予 NULL。`Type` 是 Go 语言里面表达任意类型用的。

Go 有一些特别骚的语法，可以让代码少写一点。最明显的例子就是大写字母开头的对象是导出的(export)。

- 一个nil值的slice的行为和其它任意0长度的slice一样。（slice 的 append 操作是需要接收返回值的，这是传统的 C 语言的一种编程方式，类似 K&R 的二叉树的例子。假如让我在 C 语言里实现 Go 语言这样的 slice，并且支持 append, len 等操作，我也会让 NULL 和长度为 0 的动态数组的行为一样。如果不让 nil 值的 slice 的行为和长度为 0 的 slice 一样，你就要在很多的地方写上 `if xxx == nil { xxx = make([]int) }`，或者使用前不得不先 make，显然这样没有任何的意义，只会是不得不多写一些代码）See also: [IntSet](../ch6/intset/README.md)
- map上的大部分操作,包括查找、删除、len和range循环都可以安全工作在nil值的map上,它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常。（可以把存入操作想象成调用put方法，在一个方法里面改变 this 改变不了调用这个方法的实例的地址的，所以很自然如果 this 是 NULL，我无法更新 `obj.method()` 的 obj 的地址。）
- 可以对结构体的字面量取地址（第一次看到这个操作我是觉得有点奇怪的，但是初始化一个结构体的时候，如果结构体的成员是一个结构体的指针，这个特性就变得十分有用）。
- 函数参数类型是数组的指针, 在函数内部还是一个数组, 只不过这个数组的值可以改了. 数组就好像是一个结构体, 传递数组和传递结构体一样, slice 才像是一个指针.
- 多返回值就像函数参数似的，也算是本地变量。
- 匿名结构体成员：匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针，这个成员的名字就是数据类型的名字，只是用起来的时候可以少写一次点，成员的属性和方法都可以直接访问。(好处是不用自己写成员的属性和方法的转发了，这是优点啊，虽然刚开始接触的时候我懵了一下，但是这实际上就是自动生成了成员属性和方法的转发)
- 接口是不需要显式实现的，接口是隐式实现的。只要你有一个方法名字参数返回值都符合接口定义的，那么这个类型就是实现了这个接口。（error 类型就是一个有 Error 方法的接口，什么方法都没有的接口就是 any）
