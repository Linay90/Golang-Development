package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB




func main(){
	var err error
	db,err=sql.Open("mysql","username:")
	router:=gin.Default()




}
