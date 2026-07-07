# go 

## 3 数据结构

### 3.1 数组

- 类型：

  只有存储元素类型相同且大小相同的数组在go中才是同一个类型。

- 数组初始化位置：

  当前数组初始化位置，在堆上还是栈上是在编译期确定的。

- 初始化

  ```go
  arr1 := [3]int{1,2,3}
  arr2 := [...]int{1,2,3}
  ```

  这两种对数组进行初始化的方式在运行期间得到的结果是完全一样的，后一种方式在编译期间就会被转换为前一种，也就是编译器对数组大小的推导（上限推导）

- 语句转换

  初始化数组字面量的优化

  1. 元素数小于或等于 4 ，直接将数据元素放在栈上
  2. 数组元素数大于 4，会将数据中的元素放置到静态区并在运行时取出】

- 数组访问与赋值

  数组和字符串的简单越界错误都会在编译期间发现，如直接用整数或者常量访问数组；

  但是使用变量访问数组编译器就无法提前发现错误，需要运行时阻止不合法的访问。

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

















































































































































































































































