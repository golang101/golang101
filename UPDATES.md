本页内容转移到了[gfw.go101.org](https://gfw.go101.org)网站：

* [Go语言101博客](https://gfw.go101.org/blog/101.html)
* [Go语言101在1.18版本之前的更新历史](https://gfw.go101.org/blog/2022-02-22-history.html)

### 1.16.b (2021/五月/18)

* 增加Go 1.17内容。

### 1.16.a (2021/二月/18)

* 支持Go 1.16中新增的资源内嵌功能。

### 1.15.b (2020/九月/28)

* 对[reflect.DeepEqual](https://gfw.go101.org/article/details.html#reflect-deep-equal)相关细节增加了一些描述。
* 添加了[一个语法例外](https://gfw.go101.org/article/exceptions.html#code-block-following-else).

### 1.15.a (2020/八月/07)

* 指出从Go工具链1.15开始，使用make+copy总是比使用append来克隆切片更快。

### 1.14.g (2020/六月/12)

* 删除了1.14.e中增加的一个细节：比较同时含有可比较和不可比较字段的结构体值或者同时含有可比较和不可比较字段的数组值的行为是未定义的。原因是[此行为将被明确下来](https://github.com/golang/go/issues/8606)。
* 加了[一个细节](https://gfw.go101.org/article/details.html#impossible-to-interface-assertion)：关于在编译时刻即可确定总是失败的目标类型为接口类型的断言。

### 1.14.f (2020/六月/02)

* 所用的"Go SDK"用词都替换成了"Go官方工具链"。

### 1.14.e (2020/五月/06)

* 加了[一个细节](https://gfw.go101.org/article/details.html#compare-values-with-both-comparable-and-incomparable-parts)：比较同时含有可比较和不可比较字段的结构体值或者同时含有可比较和不可比较字段的数组值的行为是未定义的。
* 加了[一个细节](https://gfw.go101.org/article/details.html#blank-fields-are-ignored-in-comparisons)：在结构体值的比较中，名为空标识符的字段将被忽略。

### 1.14.d (2020/四月/25)

* `runtime.KeepAlive`相关的内容被从[非类型安全指针](https://gfw.go101.org/article/unsafe.html)一文中删除了。
  我很抱歉之前在此文中传播了错误的信息。

### 1.13.i (2019/十月/31)

* 改正了“数组、切片和映射”一文中对[删除一段切片元素](https://gfw.go101.org/article/container.html#delete-slice-elements)一节中的错误代码。
* 更正了[延迟调用的函数值的估值时刻](https://gfw.go101.org/article/function.html#function-evaluation-time)一节中的解释。
* "在正确的位置调用内置<code>recover</code>函数"一文[改名](https://gfw.go101.org/article/panic-and-recover-more.html)为“详解panic/recover原理”。这篇文章几乎被整个重写了。

### 1.13.h (2019/十月/18)

* 修正了“表达式估值顺序规则”一文中对[赋值语句中的表达式估值和赋值执行顺序](https://gfw.go101.org/article/evaluation-orders.html#value-assignment)的欠妥解释。
* 添加了[两条总结](https://gfw.go101.org/article/101.html#compiler-optimizations)。

### 1.13.e (2019/十月/07)

* 我决定撤回1.13.d中的勘误。（重新在1.14.d中加入）

### 1.13.d (2019/九月/30)

* 在“非类型安全指针”一文中添加了<a href="https://gfw.go101.org/article/unsafe.html#fact-value-address-might-change">一个事实</a>并指出此文犯了一个<a href="https://gfw.go101.org/article/unsafe.html#pattern-convert-to-uintptr-and-back">严重错误</a>。
  
### 1.13.c (2019/九月/25)

* 删除了《在正确的位置调用内置recover函数》一文中犯了低级错误的一节。

### 1.13.b (2019/九月/19)

* 删除了不准确的描述：一个变量的地址永不改变。

### 1.13.a (2019/九月/05)

* Go 1.13就绪
* [如何优雅地关闭通道](https://gfw.go101.org/article/channel-closing.html)一文中添加了两种情形。
* [“有名类型”和“无名类型”术语](https://gfw.go101.org/article/type-system-overview.html#unnamed-type)重新加了回来，但是它们现在等同于“定义类型”和“非定义类型”。
* “数据通道”被更名为“通道”。

### 1.12.d (2019/四月/18)

* 丰富了[包级变量初始化顺序](https://gfw.go101.org/article/evaluation-orders.html#package-level-variables)一节。

### 1.12.c (2019/四月/09)

* 删除了“有名类型”和“无名类型”术语。

### 1.12.b (2019/四月/06)

* 添加了[包级变量初始化顺序](https://gfw.go101.org/article/evaluation-orders.html#package-level-variables)一节。

