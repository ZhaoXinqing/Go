// golang-map
// https://studygolang.com/articles/23184?fr=sidebar

// Map 是一种无序的键值对的集合，所以我们可以像迭代数组和切片那样迭代它
// 不过，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

// 因为map存在key和value，所以range就有两种写法:
countryCapitalMap := map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo"}

for country := range contryCapitalMap{
	fmt.Println("Capital of ",country, "is",countryCapitalMap[country])
}

for country,capital := range countryCaptalMap{
	fmt.Println("Capital of ",country,"is",capital)
}

// golang里只有一个hashmap
// map不需要指定大小，因为map会自动扩容
// 创建一个map:
myMap := make(map[string]int)
// 另一种方式创建
var myMap map[string]int
myMap = make(map[string]int)

// map在日常使用过程当中，有两个操作经常出现，一判断指定Key是否存在。二删除指定Key。
// 先来回答第一个问题，如果判断指定Key是否存在:
capital, ok := countryCapitalMap["United States"]
if(ok){
	fmt.Println("Capital of United States is", capital)
	} else {
	fmt.Println("Capital of United States is not present")
}
// ok是个bool值，用来表示指定的Key是否存在，简单明了
// 第二个问题，如何删除？
delete(countryCapitalMap,"France")
