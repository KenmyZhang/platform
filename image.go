package main
import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "image/jpeg"
)
func main() {
    imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
    //图片正则
    reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
    name := reg.FindStringSubmatch(imagPath)[0]
    fmt.Print(name)
    //通过http请求获取图片的流文件
    resp, _ := http.Get(imagPath)


    body, _ := ioutil.ReadAll(resp.Body)
    config, err := jpeg.DecodeConfig(bytes.NewReader(body))
	if err !=nil {
		fmt.Println("error",err.Error())
	}
	fmt.Println("config:",config.Width*config.Height)

    out, _ := os.Create(name)
    io.Copy(out, bytes.NewReader(body))
    return
}
/*package main
import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
)
func main() {
    imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
    //图片正则
    reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
    name := reg.FindStringSubmatch(imagPath)[0]
    fmt.Print(name)
    //通过http请求获取图片的流文件
    resp, _ := http.Get(imagPath)
    body, _ := ioutil.ReadAll(resp.Body)
    out, _ := os.Create(name)
    io.Copy(out, bytes.NewReader(body))
    return
}*/
