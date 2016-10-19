package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	pwd, _ := os.Getwd()
	urls, _ := getCSVDATA(pwd + "/csv/PRODUCT_DATA_nmax.csv")
	fmt.Print("total data ")
	fmt.Println(len(urls))
	for index := range urls {

		url := urls[index+1]
		fmt.Print(index)
		fmt.Println(" - mendownload " + url)
		response, e := http.Get(url)

		if e != nil {
			log.Fatal(e)
			fmt.Println("gagal download " + url)
		}

		defer response.Body.Close()

		s := strings.Split(url, "/")

		file, err := os.Create(pwd + "/img/" + s[9])
		if err != nil {
			log.Fatal(err)
		}
		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		fmt.Println("Success!")
	}
}

func getCSVDATA(file string) ([101]string, error) {

	f, _ := os.Open(file)
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()

	var urlImage [101]string
	for i := range result {
		urlImage[i] = result[i][1]
	}
	return urlImage, nil
}
