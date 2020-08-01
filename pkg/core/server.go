package core

import (
	"fmt"
	"gostardardk8s/pkg/global"
	"gostardardk8s/initialize"
	"net/http"
	"time"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, s.Addr)
	fmt.Printf(` 
   _     _         _               ______                                     
(_)   | |       | |             (____  \                                _   
 _____| |  ___  | |__   _____    ____)  )  ____  _   _  _____  ____   _| |_ 
|  _   _) / _ \ |  _ \ | ___ |  |  __  (  / ___)| | | |(____ ||  _ \ (_   _)
| |  \ \ | |_| || |_) )| ____|  | |__)  )| |    | |_| |/ ___ || | | |  | |_ 
|_|   \_) \___/ |____/ |_____)  |______/ |_|     \__  |\_____||_| |_|   \__)
                                                (____/
Limits,like fears,are often just an illusion.`)

	global.GVA_LOG.Error(s.ListenAndServe())
}
