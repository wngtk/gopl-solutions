# The Go Programming Language Book's Solutions

This repository provides the downloadable exercise solution programs
for the book, "The Go Programming Language"; see http://www.gopl.io.

These exercise solution programs are licensed under a <a rel="license"
href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative
Commons Attribution-NonCommercial-ShareAlike 4.0 International
License</a>.<br/> <a rel="license"
href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img
alt="Creative Commons License" style="border-width:0"
src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png"/></a>

**NOTE: there is no warranty for these code.**

## 安装项目中的可执行程序

这个项目模仿 gopl.io 组织代码的方式，有多个可执行程序。

通过 import path 安装项目中的可执行程序：

```sh
$ export GOPATH=$HOME/gobook # 选择工作目录
$ go install github.com/wngtk/gopl-solutions/ch1/helloworld # 获取/编译/安装
$ $GOPATH/bin/helloworld # 运行程序
Hello, world!
```

> 云风: 我发现我花了四年时间锤炼自己用 C 语言构建系统的能力，试图找到一个规范，可以更好的编写软件。结果发现只是对 Go 的模仿。缺乏语言层面的支持，只能是一个拙劣的模仿。https://blog.codingnow.com/2010/11/go_prime.html

