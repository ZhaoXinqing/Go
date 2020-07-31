
## ORM的介绍 ##
ORM的全称是：Object Relational Mapping（对象 关系 映射），简单的说，orm是通过使用描述对象和数据之间映射的元数据，将程序中的对象自动持久化到关系数据库中。
ORM需要解决两个问题，能否把对象的数据直接保存到数据库中，又能否直接从数据库中拿到一个对象？要想做到上面两点，则必须要有映射关系。
优点：提高开发效率。可以自动对实体Entity对象与数据库中的Table进行字段与属性的映射；不用直接SQL编码，能够像操作对象一样从数据库获取数据；（提高开发效率，降低开发成本；是开发更加对象化；可移植；可以更方便地引入数据库缓存之类的附加功能）
缺点：会牺牲程序的执行效率和会固定思维模式，在从系统结构上来看，采用orm的系统多是多层系统的，系统的层次太多，效率就会降低，orm是一种完全面对对象的做法，所以面向对象的做法也会对性能产生一定的影响；（在处理多表联查、where条件复杂之类的查询时，ORM的语法会变得复杂。就需要写入原生SQL）
总结：ORM只是一种帮助我们解决一些重复的、简单的劳动，我们不能一劳永逸的靠工具来解决问题，有些特殊问题还是需要进行特殊处理的。
ORM框架：对象关系映射，通过创建一个类，这个类与数据库的表相对应！类的对象代指数据库中的一行数据。
简述ORM原理：让用户不再写SQL语句，而是通过类以及对象的方式，和其内部提供的方法，进行数据库操作！把用户输入的类或对象转换成SQL语句，转换之后通过pymysql执行完成数据库的操作。
Orm对象关系映射
ORM（Object Relational Mapping）：对象关系映射，通过创建一个类与数据库的表相对应，类的对象代指数据库中的一行数据。让用户不再写SQL语句，而是通过类以及对象的方式，和其内部提供的方法，进行数据库操作。
功能：把用户输入的类或对象转换成SQL语句，转换之后通过pymysql执行完成数据库的操作。
优点：提高开发效率，降低开发成本；使开发更加对象化；可移植；方便地引入数据缓存之类的附加功能
不足之处：在处理多表联查、where条件之类的复杂查询时，ORM的语法会变得复杂，就需要写入原生SQL
面向对象编程和关系型数据库，都是目前最流行的技术，但是它们的模型是不一样的。面向对象编程把所有实体看成对象（object），关系型数据库则是采用实体之间的关系（relation）连接数据。很早就有人提出，关系也可以用对象表达，这样的话，就能使用面向对象编程，来操作关系型数据库。
简单说，ORM 就是通过实例对象的语法，完成关系型数据库的操作的技术，是"对象-关系映射"（Object/Relational Mapping） 的缩写。
ORM 把数据库映射成对象。
数据库的表（table） --> 类（class）
记录（record，行数据）--> 对象（object）
字段（field）--> 对象的属性（attribute）
 
举例来说，下面是一行 SQL 语句。
SELECT id, first_name, last_name, phone, birth_date, sex FROM persons WHERE id = 10
程序直接运行 SQL，操作数据库的写法如下。
res = db.execSql(sql);

name = res[0]["FIRST_NAME"];
改成 ORM 的写法如下。
p = Person.get(10);
name = p.first_name;
一比较就可以发现，ORM 使用对象，封装了数据库操作，因此可以不碰 SQL 语言。开发者只使用面向对象编程，与数据对象直接交互，不用关心底层数据库。

## ORM 优点 ##
1.数据模型都在一个地方定义，更容易更新和维护，也利于重用代码。
2.ORM 有现成的工具，很多功能都可以自动完成，比如数据消毒、预处理、事务等等。
3.它迫使你使用 MVC 架构，ORM 就是天然的 Model，最终使代码更清晰。
4.基于 ORM 的业务代码比较简单，代码量少，语义性好，容易理解。
5.你不必编写性能不佳的 SQL。

## ORM 缺点 ##
1.ORM 库不是轻量级工具，需要花很多精力学习和设置。
2.对于复杂的查询，ORM 要么是无法表达，要么是性能不如原生的 SQL。
3.ORM 抽象掉了数据库层，开发者无法了解底层的数据库操作，也无法定制一些特殊的 SQL。
CRUD 操作
数据库的基本操作有四种：create（新建）、read（读取）、update（更新）和delete（删除），简称 CRUD。
ORM 将这四类操作，都变成了对象的方法。
6.1 查询
前面已经说过，find()方法用于根据主键，获取单条记录（完整代码看这里）或多条记录（完整代码看这里）。
// 返回单条记录
// demo02.js
Customer.find(1)
// 返回多条记录
// demo05.js
Customer.find([1, 2, 3])
where()方法用于指定查询条件。
// demo04.js
Customer.where({Company: 'Apple Inc.'}).first()
如果直接读取类，将返回所有记录。
// 返回所有记录
const customers = await Customer;
但是，通常不需要返回所有记录，而是使用limit(limit[, offset])方法指定返回记录的位置和数量（完整代码看这里）。
// demo06.js
const customers = await Customer.limit(5, 10);)
上面的代码制定从第10条记录开始，返回5条记录。
6.2 新建记录
create()方法用于新建记录。
// demo12.js
Customer.create({
 Email: '',
FirstName
: 'Donald',
 LastName: 'Trump',
 Address: 'Whitehouse, Washington'
})
6.3 更新记录
update()方法用于更新记录。
// demo13.js
const customer = await Customer.find(60);
await customer.update({
 Address: 'Whitehouse'
});
6.4 删除记录
destroy()方法用于删除记录。
// demo14.js
const customer = await Customer.find(60);
await customer.destroy();
https://www.toutiao.com/a6660637718638232078/