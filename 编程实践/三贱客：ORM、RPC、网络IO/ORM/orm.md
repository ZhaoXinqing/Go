
ORM（Object Relational Mapping）对象关系映射：
    - 是什么：
        - 面向对象编程把所有实体看成对象（object），关系型数据库则是采用实体之间的关系（relation）连接数据。而ORM 就是通过使用描述对象和数据之间映射的元数据，完成关系型数据库的操作的技术
    - 优点：
        - 提高开发效率，降低开发成本；
        - 使得开发更加对象化；
        - 可移植；
        - 可以更方便地引入数据库缓存之类的附加功能。
    - 缺点：
        - 降低程序的执行效率，从系统结构上看，采用orm的系统多是多层系统的，系统的层次太多，效率就会降低，orm是一种完全面对对象的做法，所以面向对象的做法也会对性能产生一定的影响；
        - 在处理多表联查、where条件之类的复杂查询时，ORM的语法会变得复杂，就需要写入原生SQL

    - 框架及原理：
        - 通过创建一个类，将这个类与数据库的表相对应，通过把用户输入的类或对象转换成SQL语句，完成对数据库的操作；
        - 把数据库映射成对象。
            - 数据库的表（table） --> 类（class）
            - 记录（record，行数据）--> 对象（object）
            - 字段（field）--> 对象的属性（attribute）

    - CRUD 操作实例：
        - 数据库的基本操作有四种：create（新建）、read（读取）、update（更新）和delete（删除），简称 CRUD。 ORM 将这四类操作，都变成了对象的方法。
        - 查询
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
        - 新建记录
            create()方法用于新建记录。
            // demo12.js
            Customer.create({
             Email: '',
            FirstName
            : 'Donald',
             LastName: 'Trump',
             Address: 'Whitehouse, Washington'
            })
        - 更新记录
            update()方法用于更新记录。
            // demo13.js
            const customer = await Customer.find(60);
            await customer.update({
             Address: 'Whitehouse'
            });
        - 删除记录
            destroy()方法用于删除记录。
            // demo14.js
            const customer = await Customer.find(60);
            await customer.destroy();
            https://www.toutiao.com/a6660637718638232078/

举例来说，下面是一行 SQL 语句。
SELECT id, first_name, last_name, phone, birth_date, sex FROM persons WHERE id = 10
程序直接运行 SQL，操作数据库的写法如下。
res = db.execSql(sql);

name = res[0]["FIRST_NAME"];
改成 ORM 的写法如下。
p = Person.get(10);
name = p.first_name;
一比较就可以发现，ORM 使用对象，封装了数据库操作，因此可以不碰 SQL 语言。开发者只使用面向对象编程，与数据对象直接交互，不用关心底层数据库。