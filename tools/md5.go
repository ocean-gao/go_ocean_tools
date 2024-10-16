package tools

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Md5 struct{}
type Options struct {
	Salt     string
	Position int `json:"position"` // 0-左边 1-右边
}

func (m *Md5) Encipher(str string, options Options) {
	saltStr := ""
	if options.Position == 0 {
		saltStr = options.Salt + str
	} else {
		saltStr = str + options.Salt
	}

	data := []byte(str)

	hash := md5.Sum(data)
	// MD5 哈希值 以十六进制的形式输出
	enStr := fmt.Sprintf("%x", hash)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.Title = "MD5加密"

	table.Rows = [][]string{
		{"原始字符", "盐", "待加密", "加密后"},
	}
	table.Rows = append(table.Rows, []string{str, options.Salt, saltStr, enStr})

	table.BorderStyle = ui.NewStyle(ui.ColorWhite)
	table.TextStyle = ui.NewStyle(ui.ColorGreen)
	table.TitleStyle = ui.NewStyle(ui.ColorGreen)

	table.SetRect(0, 0, 80, 5)

	ui.Render(table)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		// q 或者 Ctrl+C 退出程序
		case "q", "<C-c>":
			return
		// c 复制到粘贴板
		case "c":
			_ = clipboard.WriteAll(enStr)
			return
		}
	}
}
