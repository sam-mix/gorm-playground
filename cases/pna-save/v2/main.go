package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	db.Table("t_pna_v2").Save(&model.PNA{PnaID: 2})
}
