package models

// 定义作者的数据类型
type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 定义书籍的数据类型
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// 定义Employee的数据类型
type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
}
