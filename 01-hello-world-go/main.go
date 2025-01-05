package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"01-hello-world-go/types"
	"github.com/graphql-go/graphql"
)

// 定义GraphQL Schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    types.GQLQueryType,
	Mutation: types.GQLMutationType,
})

// 处理GraphQL请求
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 从http body中读取gra
	var mmm map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&mmm); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 从请求中获取查询和变量
	query := mmm["query"].(string)
	var variables map[string]interface{} = nil
	if _, ok := mmm["variables"]; ok {
		variables = mmm["variables"].(map[string]interface{})
	}

	// 执行GraphQL查询
	params := graphql.Params{Schema: schema, RequestString: query, VariableValues: variables}
	result := graphql.Do(params)

	// 返回结果
	if len(result.Errors) > 0 {
		log.Printf("graphql error: %v", result.Errors)
		http.Error(w, fmt.Sprintf("graphql error: %v", result.Errors), http.StatusInternalServerError)
		return
	}

	// 设置响应格式为JSON
	w.Header().Set("Content-Type", "application/json")
	// 将查询结果转换为JSON并返回
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 设置GraphQL请求的路由
	http.HandleFunc("/graphql", graphqlHandler)

	// 启动HTTP服务器
	fmt.Println("GraphQL server is running at http://localhost:9080/graphql")
	log.Fatal(http.ListenAndServe(":9080", nil))
}
