Gopl 建议 Go 语言的错误处理的信息写成如下形式

> genesis: crashed: no parachute: G-switch failed: bad relay orientation

在调用 `fmt.Errorf()` 的时候的信息描述状态或者做的事情。

如果让我来设计 Go 语言我绝对不会使用 `:=` 来省略 `var`，因为 `:=` 和 `range` 是强制绑定的。

写 Go 语言的人有必要学好数据结构。

toposort 这个例子可以算是使用 map 来表示一个图。图的节点都是用字符串来表示，演示了使用 DFS 来拓扑排序。

findlinks3 演示了 BFS。因为图的结点是一个字符串，这几个图的例子都是使用 map 来记录是否访问过某一个结点。

练习 5.11：找到图里面的环（Cycle）。类似回溯的方式记录访问过的结点，和常规的遍历图一样有一个 visited/seen 记录结点是否访问过。另外用一个数据结构记录这次遍历过程中访问的结点，如果结点在这次遍历中访问两遍，那就是有环。每次遍历都是以 v 为起点, 尝试所有的顶点 v 遍历图。

---

## EOF

EOF 是一个 error 类型的变量。和 C 语言相比现代语言不使用原始类型的某一个值来表达特殊值，而是使用一个变量来表达特殊值。类似的例子有 Python 实现一个链表的时候，你可以使用 `empty = ()` 来表达空。

下面这个例子来自 https://cs61a.org/lab/lab07/

```py
class Link:
    """A linked list.

    >>> s = Link(1)
    >>> s.first
    1
    >>> s.rest is Link.empty
    True
    >>> s = Link(2, Link(3, Link(4)))
    >>> s.first = 5
    >>> s.rest.first = 6
    >>> s.rest.rest = Link.empty
    >>> s                                    # Displays the contents of repr(s)
    Link(5, Link(6))
    >>> s.rest = Link(7, Link(Link(8, Link(9))))
    >>> s
    Link(5, Link(7, Link(Link(8, Link(9)))))
    >>> print(s)                             # Prints str(s)
    <5 7 <8 9>>
    """
    empty = ()

    def __init__(self, first, rest=empty):
        assert rest is Link.empty or isinstance(rest, Link)
        self.first = first
        self.rest = rest

    def __repr__(self):
        if self.rest is not Link.empty:
            rest_repr = ', ' + repr(self.rest)
        else:
            rest_repr = ''
        return 'Link(' + repr(self.first) + rest_repr + ')'

    def __str__(self):
        string = '<'
        while self.rest is not Link.empty:
            string += str(self.first) + ' '
            self = self.rest
        return string + str(self.first) + '>'
```

一个合法的变量的地址肯定是唯一的，变量的地址会变化，这也消除了 0 地址。消除了 0 地址就不会有空指针解引用。Python 的 None 和 Go 的 nil 相同的一点是 None 也是一个对象，None 也是一个变量。

```py
>>> type(None)
<class 'NoneType'>
>>> id(None)
10746016
>>> 
```

为什么他们叫 `nil`/`None` 而不是 `NULL`，`nil`/`None` 都是一个标识符，代表的是一个变量/某一个特定类型的值的常量，而不是简单的 `(void*)(0)`。

用一个特殊的变量表达特殊值，对于有 GC 的语言，或者现代的编程语言都这样干。Rust 有 `Option` 类型，C++ 有 `std::optional` 类型，都是为了防止对空指针解引用。Rust 的 `Option` 类型是必须要处理所有的情况的，Rust 规定了 `match` 必须列出枚举的所有的情况，要么你就使用 `unwrap` 或者 `expect` 来处理错误。C++ 的 `std::optional` 也是类似，你要想获得对应的值你就必须处理错误，或者你确定你不处理。

Go 语言只要你接收了 `err` 的值，你就必须处理错误。除非你在 `err` 的地方用的是 `_` 作为占位符，表示你并不关心这个值。

Rust 是依靠强大的枚举来表达特殊值，**如果没有这样的强大的枚举的机制就可以使用一个独特的变量来表示终止或者空**。使用​​显式的特殊值对象​​（而非原始类型的魔法值）来表示“空”或“无值”状态，这种做法提升了代码的可读性和安全性。
