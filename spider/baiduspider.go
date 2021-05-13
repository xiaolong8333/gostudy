package main

import (
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
    "strings"
	//"time"
	"os"
)

var Db *sqlx.DB
var contentfile []string

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/baidu")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {
	content := Ioutil("file.txt")
	str_arr := strings.Split(content, `，`)

	// 输出 字符串数组 中的 字符串
	for _, str := range str_arr {
		 c1_spider(str)
	}
}

func c1_spider(keywords string) {
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1))
	c1.OnHTML("div[class='c-font-medium list_1V4Yg']", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, item *colly.HTMLElement) {
			title := item.Text
			// Missing Example
			_, found := Find(contentfile, title)
			if !found {
				contentfile = append(contentfile,title) // 追加1个元素
				appendToFile("sports.txt",title+"\n")
				fmt.Println("Value not found in slice")
				c1_spider(title)
			}
		})
	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c1.Visit("https://www.baidu.com/s?wd="+keywords)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func Ioutil(name string) string{
	if contents,err := ioutil.ReadFile(name);err == nil {
	//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
	result := strings.Replace(string(contents),"\n","",1)
		return result
	}
	return ""
}
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
	// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}



