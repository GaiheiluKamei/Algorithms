## ch01-入门

- 根据元素的组织，数据结构可以分成两种类型：
  - 线性数据结构(*Linear data structures*)：元素可以按顺序访问，但不是必须按顺序存储所有元素。例如：Linked Lists, Stacks and Queues.
  - 非线性数据结构(*Non-linear data structures*)：数据结构的元素以非线性顺序存储/访问。例如：Trees and graphs.

- 为了简化解决问题的过程，我们将数据结构与其操作结合在一起，称为**抽象数据结构**(*Abstract Data Types -- ADTs*)。一个抽象数据结构包括两部分：
  - 数据声明
  - 操作声明

-  常用的抽象数据结构包括：Linked Lists, Stacks, Queues, Priority Queues, Binary Trees, Dictionaries, Disjoint Sets (Union and Find), Hash Tables, Graphs……例如，Stack使用LIFO（Last-In-First-Out)的机制将数据存储在数据结构中，它的常见操作有:创建Stack，入栈，出栈，查找栈的当前顶部，查找栈中元素的数量，等等。

- **While defining the ADTs do not worry about the implementation details. They come into the picture only when we want  to  use  them**.

## ch02-递归和回溯

- Example Algorithms ofRecursion
  - Fibonacci Series, Factorial Finding
  - Merge Sort, Quick Sort
  - Binary Search
  - Tree Traversals and many Tree Problems: InOrder, PreOrder PostOrder
  - Graph Traversals: DFS [Depth First Search] and BFS [Breadth First Search]
  - Dynamic Programming Examples
  - Divide and Conquer Algorithms
  - Towers of Hanoi
  - Backtracking Algorithms

- Time Complexity: O(n). Space Complexity: O(n) for recursive stack space.

- Example Algorithms ofBacktracking
  - Binary Strings: generating all binary strings
  - Generating k−ary Strings
  - N-Queens Problem
  - The Knapsack Problem
  - Generalized Strings
  - Hamiltonian Cycles
  - Graph Coloring Problem
  
## cho3-链表

数组是一种静态结构，不方便放缩，而且插入和删除也比较费力。链表是一种动态数据结构，解决了数组的某些限制。

- 链表的属性：
  - 节点数量不固定，可以做动态增减
  - 每个节点由两部分组成：数据和对下一节点的引用
  - 最后一个节点的引用是nil
  - 链表的入口是head，head是对第一个节点的引用；如果链表为空，head是空引用

- 以下操作使链表成为一个ADT：
  - 主要操作：
    - 插入元素（Insert）
    - 删除元素（Delete）
  - 次要操作
    - 删除链表（Delete List）
    - 计算链表元素个数（Count）
    - 查找第n个元素

- 数组
  - 随机访问-访问数组元素的时间是常量：基地址+(索引*元素大小)，一次乘法和加法的操作时间是固定的
  - 优点
    - 使用方便；访问元素速度快
  - 缺点
    - 预先分配所有内存，有时候内存也许不够
    - 基于位置的插入复杂：要移动很多元素

- 链表
  - 易于放缩（常量）
  - 查找开销大，最坏情况为O(n)
  - 操作复杂：如果要删除最后一个元素，必须要遍历链表找到倒数第二个元素，将其指针设为nil
  - 指针需要额外存储空间

- 单链表
  - Each node is allocated in the heap with a call to `new()`, so the node memory continues to exist until it is explicitly deallocated with a call to ݁݁`free()`. The node called a *head* is the first node in the list. The last node's next pointer points to *nil* value.

