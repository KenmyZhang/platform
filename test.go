package main
import (
    "crypto/sha1"
    "fmt"
    "io"
    "log"
    "net/http"
    "sort"
)
//微信接口处理函数
func checkSigna(w http.ResponseWriter, r *http.Request) {
    var signature, timestamp, nonce, echostr, token string
    token = "token123" //这里应该与你网站配置接口一致
    r.ParseForm()      //解析参数，默认是不会解析的
    rF := r.Form
    //下面代码不是很标准，我自己研究的
    if v, ok := rF["signature"]; ok {
        signature = v[0]
    }
    if v, ok := rF["timestamp"]; ok {
        timestamp = v[0]
    }
    if v, ok := rF["nonce"]; ok {
        nonce = v[0]
    }
    if v, ok := rF["echostr"]; ok {
        echostr = v[0]
    }
    sArr := []string{token, timestamp, nonce}
    sort.Strings(sArr)                 //将token、timestamp、nonce三个参数进行字典序排序
    str := sArr[0] + sArr[1] + sArr[2] //排序后拼接
    //sha1加密
    t := sha1.New()
    io.WriteString(t, str)
    sha1 := fmt.Sprintf("%x", t.Sum(nil))
    //获得加密后的字符串与signature对比,相等标识该请求来源于微信，返回echostr
    if sha1 == signature {
        fmt.Fprintf(w, echostr) //这个写入到w的是输出到客户端的
    } else {
        fmt.Fprintf(w, "拒绝访问")
    }
}
func main() {
    http.HandleFunc("/signup/wechat/complete", checkSigna) //weixin为我们的访问路径，执行checkSignature
    err := http.ListenAndServe(":8065", nil)  //微信目前只接受80端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
