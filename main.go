package main

import (
	"ascii-art/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Проверяем количество аргументов командной строки
	if len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	var filename string
	args := os.Args[1:]
	var str, style string
	for _, v := range args {
		if strings.HasPrefix(v, "--output=") {
			filename = v[9:]
		}
	}
	if filename == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	if len(args) == 3 {
		style = os.Args[3]
	} else if style == "" {
		style = "standard"
	}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	filePath := filepath.Join(currentDir, filename)
	f, e := os.Create(filePath)
	if e != nil {
		fmt.Println("Error creating file:", e)
		os.Exit(1)
	}
	str = os.Args[2]
	//проверка входных данных есть ли они в аски
	cor := utils.CheckInput(str)
	if cor == true {
		//получение массива строк разделяя каждую строку
		words := strings.Split(string(str), "\\n")
		//получение файла c стилями
		styleChrs := utils.GetStyle(style)
		//вывод
		utils.Output(styleChrs, words, f)
	}
}
