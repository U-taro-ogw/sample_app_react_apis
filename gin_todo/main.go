package main

import (
	"github.com/U-taro-ogw/gin_todo/src/db"
	"github.com/U-taro-ogw/gin_todo/src/router"
)

func main() {
	dbCon := db.Init()
	router.Router(dbCon)
}
