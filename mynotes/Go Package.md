# 简要对比一下各个语言的模块系统

## Go

> The design of Go's package system combines some of the properties of libraries, name spaces, and modules into a single construct.[^1]

Go 语言的包是一个库，是命名空间（import path），也是模块。Go 语言的包是一个库很好理解，包是 Go 语言的编译单元，包的导出的（公开的）代码可以被其他代码使用。每个包都有自己的文件夹，文件夹下的 .go 文件组成了这个包，文件夹的路径决定了 import path，就形成了命名空间。Go 语言的包也是模块，Go 程序由包组成，main 包使用其他不是 main 的包，那些不是 main 的包就相当于某一个模块。

Go Modules 是用来管理依赖的。在 Go 的术语里，模块是一个 Go 语言的项目。See also: [Go Modules Reference](https://go.dev/ref/mod), [Organizing a Go module](https://go.dev/doc/modules/layout#packages-and-commands-in-the-same-repository), [Simple Go project layout with modules](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/)

Go import path 中的 internal 路径是特殊的，internal 中的包不能被其他的模块导入，internal 路径下的包只能在同一个模块里面使用。

## Rust

> As programs get larger, it's necessary to spread them over more than one file and put functions and types in different namespaces. The Rust solution for both of these is modules. [^2]

Rust 的模块系统是组合了命名空间和多文件。Rust 的编译单元是 Crate，Crate 是一个 .rs 文件和 .rs 文件引用的模块，mod（模块）可以全部写在一个 .rs 文件里，mod 也可以放在单独的 .rs 文件，模块的 .rs 文件的路径必须和模块的名字一样，一样的目录层级，一样的名字。

Rust 语言的模块系统没有包的概念，Rust 的包是 Cargo 的特性。Cargo 创建了一个包，Cargo 规定，包最多只能有一个 Library crate，可以有多个 Binary crate。

Rust 规定 `src/main.rs` 是可执行 crate，`src/lib.rs` 是库 crate。`bin/` 文件夹下的 `.rs` 文件都是可执行 crate。Rust 编译器一次考虑一个 crate，crate 不是包，也不是文件夹。crate 是一个 .rs 文件所用到的当前包下的所有的 .rs 文件。Rust 项目的代码组织是如此的高度统一。

Rust 项目的可执行程序的入口一定是 Binary crate，main 函数不会在某个奇怪的文件下。如果项目很大有多个库，每个库都会是一个独立的包，其他的文件不过是从 `src/main.rs` 或者 `src/lib.rs` 中分出去的。*Rust 项目有一种工程上的简单之美*，驾驭一个 Rust 项目的难点绝不会是构建系统。

Rust 没有声明 pub 的默认就是 private，只要没有在 `lib.rs` 声明 pub 或者 re-export，那么符号就是私有的。

## JavaScript

JavaScript 只有模块，现在一般使用 ES Modules。一个 .js 文件就是一个模块。模块机制参见：

JavaScript 本身没有包，包是 npm 和 Node.js 的机制。npm 的包 Cargo 很类似。npm 的包可以提供多个可执行的脚本，一个包只能是一个库，库里面有很多的 JavaScript 模块。

npm 包应该使用 `"exports"` 字段声明导出的模块，使用 `"exports"` 字段导出模块后其他模块默认就是私有的。For more details see the [node.js documentation on package entry points](https://nodejs.org/api/packages.html#package-entry-points)

[JavaScript 没有包](./JavaScript%20没有包.md)

## Python

模块是包含 Python 定义和语句的文件。其文件名是模块名加后缀名 `.py`。模块有自己的命名空间。模块内的对象需要通过 `modname.itemname` 来访问。

`python -m module-name` 运行一个 Python 模块，Python 关心的不是你的文件在哪，叫什么名字，只关心模块。模块名字就决定了源文件的路径和名字。Python 的包是一种通过使用“点分模块名”来组织Python模块命名空间的方式。包就是带命名空间的模块，包（文件夹）下有一个 `__init__.py` 那么包名也是一个模块，所以包就是命名空间，包也可以说是模块。包除了用来组织代码，包就是命名空间。

理解 Python 如何找模块这一点很重要，关系到你能不能理解 Python 解释器如何找到你的代码，关系到当你想运行代码的时候应该怎么运行。关于解释器如何找模块见 6. Modules

See also: [6. Modules](https://docs.python.org/3/tutorial/modules.html), [`__main__.py`](https://docs.python.org/3/library/__main__.html)

Rust、JavaScript、Go 都有办法处理私有的代码文件和符号，唯独 Python 所有的东西都是公开的，那 Python 在语义化版本号上如何做到接口的兼容性保障呢？恐怕是做不到，用户总是能访问到库的内部。Python 在软件交付和工程上还是稍逊。

## 总结

包/模块/库的用途主要是形成命名空间，方便构建、测试、分享。

[^1]: https://go.dev/talks/2012/splash.article#TOC_8.
[^2]: https://github.com/stevedonovan/gentle-intro/blob/master/src/4-modules.md
