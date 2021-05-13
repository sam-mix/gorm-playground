package main

import (
	"playground/cases/model"
	v1 "playground/cases/util/v1/conn"
)

func main() {
	db := v1.Conn()

	db.Table("t_pna_v1").Save(&model.PNAv1{PnaID: 1})

}
