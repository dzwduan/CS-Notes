# Inheritance (继承)

`Python` 是对象系统
- `Functions` 是 对象
- `Bonds methods` 也是对象 (第一个参数 `self` 已经绑定给实例的方法)
- 点式表达式求类属性的 `Bonds methods` 是 `functions`。
    
    
点式表达式 `<instance>.<method_name>`
- 左边生成一个 `object`
- 右边在 `instance` 属性内部找 `name`
- 如果没有，找类属性
- 该值将返回，除非它是一个函数(会返回绑定方法)
  
## 继承
继承是一种将类联系在一起的技术。子类可以具有与父类类相同的属性，以及一些特殊情况下的行为。 
- 两个子类可以在*定制化*的程度上有差异。

```python
class <Name>(<Base Class>):
    <suite>
```

利用继承，我们通过指定子类与父类的不同之处来实现子类。

## 在 classes 中寻找属性
步骤：
在子类中查看该属性，有则返回这个 attribute value。如果没有则去父类中找。

```python
class CheckingAccount(Account):
    """A bank account that charges for withdrawals."""
    withdraw_fee = 1
    interest = 0.01
    def withdraw(self, amount):
        return Account.withdraw(self, amount + self.withdraw_fee)

>>> ch = CheckingAccount('Tom') # Calls Account.__init__
>>> ch.interest                 # Found in CheckingAccount
0.01
>>> ch.deposit(20)              # Found in Account
20
>>> ch.withdraw(5)              # Found in CheckingAccount
14


```

## 面向对象的设计
原则：
- 不要重复自己，使用已经存在的实现方式
- 尽可能在实例上查找属性

## 继承和构成
继承最好用在 is-a 的关系上。例如，支票账户是一种特殊的账户类型。

构成最适合表示 has-a 关系。例如，一家银行管理着一系列的银行账户。
    
## 多重继承
子类可以有多个基类
```class SubclassName(BaseClass1, BaseClass2, BaseClass3, ...):```
- 先调用 subclass ，在调用 baseclass

多重继承会使得程序变得很复杂，因此应该极少使用
