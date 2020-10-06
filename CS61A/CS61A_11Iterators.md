# Iterators
容器可以提供迭代器(按顺序访问其元素)
- ```iter(iterable)```: 在可迭代值的元素上返回一个迭代器
- ```next(iterator)```: 返回迭代器中的下一个元素

这个功能可以由 带 `nonlocal index `的高阶函数实现。

## 字典的迭代器

特点：
- 字典的 `key`, `value`, `item`(输出的是 `tuple`) 都是可迭代的
- 字典的 `size` 在迭代器中不能改变

## `iter` vs `range`

特点：
- `iter` 中已经输出的值不会再出现
- `range` 每次都能输出想要的值
  
```python

# ------ range --------
>>> for i in range(3,6):
>>>     print(i)
3
4
5

>>> for i in range(3,6):
>>>     print(i)
3
4
5

# ------ iter --------
>>> ri = iter(range(3,6))
>>> for i in ri:
    print(i)
3
4
5

>>> for i in ri:
>>>     print(i)

>>> 
```

## `python` 内建的迭代器功能
许多内建的 Python 序列操作都会返回迭代器，这些迭代器会 lazily 地计算结果
  - `map(func, iterable): `: Iterate over func(x) for x in iterable 
  - `filter(func, iterable):`: Iterate over x in iterable if func(x)
  - `zip(first_iter, second_iter):` : Iterate over co-indexed (x, y) pairs
  - `reversed(sequence):` : Iterate over x in a sequence in reverse order

要查看迭代器的内容，将生成的元素放入一个容器中:
  - `list(iterable)`: Create a list containing all x in iterable
  - `tuple(iterable)`: Create a tuple containing all x in iterable
  - `sorted(iterable)`: Create a sorted list containing x in iterable

```python
>>> bcd = ['b', 'c', 'd']
>>> [x.upper() for x in bcd]  # 这是直接生成的情况
['B', 'C', 'D']

>>> caps = map(lambda x: x.upper(), bcd)  # 这是返回的迭代器 
>>> next(caps)
'B'
>>> next(caps)
'C'
```

## 生成器(`generator`)
生成器是特殊的迭代器，使用生成器函数来获得。

```python
def plus_minus(x):
"""Yield x and -x."""
    yield x
    yield -x

>>> t = plus_minus(3)
>>> next(t)
3
>>> next(t)
-3
>>> list(plus_minus(5))
[5, -5]
>>> list(map(abs, plus_minus(7)))
[7, 7]

```

特点：
- 生成函数是一个产生值 (`yields values`) 而不是返回值的函数。
- 一个普通函数只 `return` 一次，一个生成函数可以 `yield` 多次。
- 生成器是通过调用生成器函数自动创建的迭代器
- 当调用生成器函数时，它返回一个生成器，该生成器对其 `yields` 进行迭代。
- 生成函数和 迭代器 类似，都是 `lazily` 地运算的。因此 `yield` 后在那个框架会暂停

## yield from
生成器有时候也要输出迭代器, `yield from` 语句产生一个 迭代器 或 所有可迭代的值。
  
```python
def a_then_b_for(a, b):
    """The elements of a followed by those of b.

    >>> list(a_then_b_for([3, 4], [5, 6]))
    [3, 4, 5, 6]
    """
    for x in a:
        yield x
    for x in b:
        yield x

def a_then_b(a, b):
    """The elements of a followed by those of b.

    >>> list(a_then_b([3, 4], [5, 6]))
    [3, 4, 5, 6]
    """
    yield from a
    yield from b

def countdown(k):
    """Count down to zero.

    >>> list(countdown(5))
    [5, 4, 3, 2, 1]
    """
    if k > 0:
        yield k
        yield from countdown(k-1)
```