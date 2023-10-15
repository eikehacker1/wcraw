package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	var initialURL string
	var maxDepth int = 2
	if len(os.Args) > 1 {
		initialURL = os.Args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Informe a URL inicial: ")
		scanner.Scan()
		initialURL = scanner.Text()
	}

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)

		if e.Request.Depth < maxDepth {
			c.Visit(e.Request.AbsoluteURL(link))
		} else {
			fmt.Println("Profundidade máxima atingida, não visitando mais.")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	err := c.Visit(initialURL)
	if err != nil {
		fmt.Println("Erro ao visitar a página inicial:", err)

		file, err := os.Open("urls.txt")
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo urls.txt:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			nextURL := scanner.Text()
			fmt.Println("Próxima URL:", nextURL)
			err := c.Visit(nextURL)
			if err != nil {
				fmt.Println("Erro ao visitar a próxima URL:", err)
				
			}
		}
	}
}

