读完 Gopl 第五章，A tour of Go 的 Basic 部分的练习已经是小菜一碟了。

## Why Go?

学 Go 之前我更想用的是 Rust，因为 Go 语言在 DevOps 生态和软件部署上的优势，我不得不考虑 Go。有些原本用 Python 构建的命令行工具/基础设施，使用 Go 构建后更容易部署、性能更好、开发和维护更容易。Linux 发行版的大量基础设施，是使用 Python 构建的，尤其是 Fedora Infrastructure 大量使用了 Python，几乎全都是使用 Python。Rocky Linux 是继 CentOS 之后的 RHEL 开源构建，他们的基础设施使用 Go 语言开发。HashiCorp 开发的 Terraform，下一代 Vagrant，Ubuntu 的 snap，还有 Docker，Podman... 都是使用 Go 语言开发。DevOps 生态主要就是 Python 和 Go。

一个让我更进一步下决心决定学习 Go 的原因是 TypeScript 要用 Go 语言重新编写提高性能。可以预见，前端的 esbuild 和未来的 TypeScript 将会使 Go 在前端基础设施生态中不可缺少的语言。

## Go vs Rust

起初我还有点好奇，为什么有人花了精力学 Rust 最后却用 Go 用的多一些[^1]。读到第七章，Go 语言的语法，错误处理，方法，包都已经略有了解。我发现在写代码（或者某种操作的表达）上 Go 相比 Rust 并没有什么缺点。都尽可能让代码简单，编译器会尽可能让你处理错误。甚至 Go 语言的程序会比 Rust 的程序简单，代码更简短。

Go 语言是有 GC 的，除非是不能接受 GC，否则我更推荐使用 Go。因为 Go 简单，Go 的语法简单，有 GC，编译快，容易部署。正所谓人生苦短，我用 Python/Go。Python 和 Go 都是对人友好的编程语言。

[^1]: https://eli.thegreenplace.net/ Go 相关博客的数量比 Rust 多得多
