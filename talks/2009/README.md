```
The big picture
 Fundamentals:
 Clean, concise syntax.
 Lightweight type system.
 No implicit conversions: keep things explicit.
 Untyped unsized constants: no more 0x80ULL.
 Strict separation of interface and implementation.
Run-time:
 Garbage collection.
 Strings, maps, communication channels.
 Concurrency.
Package model:
 Explicit dependencies to enable faster builds
```

Go 语言确实做到了 Clean, concise syntax。轻量级的类型系统。没有隐式转换。Untyped 原来不是没有类型，而是类型暂不确定的，就不需要像 C 语言的常量字面量一样 `0x80ULL` 才表示一个 `unsinged long long` 类型。Untyped unsized constants 这确实能让代码更干净, C: `3.23f` 才是 float 类型，Rust: `3.23f32`，而 Go 只需要 `3.23` 就可以了, 编译器会在用这个常量的地方自己定类型和大小。

