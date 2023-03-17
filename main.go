package main

import (
	"fmt"
	"math/rand"

	"github.com/thedatashed/xlsxreader"
	"github.com/xuri/excelize/v2"
)

type Player struct {
	Name  string
	Email string
	CPF   string
}

func main() {
	leitura()
	criacao()
}

func leitura() {
	var ps []Player
	xl, err := xlsxreader.OpenFile("./excel/players.xlsx")

	fmt.Println(err)

	defer xl.Close()
	fmt.Println(xl)

	for row := range xl.ReadRows("Sheet1") {

		if row.Index > 1 {
			ps = append(ps, Player{
				Name:  row.Cells[0].Value,
				Email: row.Cells[1].Value,
				CPF:   row.Cells[2].Value,
			})
		}
	}
}

func criacao() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("./novoExcel/" + randSeq(10) + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
