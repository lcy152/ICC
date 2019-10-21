/*
 * ICC PROJECT
 *
 * Welcome to use ICC product
 *
 * API version: 1.0.0
 * Contact: 
 * Generated by: lcy
 */

package main

import (
	sw "go_server/service"
	"log"
	"net/http"
	mysql "go_server/mysql"
)

func main() {
	log.Printf("Server started")
	sql := mysql.GetMysql(USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	mlog := InitLog()
	router := sw.NewRouter(sql, mlog)

	log.Fatal(http.ListenAndServe(":8000", router))
}


