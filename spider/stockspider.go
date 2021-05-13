package main

import (
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
	"strconv"
	//"golang.org/x/text/encoding/simplifiedchinese"
)


var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/stocks")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {
	c1_spider()

}

func c1_spider() {
	for i := 1; i < 50; i++ { // 常见的 for 循环，支持初始化语句。
		if i==0 {
			c2_spider("index.html")
		} else {
			c2_spider("/f10/p"+strconv.Itoa(i)+".html")
		}
	}		
}
func c2_spider(url string) {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c1.OnHTML("div[class='link_style']", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, item *colly.HTMLElement) {
			//title := item.ChildText("a[class='c-gap-top-xsmall item_3WKCf']")
			href := item.Attr("href")
			c3_spider(href)
		})
	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://zs.stock.pingan.com"+url)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func c3_spider(url string) {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c1.OnHTML("/html/body/div[4]/div[5]/div[1]/div/div[2]/text()[10]", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	c1.OnHTML("div[class='content_text']", func(e *colly.HTMLElement) {
		content := string(e.Text)
		companynamec := strings.Index(content, "公司名称")
		companynamee := strings.Index(content, "英文名称")
		stockcode := strings.Index(content, "A股代码")
		stockname := strings.Index(content, "A股股票简称")
		stockcate := strings.Index(content, "证券类别")
		stockindustry := strings.Index(content, "所属行业")
		listedcate := strings.Index(content, "上市交易所")
		regulator := strings.Index(content, "所属证监会行业")
		fmt.Println(content[stockcate+15 : stockcate+21])
		var newstockcode = ""
		if content[stockcate+15 : stockcate+21] == "深交" {
			newstockcode = "SZ"+content[stockcode+13 : stockname]
		} else {
			newstockcode = "SH"+content[stockcode+13 : stockname]
		}
		fmt.Println(newstockcode)
		data := []string{
			content[companynamec+15 : companynamee],
			content[companynamee+15 : stockcode],
			newstockcode,
			content[stockname+19 : stockcate],
			content[stockcate+15 : stockindustry],
			content[stockindustry+15 : listedcate],
			content[listedcate+18 : regulator],
			content[regulator+14 : len(content)-1]}
		    data = splitstring(data)
			r, err := Db.Exec("replace into stocks(companynamec, companynamee, stockcode,stockname,stockcate,stockindustry,listedcate) values( ?,?,?,?,?,?,?)",data[0],data[1],data[2],data[3],data[4],data[5],data[6])
			if err != nil {
				fmt.Println("exec failed, ", err)
				return
			}
			id, err := r.LastInsertId()
			if err != nil {
				fmt.Println("exec failed, ", err)
				return
			}

			fmt.Println("insert succ:", id)

	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://zs.stock.pingan.com" + url)
	if err != nil {
		fmt.Println(err.Error())
	}
}
	func splitstring(pageinfo []string) []string {
		spilitinfo := pageinfo
		for i := 0; i < 5; i++ {
		spilitinfo[i] = strings.Replace(pageinfo[i], "'''", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "'", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "''", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "’", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "‘", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "“", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "”", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "，", " ", -1)
		spilitinfo[i] = strings.Replace(pageinfo[i], "？", " ", -1)
	}
		return spilitinfo
	}

