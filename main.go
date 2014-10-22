//用于测试支付宝支付
//修改alipay.go中部分带"xxxx"的参数，然后修改html中paydemo.html中对应的网址
//运行后，在本地直接打开html，或者输入"http:127.0.0.1:3000/html/paydemo.html"可以进行测试
//初学之作，对照go语言半小时速成和golang.org文档写成，水平有限，欢迎指正，联系方式: 杭州火烧云科技有限公司 yycmail@163.com(婚礼纪 老叶)

package main

import (
	//go库包
	"log"
	"net/http"
	//"runtime"
	//"text/template"

	//第三方包
	"github.com/gorilla/mux"
	
	//应用自带包
	"pay"
)

var (
	mainmux		*mux.Router
)
	
func main() {
	var err error
	//runtime.GOMAXPROCS(appCfg.CpuNumber)
	mainmux	=mux.NewRouter()
	http.Handle("/", mainmux)
	SetMainRouter(mainmux)
	err 	=http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
}

func SetMainRouter(mainmux	*mux.Router){
	//静态文件解析，实际运行时可交给nginx处理
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	
	//支付
	mainmux.HandleFunc("/alipay_wap", pay.AlipayHandler_Wap)
	mainmux.HandleFunc("/alipay_web", pay.AlipayHandler_Web)
	mainmux.HandleFunc("/alipay_web_callback", pay.Alipay_Web_CallbackHandler)
	mainmux.HandleFunc("/alipay_web_notify", pay.Alipay_Web_NotifyHandler)	
	mainmux.HandleFunc("/alipay_wap_callback", pay.Alipay_Wap_CallbackHandler)
	mainmux.HandleFunc("/alipay_wap_notify", pay.Alipay_Wap_NotifyHandler)	
	mainmux.HandleFunc("/alipay_refund", pay.Alipay_Refund_Handler)
	mainmux.HandleFunc("/alipay_refund_notify", pay.Alipay_Refund_NotifyHandler)	
	//mainmux.HandleFunc("/alipay_log", pay.GetAlipayTracklog_Handler)
}


