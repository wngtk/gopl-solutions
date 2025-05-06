# Go Modules

- [Using Go Modules](https://go.dev/blog/using-go-modules)

> Go语言的代码通过包(package)组织,包类似于其它语言里的库(libraries)或者模块(modules)。一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用。每个源文件都以一条 package 声明语句开始,这个例子里就是 package main, 表示该文件属于哪个包,紧跟着一系列导入(import)的包,之后是存储在这个文件里的程序语句。

main 包比较特殊。它定义了一个独立可执行的程序,而不是一个库。在 main 里的 main 函数也很特殊,它是整个程序执行时的入口

必须告诉编译器源文件需要哪些包,这就是跟随在 package 声明后面的 import 声明扮演的角色。

必须恰当导入需要的包,缺少了必要的包或者导入了不需要的包,程序都无法编译通过。这项严格要求避免了程序开发过程中引入未使用的包。import 声明必须跟在文件的 package 声明之后。随后,则是组成程序的函数、变量、常量、类型的声明语句。

---

import 导入一个包。import 后面的是 import path，因此 import 后面是一个字符串（路径用字符串来表示很合理），而不是一个语言可以识别的包名。包名和路径的最后一个元素相同。

Go 语言的模块和 Python，JavaScript，Rust 等其他语言中的模块完全不是同一个东西。C 语言没有模块，C 语言只有 .c 文件。Go 语言之前也没有模块，Go 语言只有 package。Go 语言的 package 是一个文件夹下的所有的 .go 文件共同组成一个包。所有的 .go 文件必须以 package 开头，一个文件夹下的 .go 文件的 package 

Go 语言的编译单元是 package，一次考虑一个包。Go 语言通过 import path 指定一个包。import path 是 module path 加上子目录的路径。

Module 在 Go 语言里面代表的是一个项目。一个项目就是一个 Go Module。Module 里面可以有很多的文件夹，每个文件夹都可以是一个 Package。

文件夹下的 .go 文件的 package 必须是一样的。同一个文件夹下的 .go 文件组成了一个 Package。不同文件夹下的 package 可以是一样的名字, 但仅仅是名字一样, 路径不一样, import path 就不一样, 就是不同的包.

package main 代表可执行程序。一个项目可以创建多个可执行程序，在不同的文件夹下的 .go 文件声明 package main 就是声明这个包是一个可执行程序。

例如 hello 文件夹是一个 Go Module 的根， foo， bar 是两个文件夹，这个两个文件夹下的 .go 文件都声明 package main：

```
hello/
├── bar
│   └── main.go
├── foo
│   └── main.go
└── go.mod
```

## 初始化 Go Module

`go mod init example.com/hello`

`example.com/hello` 是 Module path，也是 import path

## 一个项目创建多个可执行文件/库

每个 Go 程序都由包组成。程序从包 `main` 的 `func main` 开始运行。

- https://go.dev/tour/basics/1

可执行程序就是包 main，其他的包称为 non-main package。你的项目里面有多个 main 包就是有多个可执行程序，有多个非 main 的包就是有多个库。

可执行程序可以使用 go install 来安装 main 包，`go install <import path>`。例如安装 gopl 书中的 helloworld 程序：

```sh
go install gopl.io/ch1/helloworld@latest
```

go 会安装 helloworld 程序到 `$GOPATH/bin/helloworld`。

如果你的项目使用其他项目的包，通过 `go get [packages]` 下载依赖，并修改 go.mod 和 go.sum。

## 讨厌模块系统里面有双引号

设计不良的模块系统总是有双引号：Go，JS

JS 的 import 后面是模块名，模块名是目标模块的 .js 文件的相对或绝对路径。

Python 模块的文件名是模块名加后缀 .py。Python 的包是带命名空间的模块。
