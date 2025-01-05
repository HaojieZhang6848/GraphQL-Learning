# GraphQL Learning 

这个项目是用来学习GraphQL的，主要是通过实践来学习GraphQL的使用。使用Golang作为后端，实现一个GraphQL服务，前端使用ApiFox来测试GraphQL的接口。

下面描述了一些GraphQL的查询(query)和变更(mutation)。如果查询或变更需要参数，这些参数会在variables中描述。最后，展示了curl命令以描述如何向GraphQL服务发送请求，以执行查询或变更。

## 查询books(query)

**query**
```graphql
query {
    books {
      author
      id
      title
    }
}
```

**实际发送的请求(下同)**
```shell
curl --location --request POST 'http://localhost:9080/graphql' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: localhost:9080' \
--header 'Connection: keep-alive' \
--data-raw '{"query":"query {\n    books {\n      author\n      id\n      title\n    }\n}","variables":{}}'
```

## 查询authors(query)

**query**
```graphql
query {
    authors {
      name
      id
    }
}
```

## 计算斐波那契数列(带参数的query)

**query**
```graphql
query Fibonacci($n: Int!) {
    fibonacci(n: $n)
}
```

**variables**
```json
{
    "n": 30
}
```

**实际发送的请求(下同)**
```shell
curl --location --request POST 'http://localhost:9080/graphql' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: localhost:9080' \
--header 'Connection: keep-alive' \
--data-raw '{"query":"query Fibonacci($n: Int!) {\n    fibonacci(n: $n)\n}","variables":{"n":30}}'
```

## 查询employees(带参数的嵌套query)

**query**
```graphql
query Employee($department: String, $city: String, $taxRate: Float) {
    employees(department: $department) {
        id
        name
        age
        gender
        department
        salary(city: $city, taxRate: $taxRate)
    }
}
```

**variables**
```json
{
    "department": "IT",
    "city": "Shanghai",
    "taxRate": 0.1
}
```

## 查询accounts(query)

**query**
```graphql
query {
    accounts {
        id
        owner
        amount
        status
    }
}
```

## 根据owner名字查询account(带参数的query)

**query**
```graphql
query AccountsByOwner($owner: String) {
    accountsByOwner(owner: $owner) {
        id
        owner
        amount
        status
    }
}
```

**variables**
```json
{
    "owner": "Alice"
}
```

## 创建account(mutation)

**mutation**
```graphql
mutation CreateAccount($input: CreateAccountInput) {
    createAccount(input: $input) {
        id
        status
        owner
        amount
    }
}
```

**variables**
```json
{
    "input": {
        "owner": "huajuan",
        "amount": "114514"
    }
}
```

**实际发送的请求(下同)**
```shell
curl --location --request POST 'http://localhost:9080/graphql' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: localhost:9080' \
--header 'Connection: keep-alive' \
--data-raw '{"query":"mutation CreateAccount($input: CreateAccountInput) {\n    createAccount(input: $input) {\n        id\n        status\n        owner\n        amount\n    }\n}","variables":{"input":{"owner":"huajuan","amount":"114514"}}}'
```

## 更新account(mutation)

**mutation**
```graphql
mutation UpdateAccount($input: UpdateAccountInput) {
    updateAccount(input: $input) {
        id
        status
        owner
        amount
    }
}
```

**variables**
```json
{
    "input": {
        "id": "bcc9e5b8-75ca-4c55-802a-ad2217ec5070",
        "owner": "chengxi",
        "amount": "20",
        "status": "Normal"
    }
}
```

## 大杂烩查询(带参数的嵌套query)

**query**
```graphql
query HotchPotchQuery($fibN: Int, $department: String, $city: String, $taxRate: Float, $owner: String) {
    books {
      author
      id
      title
    }
    authors {
      name
      id
    }
    fibonacci(n: $fibN)
    employees(department: $department) {
        id
        name
        age
        gender
        department
        salary(city: $city, taxRate: $taxRate)
    }
    accounts {
        id
        owner
        amount
        status
    }
    accountsByOwner(owner: $owner) {
        id
        owner
        amount
        status
    }
}
```

**variables**
```json
{
    "fibN": 20,
    "department": "Engineering",
    "city": "Shanghai",
    "taxRate": 0.2,
    "owner": "chengxi"
}
```