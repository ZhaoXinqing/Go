## 一、变量类型 ##
变量分为值类型，指针类型和引用类型。

- 在golang中故意淡化了指针的概念，我们只需要关注值类型和引用类型就可以。你在官方介绍中也很少看到指针类型这一概念

## 二、值类型和引用类型的区别 ##
- 值类型变量：除开slice,map, channel类型之外的变量都是值类型
- 引用类型变量：slice, map, channel这三种。
### 1，零值不同 ###
- 指针类型的变量，零值都是nil。
- 值类型的变量，零值是其所在类型的零值。
	1. int32类型的零值是0
	1. string类型的零值是""
	1. bool类型的零值是false
	1. 符合结构struct类型的零值是其每个成员的零值的组合

###2，变量申明后是否需要初始化才能使用
- 指针类型的变量，需要初始化才能使用。(slice是一个特例，slice的零值是nil，但是可以直接append)
- 值类型的变量，不用初始化，可以直接使用（struct声明后可以直接使用直接赋值，但是map就不行）
###3，初始化方法不同

- 值类型的变量，其实不需要初始化就可以使用。如果有良好的代码习惯，使用前进行初始化也是非常提倡的。
	- 基本类型的初始化非常简单:
		- var i int; i = 1;
		- var b bool; b = true
		- var s string; s = ""
	- 符合类型struct的初始化有两种:
		- s1 = Student1{}
		- s1 = new(Student1)


- 引用类型的变量，初始化方式也不一样：
	- slice类型，用make，new，{}都可以
	- map类型，用make，new，{}都可以
	- channel类型，只能用make活着new初始化

	

	//map可以用{}，make，new三种方式初始化

	- s2 := Student2{} //map[]
	- s2 := make(Student2) //map[]
	- s2 := new(Student2) //&map[]

	//slice可以用{},make,new,但是make的时候需要带len参数
	- type S3 []string
	- s3 := S3{} //[]
	- s3 := new(S3) //&[]
	- s3 := make(S3, 10) //[ ]

	//channel只能用make或者new

	- type Student4 chan string
	- s4 := new(S4) //0xc000096000
	- s4 := make(S4) //0xc000082008
	- s4 := S4{} //编译器报错：invalid type for composite literal: Student4

##三、make和new的区别


- make返回的是对象。
	- 对值类型对象的更改，不会影响原始对象的值
	- 对引用类型对象的更改，会影响原始对象的值
- new返回的是对象的指针，对指针所在对象的更改，会影响指针指向的原始对象的值。

##四、golang没有引用传递，都是值传递

- 如果函数形参是值类型，则会对值类型做一份拷贝作为函数形参。在函数内对形参变量做的修改，不会影响函数外的那个被传入的变量。
- 如果函数形参是引用类型，则会对引用类型变量做一次拷贝。但是拷贝得到的引用类型变量的值，和被传入调用函数的原始引用类型变量的值，是一样的，即指向的是同一个变量的地址(参考前面值类型变量和引用类型变量图)。所以在函数里面的修改，会影响原始引用变量指向的变量的值。

   type Student struct{
    Age int
	}
   
   
	func main(){
    s := Student{30}
    modify1(s)
    	fmt.Println(s) //{30}
    	
    modify2(&s)
    fmt.Println(s) //{32}
     }
    
     func modify1(s Student){
    s.Age = 31
     }
    
     func modify2(s *Student){
    s.Age = 32
     }