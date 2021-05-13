package main

import (
	"fmt"
	"github.com/gocolly/colly"
)


func main() {
/*	var zm = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 输出 字符串数组 中的 字符串
	for _, s := range zm {
		fmt.Printf("Unicode: %c  \n", s, s)
	}*/
	c1_spider()
}
func c1_spider(text) {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c1.OnHTML("div[class='search-box']", func(e *colly.HTMLElement,text string) {///html/body/div[1]/div[2]/div/div[2]/div
		fmt.Println(text.Text)
/*		e.ForEach("div[class='letter-panel search-tabs-panel']", func(i int, item *colly.HTMLElement) {
			title :=item.ChildText("ul[class='clearfix']")
			fmt.Println(title)
		})*/
	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://www.jiaoyimao.com/")
	if err != nil {
		fmt.Println(err.Error())
	}
}




