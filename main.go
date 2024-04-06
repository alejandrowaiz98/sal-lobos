package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {

	f, err := excelize.OpenFile("excelqlo.xlsx")

	if err != nil {
		panic(err)
	}

	err = readAndWrite(f)

	if err != nil {
		panic(err)
	}

}

func readAndWrite(f *excelize.File) error {
	rand.Seed(time.Now().UnixNano()) // Semilla para generar números aleatorios

	var rows [][]string

	allRows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for i, row := range allRows {
		if i > 10 {
			rows = append(rows, row)
		}
	}

	for i := 0; i < len(rows); i += 15 {
		endIndex := i + 14
		if endIndex >= len(rows) {
			endIndex = len(rows) - 1
		}
		if i+15 >= len(rows) { // Verificar si hay suficientes filas para formar el siguiente bloque de 15 filas
			break
		}
		start := strings.Split(rows[i][0], ",")
		end := strings.Split(rows[i+15][0], ",")

		if len(start) < 3 || len(end) < 3 {
			continue
		}

		startVal2, _ := strconv.ParseFloat(start[1], 64)
		endVal2, _ := strconv.ParseFloat(end[1], 64)
		startVal3, _ := strconv.ParseFloat(start[2], 64)
		endVal3, _ := strconv.ParseFloat(end[2], 64)

		for j := i; j <= endIndex; j++ {
			vals := strings.Split(rows[j][0], ",")
			if len(vals) < 3 {
				continue
			}
			if j != i && j != endIndex {
				val2 := generateRandomValue(startVal2, endVal2)
				val3 := generateRandomValue(startVal3, endVal3)
				vals[1] = strconv.FormatFloat(val2, 'f', 2, 64)
				vals[2] = strconv.FormatFloat(val3, 'f', 2, 64)
			}
			rows[j][0] = strings.Join(vals, ",")
		}
	}

	newFile := excelize.NewFile()

	for i, value := range rows {
		coord := fmt.Sprintf("A%v", i+1)
		newFile.SetCellValue("Sheet1", coord, value)
	}

	err = newFile.SaveAs("NewFile.xlsx")

	if err != nil {
		return err
	}

	return nil
}

func generateRandomValue(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
