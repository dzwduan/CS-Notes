# Non-Local Assignment
在函数内部创建的变量，只在自己的 `Frame` 里使用。  
使用高阶函数时，内部函数想要无法引用/改变 `parent frame` 的变量。

解决方法：使用 `nonlocal` 声明变量使内部函数的变量绑定为 `parent frame` 的变量。

```python
def make_withdraw(balance):
    "Return a withdraw function with a starting balance."
    def withdraw(amount):
        nonlocal balance # declare the name "balance" nonlocal
        if amount > balance:
            return "Insufficient funds"
        balance -= amount
        return balance
    return withdraw

withdraw = make_withdraw(100)
withdraw(25)
withdraw(25)
wtihdraw(60)

```
## Non-Local 细节

对经过 `nonlocal` 声明过的 `name` 重新赋值，改变的是 `name` 的第一个非本地框架 (**first non-local framework**) 中预先存在的值
- 该 `name` 要已经声明过
- 该 `name` 不能与 `local scope` 的 `name` 冲突


|状态(x=2)|影响|
|:---:|:---:|
|没有nonlocal,x 没有在本地绑定|在当前环境的第一个 frame 中,创建一个 x 到对象 2 的新绑定|
|没有nonlocal,x 已经在本地绑定过|在当前环境的第一个 frame 中,将 x 重新绑定到对象 2 |
|有nonlocal x,x 已经在 non-local frame 绑定过|在当前环境的第一个 non-local frame 中,将 x 重新绑定到对象 2 |
|有nonlocal x,x 没有在 non-local frame 绑定过|语法错误|
|有nonlocal x,x 已经在 non-local frame 绑定过, x也在本地绑定过|语法错误|


## Python particulars(细节)

python 在执行方法前，先求出 `name` 属于的 `frame`

- 在 `function` 内部，单一 `name` 的所有实例必须指向同一个框架

  
```python
def make_withdraw(balance):
    """Return a withdraw function with a starting balance."""
    def withdraw(amount):
        if amount > balance:
            return 'Insufficient funds'
        balance = balance - amount  # 执行前计算，得知 balance 是本地变量
        return balance
    return withdraw
```

```python
def make_withdraw(balance):
    """Return a withdraw function with a starting balance."""
    def withdraw(amount):
        if amount > balance:
            return 'Insufficient funds'
        # balance = balance - amount # 如果去掉，可以正常执行，因为这个 name 指向前一个 frame
        return balance
    return withdraw

```

其他可选的方法：  
用可变值，如列表。

```python
def make_withdraw_list(balance):
    b = [balance]
    def withdraw(amount):
        if amount > b[0]:
            return 'Insufficient funds'
        b[0] = b[0] - amount
        return b[0]
    return withdraw
```


## 突变会导致 函数的引用透明性 (referential transparency) 丢失
返回值只依赖于输入值，这种特性称为引用透明性。  
`Mutation operations` 违反引用透明度，因为它们改变了环境。