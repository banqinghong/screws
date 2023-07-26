package huawei

import (
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	defaultSheet = "Sheet1"
	firstColumn  = 'A'
	Suffix       = ".xlsx"
)

type ExcelContent struct {
	Title   []string
	Content [][]string
	OutFile string
}

func toCharStr(i int) string {
	return string(rune('A' + i))
}

//SaveExcel save result to excel file
func SaveExcel(excel *ExcelContent) {
	log.Println("start save result to excel.")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet(defaultSheet)

	style, err := f.NewStyle(`{"font":{"family":"宋体","size":12},"alignment":{"horizontal":"center","vertical":"center"}, "border":[{"type": "left", "style": 2, "color": "#000000"}, {"type": "right", "style": 2, "color": "#000000"}, {"type": "top", "style": 2, "color": "#000000"}, {"type": "bottom", "style": 2, "color": "#000000"}]}`)
	if err != nil {
		fmt.Println(err)
	}
	titleStyle, err := f.NewStyle(`{"font":{"family":"宋体","size":16},"alignment":{"horizontal":"center","vertical":"center"}, "fill":{"type":"gradient","pattern":1,"color":["#6495ED","#6495ED"],"shading":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	f.MergeCell(defaultSheet, "G2", "G10")
	f.SetCellValue(defaultSheet, "G2", "测试文字")
	f.SetCellStyle(defaultSheet, "G2", "G10", style)

	f.SetColWidth(defaultSheet, "A", "H", 26)

	// 写入title
	for k, v := range excel.Title {
		//fmt.Printf("column: %s\n", column)
		cell := toCharStr(k) + "1"
		f.SetCellValue(defaultSheet, cell, v)
		f.SetCellStyle(defaultSheet, cell, cell, titleStyle)
	}
	f.SetActiveSheet(index)
	for i, columnContent := range excel.Content {
		for k, v := range columnContent {
			cell := toCharStr(k) + strconv.Itoa(i+2)
			f.SetCellValue(defaultSheet, cell, v)
			f.SetCellStyle(defaultSheet, cell, cell, style)
		}
	}

	// Save xlsx file by the given path.
	outFile := "./dist/" + excel.OutFile + Suffix
	if err := f.SaveAs(outFile); err != nil {
		fmt.Println(err)
	}

	log.Printf("save result to excel[%s] finish\n", outFile)
}
