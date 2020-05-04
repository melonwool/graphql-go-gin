# graphql-go-gin
gin graphql demo

### 说明
由于不想安装其他数据库，db目前使用的sqllite, 可以随意更换其他db

### 开始
启动
```
go run server.go   // 0.0.0.0:8080 
```
测试地址
http://ip:8080/ 使用了[GraphiQL](https://github.com/graphql/graphiql)
### 实现的功能
- 查询
- 插入
- 修改
- 删除
- TODO 分页查询
### 示例
通过用户ID查询数据

```
{
  getUser(id: 1) {
    name
    books {
      name
      id
    }
  }
}
```

通过书籍ID查询数据

```
{
  getBook(id: 1) {
    name
    owner {
      name
      id
    }
  }
}
```

添加书籍

```

mutation addBook($book: BookInput!) {
  addBook(book: $book) {
    name
  }
}

// query variables
{
  "book": {
    "name": "《笑傲江湖》",
    "ownerID": 1
  }
}

```

修改书籍

```
mutation UpdateBook($book: BookInput!) {
  updateBook(book: $book) {
    name
    id
  }
}
// query variables
{
  "book": {
    "id": "1",
    "name": "《笑傲江湖》",
    "ownerID": 1,
    "tagIDs": [1]
  }
}
```

删除书籍
```
mutation DeleteBook($userID: ID!, $bookID: ID!) {
  deleteBook(userID:$userID, bookID: $bookID) {

  }
}
// query variables
{
  "bookID":"1",
  "userID":"1"
}
```
