# 学习笔记

## 字符串



## 字符串匹配

暴力解法： O(mn)

rabin-karp

1. 假设字串的长度为M，目标字符串的长度为N
2. 计算子串的hash
3. 计算目标字符串txt中每个长度为M的字串的hash
4. 比较hash，如果hash值不同，字符串不匹配，如果hash值相同，使用朴素算法再次判断

KMP

当子串与字符串不匹配时，已经知道了前面已经匹配成功的那一部分的字符。KMP算法的想法时，设法利用这个已知的信息，不要吧搜索位置移回已经比较过的位置，继续把他向后移