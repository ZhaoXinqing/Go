算法思想：
    贪心算法，局部最优解，不从总体最优上加以考虑；（活动安排）
    动态规划：弥补了贪心算法的不足,最优子结构（凑单问题）
    分治算法：MapReduce
    回溯：递归与非递归：

排序算法：
    冒泡（两两对比），
    选择（最大值放队首），
    插入（未排序插入已排序队列），希尔排序（插入排序的改进版本，递减增量），
    快速（分治思想，基准值），
    堆排序（建堆，堆首堆尾呼唤，缩小堆尺寸-1）
    - 这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：
            - 基数排序：根据键值的每位数字来分配桶；
            - 计数排序：每个桶只存储单一键值；
            - 桶排序：每个桶存储一定范围的数值；
        稳定和非稳定：例如：原始数据，a2和a4的位置都是3。对于稳定排序来说，排序后的序列，a2一定还是在a4前面。但是对于非稳定排序来说，就不一定了，可能排完序之后，a4反而在a2的前面了。

查询算法：
    1、深度优先(BFS)：从初始状态开始，总是先搜索至距离初始状态近的状态。每个状态都只经过一次，因此复杂度为O(状态数*转移方式数)。通常采用循环或队列（Queue）实现。
    2、广度优先(DFS)：从某个状态开始，不断转移，直至无法转移，回退到前一步，再继续转移到其他状态，直到找到最终解。通常采用递归函数或者栈（Stack）来实现。
    3、二分查找：对已经的排序好的队列进行中间值对比查找；

数据结构：
（1）线性数据结构：
    1、数组：在内存中连续存储，预先指定长度，大小固定，不能动态扩展，存放一组相同数据类型的的数据结构；（可利用下标进行随机访问，查找很快，但插入和删除效率低，可能浪费空间；从栈中分配空间（用NEW创建的在堆中））
    2、链表：采用动态分配内存的形式实现，每个元素增设指针域，用来指向后继元素，可用一组任意的存储单元存放数据元素，不会浪费内存，拓展灵活（访问元素的时候只能通过线性的方式由前到后顺序访问，效率较低，但添加删除元素效率高，通常从堆中分配空间）
    3、栈、队列和线性表：可采用顺序存储和链式存储的方法进行存储
    （线性表是逻辑结构，顺序表和链表是存储结构）
        - 顺序存储：借助数据元素在存储空间中的相对位置来表示元素之间的逻辑关系
        - 链式存储：借助表示数据元素存储地址的指针表示元素之间的逻辑关系
    - 队列：先进先出的线性结构，只允许在序列两端进行操作 
    - 栈：先进后出的线性结构，只允许在序列末端栈顶进行操作；
    - 线性表：包括顺序表（元素的地址是连续的）和链表（链表里面节点是通过指针连接的）（单、双，循环链表）
（2）树形结构：
2、树形结构：
结点间具有层次关系，每一层的一个结点能且只能和上一层的一个结点相关，但同时可以和下一层的多个结点相关，称为“一对多”关系。常见类型有：树、堆
二叉树：
是一种递归数据结构，是含有n(n>=0)个结点的有限集合，可以是空树，每个结点都恰好有两棵子树，其中一个或两个可能为空；二叉树中每个结点的左、右子树的位置不能颠倒，若改变两者的位置，就成为另一棵二叉树
完全二叉树：
从根起，自上而下，自左而右，给满二叉树的每个结点从1到n连续编号，如果每个结点都与深度为k的满二叉树中编号从1至n的结点一一对应，则称为完全二叉树
采用顺序存储结构既可利用数组元素的下标值确定结点在二叉树中的位置及结点之间的关系，但采用顺序存储结构存储一般二叉树容易造成空间浪费，链式结构可以克服这个缺点
二叉查找树：
a、左子树不空，则左子树上所有结点的值均小于根结点的值
b、右子树不空，则右子树上所有结点的值均大于根结点的值
c、它的左、右子树也分别是二叉查找树
平衡二叉树：
是左子树和右子树的高度之差的绝对值不超过1的二叉查找树（解决二叉查找树退化为链表的情况。），平衡二叉树的左子树和右子树都是平衡二叉树。
平衡二叉树的失衡及调整主要有四种情况：LL型、RR型、LR型、RL型
红黑树：
红黑树也是均衡二叉树（基础是二叉查找树），但具备自动维持平衡的性质，解决了平衡树在插入、删除等操作需要频繁调整的情况，
（1）从根节点到叶子节点的最长路径不大于最短路径的2倍；
①最长路径：若有红色节点，则必然有一个连接的黑色节点，当红色节点和黑色节点数量相同时，就是最长路径，也就是黑色节点（或红色节点）* 2；
②最短路径：从根节点到每个叶子节点的黑色节点数量是一样的，那么纯由黑色节点组成的路径就是最短路径
（2）新加入到红黑树中的节点为红色节点；
①红黑树中从根节点到每个叶子节点的黑色节点数量是一样的，此时假如新的黑色节点的话，必然破坏规则
②维持平衡主要通过两种方式【变色】和【旋转】，【旋转】又分【左旋】和【右旋】，两种方式可相互结合。       
堆：
是一种特殊的完全二叉树，其所有非叶子结点均大于（或小于）其左右孩子结点。均大于，则为大顶堆，均小于，则为小根堆；
并查集：
并查集是一种树型的数据结构，用于处理一些不相交集合（Disjoint Sets）的合并及查询问题。常常在使用中以森林来表示。
B树，B+树：
通常应用在数据库索引，可以认为是m（>2）叉的多路平衡查找树，为了追求性能，要减少IO次数，对于树来说，IO次数就是树的高度，而“矮胖”就是b树的特征之一;
其中，b+树是b树的一种变体，查询性能更好，其主要特点是：
（1）非叶子节点仅具有索引作用，所有数据都保存在叶子节点；
（2）对于范围查找来说，b+树只需遍历叶子节点链表即可，b树却需要重复地中序遍历
字典（trie）树：
就是利用字符串之间的公共前缀，将重复的前缀合并在一起，特别适用自动补全类应用，但比较消耗内存：因为它每一层都保存一个字典表，就算这层的节点只有一个也要有一组表。（使用的是指针类型，内存不连续对存储不友好，性能打折扣） 优点：查询效率比较高，对于一些范围较小的或者内存不敏感的应用可以使用

3、图形结构：
        - 在图形结构中，允许多个结点之间相关，称为“多对多”关系，可分为有向图和无向图
4、散列表
散列表（Hash table，也叫哈希表），是根据关键码值(Key value)而直接进行访问的数据结构。也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表。
理想散列表（哈希表）是一个包含关键字的具有固定大小的数组，它能够以常数时间执行插入，删除和查找操作。
每个关键字被映射到0到数组大小N-1范围，并且放到合适的位置，这个映射规则就叫散列函数
理想情况下，两个不同的关键字映射到不同的单元，然而由于数组单元有限，关键字范围可能远超数组单元，因此就会出现两个关键字散列到同一个值得时候，这就是散列冲突


二叉搜索树中，左子节点必须比父节点小，右子节点必须必比父节点大；
堆中，在最大堆中两个子节点都必须比父节点小，而在最小堆中都比父节点大。

一致性hash，hash槽：

复杂度分析：

散列表、哈希算法：

字符串处理：

字典树trie树；


相似度分析算法：
    - 欧几里得距离（Euclidean distance）（音乐推荐系统）
    - 朴素贝叶斯算法（过滤垃圾短信）

布隆过滤器：