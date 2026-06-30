# go 

## 5 常用关键字

### 5.2 select

- 功能

  Go 语言中的 `select` 能够让 Goroutine 同时等待多个 Channel 可读或者可写，在 Channel 状态改变之前，`select` 会一直阻塞当前线程或者 Goroutine。

- 语法

  ```go
  select{
      case <condition>:
      <excution>
      case <condition>:
      <excution>
      ...
  }
  ```

- 注意事项：
  1. case中的表达式必须是Channel的收发操作

#### 5.2.1 实现原理

1. select 中没有 case

   直接阻塞当前的 goroutine，导致当前的 goroutine 进入无法被唤醒的永久睡眠状态。

   编译器直接将 `select{}` 转换为 `runtime.block` 函数，实现对 `goroutine` 的阻塞

2. select 中只有一个 case

   编译器会将 select 改写成 if 条件语句

   ```go
   // 改写前
   select {
   case v, ok <-ch: // case ch <- v
       ...
   }
   
   // 改写后
   if ch == nil {
       block()
   }
   v, ok := <-ch // case ch <- v
   ...
   ```

3. select 中有两个 case，其中一个是default

4. select 中有多个 case

---

### 5.3 make & new

- 功能

  `make` : 初始化内置的数据结构，slice, map, channel。

  `new` : 根据传入的类型分配一片内存空间，并返回指向这片内存空间的指针。

---

## 6 并发编程

### 6.1 上下文 Context

















































































































































































































































