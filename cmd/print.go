package cmd

import (
	"bytes"
	"fmt"
	static "rac-tester/aa"
)

func printAa(filename string) {
	file, err := static.Aa.Open(filename) // 相対パスなので / を取り除いてファイル名を指定
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// ファイルを読み込んで出力
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	fmt.Print(buf.String())
}
