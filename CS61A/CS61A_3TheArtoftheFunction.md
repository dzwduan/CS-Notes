# 实践指南：函数的艺术

## 如何编写良好的函数
 - 每个函数都应只做**一个任务**
  - **不要重复劳动**是软件工程的中心法则
  - 函数应**泛型**，最好在任何情况下可以使用

##  文档字符串
描述这个函数的文档，叫做文档字符串
- 使用 `"""`，描述函数的任务和参数
- [文档字符串准则](https://www.python.org/dev/peps/pep-0257/)

```python
def pressure(v, t, n):
    """Compute the pressure in pascals of an ideal gas.

    Applies the ideal gas law: http://en.wikipedia.org/wiki/Ideal_gas_law

    v -- volume of gas, in cubic meters
    t -- absolute temperature in degrees kelvin
    n -- particles of gas
    """
    k = 1.38e-23  # Boltzmann's constant
    return n * k * t / v

help(pressure) # 查看文档字符串
```

## 参数默认值
```python
k_b=1.38e-23  # Boltzmann's constant
def pressure(v, t, n=6.022e23):
    return n * k_b * t / v
```
- 函数体内的常见数据，应表示为具名参数的默认值 （`n=6.022e23`）
- 永远不会改变的数据，应该定义在全局帧中 （基本常数`k_b=1.38e-23`）
  
## python 可以有多个返回值
```python
def divide_exact(n,d):
    return n // d, n % d
quotient, remainder = divide_exact(2013, 10)
>>> quotient
201
>>> remainder 
3
```
- 实际上返回的是一个元组 tuple
