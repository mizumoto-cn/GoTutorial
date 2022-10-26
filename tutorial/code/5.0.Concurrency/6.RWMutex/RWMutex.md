# RWMutex

RWMutex 是一个读/写互斥锁，在某一时刻只能由任意数量的 reader 持有 或者 一个 writer 持有。也就是说，要么放行任意数量的 reader，多个 reader 可以并行读；要么放行一个 writer，多个 writer 需要串行写。

RWMutex 对外暴露的方法有五个：

RLock()：读操作获取锁，如果锁已经被 writer 占用，会一直阻塞直到 writer 释放锁；否则直接获得锁；
RUnlock()：读操作完毕之后释放锁；
Lock()：写操作获取锁，如果锁已经被 reader 或者 writer 占用，会一直阻塞直到获取到锁；否则直接获得锁；
Unlock()：写操作完毕之后释放锁；
RLocker()：返回读操作的 Locker 对象，该对象的 Lock() 方法对应 RWMutex 的 RLock()，Unlock() 方法对应 RWMutex 的 RUnlock() 方法。

一旦涉及到多个 reader 和 writer ，就需要考虑优先级问题，是 reader 优先还是 writer 优先：

reader优先：只要有 reader 要进行读操作，writer 就一直等待，直到没有 reader 到来。这种方式做到了读操作的并发，但是如果 reader 持续到来，会导致 writer 饥饿，一直不能进行写操作；
writer优先：只要有 writer 要进行写操作，reader 就一直等待，直到没有 writer 到来。这种方式提高了写操作的优先级，但是如果 writer 持续到来，会导致 reader 饥饿，一直不能进行读操作；
没有优先级：按照先来先到的顺序，没有谁比谁更优先，这种相对来说会更公平。

RWMutex 可以看作是没有优先级，按照先来先到的顺序去执行，只不过是 多个reader 可以 并行去执行罢了。