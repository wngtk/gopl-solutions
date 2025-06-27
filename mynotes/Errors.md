# Errors

- https://go.dev/talks/2012/splash.article#TOC_16.

Go does not have an exception facility in the conventional sense, that is, there is no control structure associated with error handling. (Go does provide mechanisms for handling exceptional situations such as division by zero. A pair of built-in functions called panic and recover allow the programmer to protect against such things. However, these functions are intentionally clumsy, rarely used, and not integrated into the library the way, say, Java libraries use exceptions.)

The key language feature for error handling is a pre-defined interface type called error that represents a value that has an Error method returning a string:

```
type error interface {
    Error() string
}
```

Libraries use the error type to return a description of the error. Combined with the ability for functions to return multiple values, it's easy to return the computed result along with an error value, if any. For instance, the equivalent to C's getchar does not return an out-of-band value at EOF, nor does it throw an exception; it just returns an error value alongside the character, with a nil error value signifying success. Here is the signature of the ReadByte method of the buffered I/O package's bufio.Reader type:

```
func (b *Reader) ReadByte() (c byte, err error)
```
This is a clear and simple design, easily understood. Errors are just values and programs compute with them as they would compute with values of any other type.

It was a deliberate choice not to incorporate exceptions in Go. Although a number of critics disagree with this decision, there are several reasons we believe it makes for better software.

First, there is nothing truly exceptional about errors in computer programs. For instance, the inability to open a file is a common issue that does not deserve special linguistic constructs; if and return are fine.

```
f, err := os.Open(fileName)
if err != nil {
    return err
}
```
Also, if errors use special control structures, error handling distorts the control flow for a program that handles errors. The Java-like style of try-catch-finally blocks interlaces multiple overlapping flows of control that interact in complex ways. Although in contrast Go makes it more verbose to check errors, the explicit design keeps the flow of control straightforward—literally.

There is no question the resulting code can be longer, but the clarity and simplicity of such code offsets its verbosity. Explicit error checking forces the programmer to think about errors—and deal with them—when they arise. Exceptions make it too easy to ignore them rather than handle them, passing the buck up the call stack until it is too late to fix the problem or diagnose it well.

---

> no control structure associated with error handling

这句话说得很准确，try-catch 就和 if-else 语句似的是一种控制结构。判断，循环都是一种控制结构。

之前初学异常的时候不知道什么时候用，也不知道该怎么用。第一次学 Java 的异常一上来就给我甩出 Checked 和 Unchecked，我当时看得是一脸懵逼，后来很长一段时间都害怕看到异常。后来看到 Python 的 cs61a 的例子居然使用异常机制 `HogLoggingException` 来控制循环结束，当时也不理解为什么要用异常来控制循环结束。函数里面不可能调用 break 来让 caller 的循环结束，当时也就认为这是一种巧妙利用语言的机制的做法。从控制结构的角度来理解的话，就说得通了。错误处理的控制结构是跨函数的控制结构，是专门用来处理错误的控制结构，有这种控制结构就可以将判断错误的代码和正常的业务逻辑分开。

> The key language feature for error handling is a pre-defined interface type called error that represents a value that has an Error method returning a string:

```
type error interface {
    Error() string
}
```
> Errors are just values and programs compute with them as they would compute with values of any other type.

Go 语言的 error 就是一个值，任何值都行，只要有一个 Error 方法能获得一个字符串。
