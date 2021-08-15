## mutex

- sync.Mutex

Lock()

Unlock()

利用defer

- sync.WaitGroup

WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。

Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。

https://studygolang.com/articles/12972

## reference 

https://zhuanlan.zhihu.com/p/295397902