// 多元判断，switch的语法看起来就会比If .. else if 要简洁。
// switch <表达式> {
// 	case <值1>:
// 	case<值2>:
// 	default:
// }
// switch判断不需要添加break
// default是可选项，如果都匹配不成功，没有异常，没有报错。所以有可能程序出现非预期结果
// 表达式部分必须是一个完整的布尔表达式，或者是一个推算出唯一结果的函数。
// 如果表达式为空，或者计算失败，则默认为True。
// 所有的case必须是同一种数据类型，如果类型不一致，就有可能出现非预期的计算结果
// 通过fallthrough可以跳过case判断。

// Select
// select后面没有表达式
// case后面的类型必须是chan。(chan是一种数据通信管道，golang封装成了数据类型)
// default是可选的。如果有default，必须放到select最后。

// Golang中有几种退出循环的方式
// Break。 使用率最高的方式，当满足退出条件时，直接break退出
// Continue。跳出了本次循环，继续下次循环
// Goto。 武林界失传已久的Goto大法，在Golang中重出江湖。 与其它语言，强制半强制的屏蔽Goto不同，
// golang很大方的推出goto。没有什么其它限制条件，只要你认为goto使用的正确就可以大胆的使用。(小心陷
// 入死循环)。使用Goto时，需要提前定义标签，然后goto 标签即可。
