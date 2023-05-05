package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main(){
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


	// titleとauthorが同じcontentが存在する行を特定してdelete_contentsに追加
	delete_contents := [][]string{
		rows[0],
	}
	// すでに同じtitle＆authorが存在するか確認
	seen := make(map[string]bool)
	// delete_contents内の重複を避けるための確認
	is_exist_in_delete_list := make(map[string]bool)


	// titleとauthorが完全に一致するコンテンツをdelete_contentsに追加
	for _, row := range rows {
		title := row[3]
		author := row[6]
		key := title + ":" + author
		if seen[key] && !is_exist_in_delete_list[key] {
			delete_contents = append(delete_contents, row)
			is_exist_in_delete_list[key] = true
		} else {
			seen[key] = true
		}
	}


	// ここからcsvにdelete_contentsを書き込む
	newDeleteFile, err := os.Create("delete.csv")
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(newDeleteFile)
	w.WriteAll(delete_contents)


	// // ここからcsvに消すかどうか確認が必要なコンテンツを書き込む
	// confirmationRequiredFile, err := os.Create("delete.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w := csv.NewWriter(confirmationRequiredFile)
	// w.WriteAll(confirmation_content)
}