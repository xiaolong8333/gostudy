package main
import (
	"fmt"
	"bytes"
    "io"
    "net/http"
    "time"
)
func main(){
	result := Get("http://hq.sinajs.cn/list=sh600297")
	fmt.Println(result)
}

func Get(url string) string {

    // 超时时间：5秒
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    var buffer [512]byte
    result := bytes.NewBuffer(nil)
    for {
        n, err := resp.Body.Read(buffer[0:])
        result.Write(buffer[0:n])
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
    }

    return result.String()
}