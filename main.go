package main

import (
	"fmt"

	"github.com/thedatashed/xlsxreader"
)

type Player struct {
	Name  string
	Email string
	CPF   string
}

func main() {
	var ps []Player
	xl, err := xlsxreader.OpenFile("./excel/players.xlsx")

	fmt.Println(err)

	defer xl.Close()

	for row := range xl.ReadRows("Sheet1") {
		if row.Index > 1 {
			ps = append(ps, Player{
				Name:  row.Cells[0].Value,
				Email: row.Cells[1].Value,
				CPF:   row.Cells[2].Value,
			})
		}
	}
	fmt.Println(ps)
}
