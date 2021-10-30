# monkey-test

该项目是一个测试项目，用于个人独立学习。

测试对象：gomonkey

库介绍：用于打桩测试使用的go mock库

项目包含内容：

- 测试文件json_test，mock一个函数；
- 测试文件method_test，mock一个方法；
- 测试文件global_var_test，mock一个全局变量；
- 测试文件compute_test，mock一个函数序列；
- 测试文件method_seq_test,mock一个方法序列。


### 库基本介绍

> github.com/agiledragon/gomonkey

- support a patch for a function
- support a patch for a public member method
- support a patch for a private member method
- support a patch for a interface
- support a patch for a function variable
- support a patch for a global variable
- support patches of a specified sequence for a function
- support patches of a specified sequence for a member method
- support patches of a specified sequence for a interface
- support patches of a specified sequence for a function variable

由于go编译器的优化功能，会将简单地函数置换成内置函数，会导致mock测试可能失效。
为了避免这种情况需要在测试的时候关闭编译器的内置函数优化功能。


> github.com/smartystreets/goconvey

这是一款go测试工具，在此项目中利用到了该库的各种测试辅助方法。

### 项目运行

示例：

```shell
cd pkg/compute
go test gcflags=all=-l
```

### 基本使用介绍

- Convey

该函数起到作用域的效果，能够反复嵌套，在该函数中mock的对象将会在该作用域和子作用域生效。能够隔离mock带来的负面影响。

官方推荐在在顶层使用`Convey("test func name", *testing.T, func(){})`包裹。


