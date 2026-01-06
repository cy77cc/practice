package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID          int    `orm:"id"`
	Name        string `orm:"username"`
	Email       string // 默认使用 email
	IgnoreField string `orm:"-"` // 应该忽略
}

// InsertSQL 生成 INSERT 语句和参数
func InsertSQL(v interface{}) (string, []interface{}, error) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("unsupported type")
	}
	tableName := t.Name()
	columns := make([]string, 0, t.NumField())
	args := make([]interface{}, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("orm") == "-" {
			continue
		}
		col := f.Tag.Get("orm")
		if col == "" {
			col = strings.ToLower(f.Name)
		}
		columns = append(columns, col)
		args = append(args, val.Field(i).Interface())
	}
	if len(columns) == 0 {
		return "", nil, fmt.Errorf("no columns")
	}
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	return sql, args, nil
}

func main() {
	u := User{
		ID:          1,
		Name:        "Alice",
		Email:       "alice@example.com",
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
