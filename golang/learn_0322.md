go中声明是给一个程序实体命名，并且设定其部分或全部属性，有四个主要的申请：变量（var）、常量（const）、类型（type）和函数（func）。


byte是uint8的别名类型，而rune是int32的别名类型。

Go的数据类型分四大类：基础类型（basic type）、聚合类型（aggregate type）、引用类型（reference type）和接口类型（interface type）。


##  输出

Print:   输出到控制台(不接受任何格式化，它等价于对每一个操作数都应用 %v)
         fmt.Print(str)
         
Println: 输出到控制台并换行
         fmt.Println(tmp)
         
Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出整形变量和整形 等）
         fmt.Printf("%d",a)
         
Sprintf：格式化并返回一个字符串而不带任何输出。
         s := fmt.Sprintf("a %s", "string") fmt.Printf(s)
         
Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
         fmt.Fprintf(os.Stderr, “an %s\n”, “error”)

