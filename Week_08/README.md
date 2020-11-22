# 学习笔记

## 位运算

x ^ 0 = x
x ^ 1s = ~x
x ^ (~x) = 1s
x ^ x = 0
c = a ^ b => a ^ c = b; b ^ c = a
a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c

指定位置的位运算

1. x最右边的n位清零：x & (~0 << n)
2. x第n位的值：(x >> n) & 1
3. x第n位的幂值：x & (1 << n)
4. 第n位置为1：x | (1 << n)
5. 第n位置位0：x & (~(1 << n))
6. 将x最高位至第n位清零：x & ((1 << n) -1)

x = x & (x-1) 清楚最低位的1
x & -x 得到最低位的1

## 排序算法

```Go
// 冒泡
func bobbleSort(arr []int) {
    for i := 0; i < len (arr); i ++ {
        for j := i; i <= len(arr); j++ {
            if arr[i] > arr[j] {
                arr[i],arr[j] = arr[j],arr[i]
            }
        }
    }
}
```

```Go
// 选择排序
func selectionSort(arr []int) {
    pos := 0
    for i := 0; i < len(arr) - 1; i ++ {
        minIdx := i
        for j :=i + 1; j < len(arr); j ++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}
```

```Go
// 插入排序
func insertionSort(arr int[]) {
    if len(arr) < 2 {
        return
    }
    for i := 1; i < len(arr); i++ {
        for j := i; j > 0; j-- {
            if arr[j] < arr[j-1] {
                arr[j], arr[j-1] = arr[j-1], arr[j]
            }
        }
    }
}

```

## 布隆过滤器

快速判断一个元素是否存在（全部为1可能存在，否则一定不存在）

- 一个很长的二进制向量 （位数组）
- 一系列随机函数 (哈希)
- 空间效率和查询效率高
- 有一定的误判率（哈希表是精确匹配）