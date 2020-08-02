
ORM（Object Relational Mapping）对象关系映射：
    - 是什么：
        - 通过创建一个类，将这个类与数据库的表相对应，通过把用户输入的类或对象转换成SQL语句，完成对数据库的操作；
        - 把数据库映射成对象。
            - 数据库的表（table） --> 类（class）
            - 记录（record，行数据）--> 对象（object）
            - 字段（field）--> 对象的属性（attribute）
    - 操作实例：
        - 数据库的基本操作有四种：create（新建）、read（读取）、update（更新）和delete（删除），简称 CRUD。 ORM 将这四类操作，都变成了对象的方法。
        - 查询：
            find()方法用于根据主键，获取单条记录：
                - 返回单条记录- Customer.find(1)
                - 返回多条记录- Customer.find([1, 2, 3])
            where()方法用于指定查询条件：
                - Customer.where({Company: 'Apple Inc.'}).first()                
                - const customers = await Customer;如果直接读取类，将返回所有记录。
            limit(limit[, offset])方法指定返回记录的位置和数量：
                - const customers = await Customer.limit(5, 10);)   从第10条记录开始，返回5条记录
        - 新建：
            create()方法用于新建记录。
                - Customer.create({
                    Email: '',
                    FirstName: 'Donald',
                    LastName: 'Trump',
                    Address: 'Whitehouse, Washington'
                    })
        - 更新：
            update()方法用于更新记录。
                - const customer = await Customer.find(60);
                - await customer.update({
                    Address: 'Whitehouse'
                    });
        - 删除：
            destroy()方法用于删除记录。
                - const customer = await Customer.find(60);
                    - await customer.destroy();
    - 举例：
        - SQL语句：
            - SELECT id, first_name, last_name, phone, birth_date, sex FROM persons WHERE id = 10
                - 程序直接运行 SQL，操作数据库：
                    - res = db.execSql(sql); 
                    - name = res[0]["FIRST_NAME"];
        - 改成ORM ：
            - p = Person.get(10);
            - name = p.first_name;
    - 比较就发现，ORM 使用对象，封装了数据库操作，因此可以不碰 SQL 语言。开发者只使用面向对象编程，与数据对象直接交互，不用关心底层数据库。

    - 优点：
        - 使得开发更加对象化，提高开发效率，降低开发成本；
        - 提高了系统的可移植性，还可以更方便地引入数据库缓存之类的附加功能。
    - 缺点：
        - 从系统结构上看，采用orm的系统多是多层系统的，太多的系统层次，会降低执行效率；并且ORM作为一种完全面对对象的做法，也会对系统性能产生一定的影响；
        - 在处理多表联查、where条件之类的复杂查询时，ORM的语法会变得复杂，就需要写入原生SQL；