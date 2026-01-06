package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID    int    `orm:"id"`
	Name  string `orm:"username"`
	Email string // 默认使用 email
	IgnoreField string `orm:"-"` // 应该忽略
}

// InsertSQL 生成 INSERT 语句和参数
func InsertSQL(v interface{}) (string, []interface{}, error) {
	// TODO: 实现反射逻辑
	// 1. 检查 v 是否为结构体或结构体指针
	// 2. 获取表名（结构体名）
	// 3. 遍历字段，获取列名和值
	// 4. 拼接 SQL 字符串
	
	return "", nil, nil
}

func main() {
	u := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		IgnoreField: "ignore me",
	}

	sql, args, err := InsertSQL(u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("SQL:", sql)
	fmt.Println("Args:", args)
	
	// 预期输出:
	// SQL: INSERT INTO User (id, username, email) VALUES (?, ?, ?)
	// Args: [1 Alice alice@example.com]
}
