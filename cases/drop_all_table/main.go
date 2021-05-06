package main

import v2 "playground/cases/util/v2/conn"

func main() {
	db := v2.Conn()
	db.Exec("SELECT DATABASE();")
}
