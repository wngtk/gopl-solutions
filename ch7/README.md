读完 Gopl 第五章，A tour of Go 的 Basic 部分的练习已经是小菜一碟了。

## Why Go?

学 Go 之前我更想用的是 Rust，因为 Go 语言在 DevOps 生态和软件部署上的优势，我不得不考虑 Go。有些原本用 Python 构建的命令行工具/基础设施，使用 Go 构建后更容易部署、性能更好、开发和维护更容易。Linux 发行版的大量基础设施，是使用 Python 构建的，尤其是 Fedora Infrastructure 大量使用了 Python，几乎全都是使用 Python。Rocky Linux 是继 CentOS 之后的 RHEL 开源构建，他们的基础设施使用 Go 语言开发。HashiCorp 开发的 Terraform，下一代 Vagrant，Ubuntu 的 snap，还有 Docker，Podman... 都是使用 Go 语言开发。DevOps 生态主要就是 Python 和 Go。

一个让我更进一步下决心决定学习 Go 的原因是 TypeScript 要用 Go 语言重新编写提高性能。可以预见，前端的 esbuild 和未来的 TypeScript 将会使 Go 在前端基础设施生态中不可缺少的语言。

## `sort.Interface`

`sort.Interface` 的例子非常好，`byArtist` 底层还是原本的类型，通过不同的类型调用不同的方法。类似的例子还有 LimitReader。**包裹原来的类型，提供一样的方法。包裹类型，可以直接是另外类型别名**。接口不管你是什么类型，只要你有接口的方法，就可以是当作这个接口类型。

## 7.3 实现接口 ⭐(没有继承真是太好了)

> 在 Go 语言里我们可以在需要时才定义新的抽象和分组，并且不用修改原有类型的定义。当需要使用另一个作者写的包里的具体类型时，这一点特别有用。。

Java 类实现接口需要修改 class 的代码, 而在 Go 语言里面不需要. 按照 OOP 的思维应该是 Audio 和 Video 都是 Streamer 的子类, 或者 Audio 和 Video 都实现 Streamer 接口. 所以在 Java 要加一个 Streamer 就要修改 Audio 和 Video. 但是在 Go 语言我需要将 Audio 和 Video 都当作同一种方式来处理的时候我再定义一个 Streamer, 完全不需要修改现有的 Audio 和 Video. 

## Caveats

7.5 节有个例子。buf 的类型在一开始就应该写成 `io.Writer`，而不是 `*bytes.Buffer`。**当使用接口的时候，变量的类型也应该声明为接口类型**。

```go
var buf io.Writer
if debug {
    buf = new(bytes.Buffer)
}
f(buf) // OK
```
