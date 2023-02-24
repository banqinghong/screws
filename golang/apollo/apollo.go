package apollo

import (
	"fmt"

	"github.com/shima-park/agollo"
)
func GetConfig() {
	agollo.Init(
		"apollo.xgimi.com",
		"application",
	)

	fmt.Println("初始化Apollo配置成功")

	//Use your apollo key to test
	fmt.Println("timeout:", agollo.Get("server.port"))
}
