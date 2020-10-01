# List

## list 的特性
- 可以使用 + 将两个 list 相加
- in 判断元素是否在 list
- 成对出现的 list 可以用 for 循环提取
```python
>>> [41,42,43,44]
[41,42,43,44]
>>> [41,42,43,44] + [41,42,43,44]
[41,42,43,44,41,42,43,44]
>>> len([41,42,43,44])
4
>>> 41 in [41,42,43,44]
True
>>> 41 not in [41,42,43,44]
False
>>> [41,42] in [41,42,43,44]
False
>>> for i in [41,42,43,44]:
        print(i)
41
42
43
44
>>> for i,j in [[41,42],[43,44]]:  # 适用于成对出现的序列
        print(i,j)
41 42
43 44

```

## range
range 是连续数字的序列
- 用_表示不关心序列里的值
```python
>>> list(range(-2,2))
[-2, -1, 0, 1]
>>> list(range(2))
[0, 1]
>>> for _ in range(3):   # 用_表示不关心序列里的值
        print("Go bear!")
Go bear!
Go bear!
Go bear!
```

## 列表推导 (List Comprehensions)
- `[function(x) for x in list if boolean(x)]`
```python
>>> odds = [1,3,5]
>>> [x+1 for x in odds]
[2,4,6]
>>> [x for x in odds if 25 % x == 0]
[1,5]
>>> [x+1 for x in odds if 25 % x == 0]
[2,6]
```

## 字符串 (Strings)
字符串本身是字符的不可改变的 array
- 可以用 `len(string)` 求的字符串数量
- 可以用 `string[i]` 提取单个字符
- 可以用 in 判断子字符是否在字符中

```python
>>> city = 'Berkeley'
>>> len(city)
8
>>> city[3]
'k'
>>> 'here' in "Where's"
True
```

## 切片 (Slice)
slice 是创建出了一个新的 list

```python
>>> odds = [1,3,5,7,9]
>>> odds[1:3]
[3, 5]
>>> odds[:3]
[1, 3, 5]
>>> odds[1:]
[3, 5, 7, 9]
>>> odds[:]
[1, 3, 5, 7, 9]
```

## 序列函数
### sum
`sum(iterable, start=0)` -> value
- 返回的 value 是：`start` (默认值:0) 加上一个可迭代的数字，当可迭代为空时，返回开始值。

```python
>>> sum([[2,3],[4]],[])  # 相当于 [] + [2,3] + [4]
[2, 3, 4]
```

### max
- `max(iterable, key=func)` -> value
- `max(arg1, arg2, *args, *[, key=func])` -> value

```python
>>> max(range(10),key = lambda x: 7-(x-4)*(x-2))  # y = 7-(x-4)*(x-2) 当 y 取最大值时， x是？
3
```

### all & any
`all(iterable, /)`
- 如果 `bool(x)` 迭代中的**所有**x都为真，则返回真。
- 如果可迭代为空，则返回**True**。
`any(iterable, /)`
- 如果 `\bool(x)` 迭代中的**一个**x为真，则返回真。
- 如果可迭代为空，则返回**False**。

```python
>>> all([2,4])
True
>>> all([2,4,0])
False
>>> all([2,4,[]])
False
>>> all([2,4,[1]])
True
```