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

// 定义Account的数据类型
type Account struct {
	ID     string  `json:"id"`
	Owner  string  `json:"owner"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

const AccountStatusNormal = "Normal"
const AccountStatusClosed = "Closed"
const AccountStatusSuspended = "Suspended"
const AccountStatusPending = "Pending"
const AccountStatusUnknown = "Unknown"
const AccountStatusError = "Error"

var AccountAllPossibleStatus = []string{
	AccountStatusNormal,
	AccountStatusClosed,
	AccountStatusSuspended,
	AccountStatusPending,
	AccountStatusUnknown,
	AccountStatusError,
}
