# data abstraction
复合数据 可以看做 单个概念单元

解决两个问题
  - 如何表示数据
  - 如何处理数据

## 分数是用 一对数字 来表示的经典例子

```python
>>> pair = [1,2]
>>> pair
[1,2]

>>> x,y = pair # unpacking a list
>>> x
1
>>> y
2
```

## 复合数据体现的抽象原则
抽象屏障 (Abstraction Barriers)
  - 选择适合的抽象层方法，不要用不属于自己抽象层的方法
  
优点：
  - 使得程序更容易被更改
  - 可以改变 data representation 无需重写整个程序

例如：
```python
def rational_mul(x,y):
    return rational(numer(x) * numer(y), denom(x) * denom(y))

#---------用列表表示-------------------

def rational(x,y):
    return [x,y]

def numer(x):
    return x[0]

def denom(x):
    return x[1]

#---------用HOF表示-------------------

def rational(x,y):
    def select(keyword):
        if keyword == "n":
            return x
        if keyword == "d":
            return y
    return select

def numer(x):
    return x("n")

def denom(x):
    return x("d")
```

## dict
特点
- 无顺序
- 键值对
- 两个 key 不能相同
- key 不能是列表


```python
>>> numberals = {'I':1,'V':5,'X':10}    # dict 没有顺序
>>> numberals
{'I': 1, 'V': 5, 'X': 10}

#----- keys() --------
>>> numberals.keys()
dict_keys(['I', 'V', 'X'])

#----- values() ------
>>> numberals.values()
dict_values([1, 5, 10])

#------ items() ------
>>> numberals.items()
dict_items([('I', 1), ('V', 5), ('X', 10)])

#---- list -> dict ------
>>> items = [('I', 1), ('V', 5), ('X', 10)]
>>> dict(items)
{'I': 1, 'V': 5, 'X': 10}

#--- 根据 key 获取 value ------
>>> dict(items)['X']
10

#--------- in -----------
>>> 'X' in numberals
True
>>> 'X-ray' in numberals
False

#---- get(key,default) ---------
>>> numberals.get('X',0)
10
>>> numberals.get('X-ray',0)
0

#------ dict 推导式 ------------
>>> {x:x*x for x in range(5)}
{0: 0, 1: 1, 2: 4, 3: 9, 4: 16}

# ---- 两个 key 不能相同 --------
>>> {1:2,1:3}
{1: 3}

# ----- value 能是列表 --------
>>> {1:[2,3]}
{1: [2, 3]}

# ----- key 不能是列表 --------
>>> {[1]:2}
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: unhashable type: 'list'
```

## 如何知道违反了数据抽象原则
简而言之，违反了数据抽象是指绕过 ADT 的构造函数和选择器，直接使用其在其余代码中的实现方式，并假设其实现不会更改。
- 我们不能假设我们知道 ADT 是如何构造的，只能使用构造函数。
- 我们不能假设我们知道如何访问 ADT 的详细信息，除非使用选择器。 
- 细节应该由构造函数和选择器抽象出来。 

如果我们绕过构造函数和选择器，直接访问细节，那么对 ADT 实现的任何微小更改都可能破坏整个程序。

## 为什么要保持数据抽象
通过正确实现这些数据类型，可以使代码更具可读性，因为：
- 可以使构造函数和选择器的名称更具信息性
- 让团队协作更轻松
- 其他程序员不必担心实现细节
- 防止错误的传播
  - 可以在单个函数中修复错误，而不是在整个程序中修复错误
