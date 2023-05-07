package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
    "hash/fnv"
    // "strconv"
)


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


	// それぞれのコンテンツのハッシュ値をkeyとしてマップにまとめる
    contentMap := make(map[uint32][][]string)
    for _, row := range rows {
		title := row[3]
        h := CalculateHash(title)
        contentMap[h] = append(contentMap[h], row)
    }
    // // マップのキーだけを出力する
    // for k := range contentMap {
    //     fmt.Println(k)
    // }


	// // ここからcsvに消すかどうか確認が必要なコンテンツを書き込む
	// confirmationRequiredFile, err := os.Create("confirmation.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w := csv.NewWriter(confirmationRequiredFile)
	// w.WriteAll(confirmation_contents)
}

func CalculateHash(s string) uint32 {
	h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}