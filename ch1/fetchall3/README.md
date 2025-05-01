对 goroutine 还不是很熟悉，fetchall 只会发一次消息然后 return，main 函数里面也是发起多少次 go 就读多少次消息。
