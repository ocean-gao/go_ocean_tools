package main

import (
	"github.com/fatih/color"
	"github.com/ocean-gao/go_ocean_tools/cmd"
)

var LogoTips = `
欢迎使用 go_ocean_tools, 请按照下面的指示操作。
`

func main() {
	// 实例化一个新的 color 对象，设置前景色为绿色，文字加粗
	colorPrint := color.New(color.Bold)
	colorPrint.Add(color.FgGreen)

	_, _ = colorPrint.Println(LogoTips)

	cmd.Execute()
}
