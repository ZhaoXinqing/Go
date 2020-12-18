package main

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	grom.Model
	Code  string
	Price uint
}

func main() {
	db, err := grom.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询 id 为 1 的 product
	db.First(&product, "code = ?", "L1212") // 查询 code 为 L1212 的 product

	// 更新 （更新product的price为 2000）
	db.Model(&product).Update("Price", 2000)

	// 删除 （删除 product)
	db.Delete(&product)
}
