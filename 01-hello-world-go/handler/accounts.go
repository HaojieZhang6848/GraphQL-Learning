package handler

import (
	"01-hello-world-go/db"
	"01-hello-world-go/models"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/graphql-go/graphql"
)

// ResolveAccounts 定义查询所有账户的处理函数
var ResolveAccounts = func(p graphql.ResolveParams) (interface{}, error) {
	return db.Accounts, nil
}

// ResolveAccountsByOwner 定义根据所有者查询账户的处理函数
var ResolveAccountsByOwner = func(p graphql.ResolveParams) (interface{}, error) {
	owner := p.Args["owner"].(string)
	ret := []interface{}{}
	for _, account := range db.Accounts {
		if account.Owner == owner {
			ret = append(ret, account)
		}
	}
	return ret, nil
}

// ResolveCreateAccount 定义创建账户的处理函数
var ResolveCreateAccount = func(p graphql.ResolveParams) (interface{}, error) {
	owner := p.Args["input"].(map[string]interface{})["owner"].(string)
	amount := p.Args["input"].(map[string]interface{})["amount"].(float64)
	status := models.AccountStatusNormal
	account := models.Account{
		ID:     gofakeit.UUID(),
		Owner:  owner,
		Amount: amount,
		Status: status,
	}
	db.Accounts = append(db.Accounts, account)
	return account, nil
}

// ResolveUpdateAccount 定义更新账户的处理函数
var ResolveUpdateAccount = func(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["input"].(map[string]interface{})["id"].(string)
	amount := p.Args["input"].(map[string]interface{})["amount"].(float64)
	status := p.Args["input"].(map[string]interface{})["status"].(string)
	owner := p.Args["input"].(map[string]interface{})["owner"].(string)
	for i, account := range db.Accounts {
		if account.ID == id {
			db.Accounts[i].Amount = amount
			db.Accounts[i].Status = status
			db.Accounts[i].Owner = owner
			return db.Accounts[i], nil
		}
	}
	return nil, fmt.Errorf("account with id %s not found", id)
}
