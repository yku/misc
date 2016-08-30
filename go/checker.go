package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

func getDetailTitle(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	title := ""
	doc.Find("#detail_title > p").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})

	date := ""
	doc.Find("#detail_title > div > span.newsday").Each(func(i int, s *goquery.Selection) {
		date = s.Text()
	})
	return "[" + date + "] " + title
}

func readFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func writeFile(path string, data []string) {
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for i := range data {
		writer.WriteString(data[i] + "\n")
	}

	writer.Flush()
}

func main() {
	in_file := flag.String("i", "./url_list.txt", "input file.")
	out_file := flag.String("o", "./url_list.txt", "output file.")
	start := flag.Int("s", 27150, "start url")
	count := flag.Int("c", 100, "count")
	flag.Parse()

	// read datafile and create map. key => url, value => title
	data := readFile(*in_file)
	m := map[string]string{}
	for i := range data {
		slice := strings.Split(data[i], ",")
		m[slice[0]] = slice[1]
	}

	end := *start + *count
	for idx := *start; idx <= end; idx++ {
		target := fmt.Sprintf("http://sen-no-kaizoku.sega-net.com/info/detail/info_detail%06d.html", idx)
		_, exists := m[target]
		if exists {
			continue
		}
		resp, err := http.Get(target)
		if err == nil && resp.StatusCode == 200 {
			title := getDetailTitle(target)
			m[target] = title
			s := fmt.Sprintf("%s,%s", target, title)
			data = append(data, s)
			fmt.Println(title)
			fmt.Println(target)
			resp.Body.Close()
		}
	}
	writeFile(*out_file, data)
}
