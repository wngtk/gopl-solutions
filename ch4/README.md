# 终于来到了第四章

我还是想吐槽一下 Go 语言的设计。for range 的用法太多了。循环 n 次, 用 `for i := range n` 这样的方式来写特别像 Python 写成 `for i in range(n)`。

Python 的 range() 是创建一个序列。for in 循环可以用在所有的序列类型上，自定义的类型只要实现对应的迭代器方法就能使用 for in 循环。一切都是这么的优雅。

看看 Go 语言，要想实现 `for i in range(1, 10)` 就不得不回到 C 语言的三段式 for 循环。倒不是传统的 C 语言语法有什么不好，而是 Go 语言一个语法的变体也太多了。`for range` 和普通 `for init; cond; post` 太不一致了(要是一致了好像也没有必要有两种)，非常不一致的地方还有 range 的左边可以有一个接收变量，也可以是两个。

不过我相信 Go 语言是解决工程问题的语言，Go 语言最大的优势在于编译速度快，容易部署。我对 Go 语言的吐槽或许是我当下还存在着一些疑惑吧。

> Go's purpose is therefore not to do research into programming language design; it is to improve the working environment for its designers and their coworkers. Go is more about software engineering than programming language research. Or to rephrase, it is about language design in the service of software engineering. -- https://go.dev/talks/2012/splash.article

Go 的目的不是进行编程语言设计研究；它的目的是改善设计人员及其同事的工作环境。Go 更多关注的是软件工程而不是编程语言研究。或者换句话说，它关注的是语言设计为软件工程服务。

## https://go.dev/talks/2012/splash.article

### 痛点

> slow builds

C/C++，Rust 编译比较慢。

> each programmer using a different subset of the language

说的就是 C++，现代 C++ 趋向于让大家使用同一个 C++ 子集。

> poor program understanding (code hard to read, poorly documented, and so on)

有些老的 C/C++ 代码写得十分晦涩，也缺乏文档工具，编程语言更是没有自己的文档，写一个库也可能没有文档。Java 的文档稍微好一点，现代的语言 JavaScript，Python，Go 和 Rust 基本上在文档和代码可读性都有了长足的进步。

> As a simple, self-contained example, consider the representation of program structure. Some observers objected to Go's C-like block structure with braces, preferring the use of spaces for indentation, in the style of Python or Haskell. However, we have had extensive experience tracking down build and test failures caused by cross-language builds where a Python snippet embedded in another language, for instance through a SWIG invocation, is subtly and invisibly broken by a change in the indentation of the surrounding code. Our position is therefore that, although spaces for indentation is nice for small programs, it doesn't scale well, and the bigger and more heterogeneous the code base, the more trouble it can cause. It is better to forgo convenience for safety and dependability, so Go has brace-bounded blocks.

为了安全和可靠 Go 语言还是采用大括号作为块结构

### 依赖 in Go

> Through the design of the standard library, great effort was spent on controlling dependencies. **It can be better to copy a little code than to pull in a big library for one function.** (A test in the system build complains if new core dependencies arise.) Dependency hygiene trumps code reuse. One example of this in practice is that the (low-level) net package has its own integer-to-decimal conversion routine to avoid depending on the bigger and dependency-heavy formatted I/O package. Another is that the string conversion package strconv has a private implementation of the definition of 'printable' characters rather than pull in the large Unicode character class tables; that strconv honors the Unicode standard is verified by the package's tests.

如果只是很简单的一个函数，确实没必要拉取一个很大的库。在写练习的过程中，有时候直接把整个程序的代码（总共也没几行）拷贝过来，比我依赖 Go 语言书的代码仓库要好得多，一来是代码都在这里，一下子就能看到所有的代码；二来是 Dependency hygiene trumps code reuse，不必在乎作者是否撤回仓库，也不用担心网络问题，

### Syntax

Go 语言的类型放在变量名字的右边，没记错的话这样的语法在实现编译器的时候是更简单的。Go 语言有其理由 https://go.dev/s/decl-syntax，把类型放在右边确实在函数指针上要更简单一点。

> One feature missing from Go is that it does not support default function arguments. This was a deliberate simplification. Experience tells us that defaulted arguments make it too easy to patch over API design flaws by adding more arguments, resulting in too many arguments with interactions that are difficult to disentangle or even understand. The lack of default arguments requires more functions or methods to be defined, as one function cannot hold the entire interface, but that leads to a clearer API that is easier to understand. Those functions all need separate names, too, which makes it clear which combinations exist, as well as encouraging more thought about naming, a critical aspect of clarity and readability.

没有函数参数默认值，看来 Go 的意思是设计更多的 API。C 语言也没有默认函数参数值，那写 C 代码的时候也应该定义更多的函数。可能在某种程度上来讲会有更清晰更易于理解的 API，**不是根据传参的值和数量来决定函数的行为，而是有明确的函数名字，表明了函数的行为函数会做什么操作**。

这算是 Go 改变代码编写习惯的一点，使用 Go 要求定义更多的函数或方法。Python 却是一个函数就能涵盖多种行为，通过参数默认值来调整逻辑。有些 Python 项目大量使用函数参数默认值（rpkg），要理解此次函数调用的意图，必须去读函数的注释或者代码。如果在函数名字上就表明了函数做什么事情，代码更清晰。

**Go 更加强调显式设计**：deffer 显式 RAII，range 显式迭代，不支持 default function arguments 必须显式定义不同的函数来表达不同的函数行为。

## Naming

没有函数重载，凭函数名字就能确定所指的函数。

## Error

This is a clear and simple design, easily understood. Errors are just values and programs compute with them as they would compute with values of any other type.

错误只是值，有一个 Error 方法，描述发生了什么错误。
