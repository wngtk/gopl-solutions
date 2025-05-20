编程珠玑的第一章就有这样的位图的（bitmap）一个例子。印象中我见过其他语言的实现。

要不是 Go 语言 nil 的 slice 和长度为 0 的 slice 的行为一样，IntSet 的初始化就会很麻烦，就不得不弄一个 `NewIntSet()` 函数。所以 nil 值的 slice 和长度为 0 的 slice 的行为一样，是一个明智之举。
