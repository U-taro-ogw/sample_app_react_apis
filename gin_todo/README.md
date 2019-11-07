# go_todo
golang + gin + gormを使った基本的なCRUD（API）


## route
```
[GIN-debug] GET    /todos                    --> github.com/U-taro-ogw/gin_todo/src/controllers.(*TodoHandler).GetAll-fm (3 handlers)
[GIN-debug] GET    /todos/:id                --> github.com/U-taro-ogw/gin_todo/src/controllers.(*TodoHandler).GetTodo-fm (3 handlers)
[GIN-debug] POST   /todos                    --> github.com/U-taro-ogw/gin_todo/src/controllers.(*TodoHandler).CreateTodo-fm (3 handlers)
[GIN-debug] DELETE /todos/:id                --> github.com/U-taro-ogw/gin_todo/src/controllers.(*TodoHandler).DeleteTodo-fm (3 handlers)
[GIN-debug] PUT    /todos/:id                --> github.com/U-taro-ogw/gin_todo/src/controllers.(*TodoHandler).UpdateTodo-fm (3 handlers)
```


## 準備
`git clone`してgin_todoディレクトリ配下に`.env`ファイルを作成する  
作成したファイルに以下3つを記述する
```
MYSQL_USER=Mysqlのユーザ名
MYSQL_ROOT_PASSWORD=Mysqlのrootパスワード
MYSQL_DATABASE=使用するデータベース名
```


## 動かしてみる

### 起動
```
docker-compose up -d
```

### 新規作成(C)
```
curl -X POST http://localhost:8083/todos \
-H "Content-Type: application/json" \
-d '{"text": "post_todo_1", "status": "1"}'
```

### 一覧取得(R)
```
curl -X GET http://localhost:8083/todos
```

### 1件取得(R)
```
curl -X GET http://localhost:8083/todos/:id
```

### 更新(U)
```
curl -X PUT http://localhost:8083/todos/1 \
-H "Content-Type: application/json" \
-d '{"text": "post_todo_1_update", "status": "2"}'
```

### 削除(D)
```
curl -X DELETE http://localhost:8083/todos/:id
```
