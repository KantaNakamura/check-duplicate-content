package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)


func main() {
	// csvを読み込む
	file, err := os.Open("content.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 読み込んだcsvのデータをrowsに追加
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	delete_contents := searchDeleteContent(rows)
	// confirmation_contents := 

	// ここからcsvにdelete_contentsを書き込む
	newDeleteFile, err := os.Create("delete.csv")
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(newDeleteFile)
	w.WriteAll(delete_contents)
	fmt.Println(len(rows))
	fmt.Println(len(delete_contents))
}


// 消すべきコンテンツ（titleとauthorが完全に一致するコンテンツ）を探す
func searchDeleteContent(rows [][]string) [][]string {	// titleとauthorが同じcontentが存在する行を特定してdelete_contentsに追加
	delete_contents := [][]string{
		rows[0],
	}
	// すでに同じtitle＆authorが存在するか確認
	seen := make(map[string]bool)

	// titleとauthorが完全に一致するコンテンツをdelete_contentsに追加
	for _, row := range rows {
		title := row[3]
		author := row[6]
		key := title + ":" + author
		if seen[key] {
			delete_contents = append(delete_contents, row)
		} else {
			seen[key] = true
		}
	}

	return delete_contents
}
