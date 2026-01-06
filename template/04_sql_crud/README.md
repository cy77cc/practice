# 题目 04: 简易 SQL 生成器 (Reflection 实战)

## 背景
ORM (Object Relational Mapping) 框架的核心功能之一是将 Go 的结构体映射为 SQL 语句。这需要大量使用 Go 的反射机制 (`reflect` 包)。

## 需求
实现一个通用的 SQL 生成器函数 `InsertSQL(v interface{}) (string, []interface{}, error)`。

1.  **输入**：任意结构体实例（指针或值）。
    结构体定义示例：
    ```go
    type User struct {
        ID   int    `orm:"id"`
        Name string `orm:"username"`
        Age  int    `orm:"age"`
    }
    ```
2.  **处理**：
    *   使用 `reflect` 包解析结构体的字段名和 Tag (`orm`)。
    *   如果字段没有 `orm` tag，则使用字段名的小写形式作为列名。
    *   提取字段的值。
3.  **输出**：
    *   生成的 SQL 语句，例如：`INSERT INTO User (id, username, age) VALUES (?, ?, ?)`
    *   参数切片：`[]interface{}{1, "Alice", 18}`
    *   表名默认使用结构体名称。

## 考察点
*   `reflect.TypeOf` 和 `reflect.ValueOf`
*   `reflect.StructTag` 解析 (`Get` 方法)
*   处理不同类型的字段 (int, string 等)
*   字符串拼接 (`strings.Builder`)

## 扩展（可选）
*   支持忽略某个字段（例如 tag 为 `orm:"-"`）。
*   实现 `UpdateSQL`，假设第一个字段为主键。
