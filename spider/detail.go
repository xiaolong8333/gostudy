package main

import (
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
type member struct {
	stockcode string `db :"stockcode"`

}

var stockdetail  []string


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
	rows ,err :=Db.Query("select stockcode from stocks")
	check(err)
	//遍历查询的结果集合
	for rows.Next(){
		stockdetail := stockdetail[:]
		var s member
		//将从数据库中查询到的值对应到结构体中相应的变量中
		err = rows.Scan(&s.stockcode)
		check(err)
			fmt.Println(s.stockcode)
			stockcode := s.stockcode
			// 输出 字符串数组 中的 字符串
			stockdetail = append(stockdetail, stockcode)
			c1_spider(stockcode,stockdetail)
	}
}

func c1_spider(stockcode string,stockdetail []string) {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c1.OnHTML("table[class='quote-info'] > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, item *colly.HTMLElement) {
			item.ForEach("td", func(y int, item *colly.HTMLElement) {
				str :=item.ChildText("span")
				stockdetail = append(stockdetail, str)
			})
		})
		if len(stockdetail)<26 {
			return
		}

		_, err := Db.Exec("replace into stocksdetail(stockcode,v1,v2,v3,v4,v5,v6,v7,v8,v9,v10,v11,v12,v13,v14,v15,v16,v17,v18,v19,v20,v21,v22,v23,v24,v25,v26,v27) " +
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", stockdetail[0],stockdetail[1],stockdetail[2],stockdetail[3],stockdetail[4],stockdetail[5],stockdetail[6],stockdetail[7],stockdetail[8],stockdetail[9],stockdetail[10],stockdetail[11],stockdetail[12],stockdetail[13],stockdetail[14],stockdetail[15],stockdetail[16],stockdetail[17],stockdetail[18],stockdetail[19],stockdetail[20],stockdetail[21],stockdetail[22],stockdetail[23],stockdetail[24],stockdetail[25],stockdetail[26],stockdetail[27])
		if err != nil {
			fmt.Println("exec failed, ", err)
		}
		fmt.Println("asdsa")
	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://xueqiu.com/S/"+stockcode)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func check (err error){
	if err != nil{
		fmt.Println(err)
	}
}




