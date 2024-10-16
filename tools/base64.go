package tools

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Cryption struct{}

func splitStringByLength(s string, length int) []string {
	var result []string
	for i := 0; i < len(s); i += length {
		if i+length > len(s) {
			result = append(result, s[i:])
		} else {
			result = append(result, s[i:i+length])
		}
	}
	return result
}

func (c *Cryption) EncodeToString(str string) {
	data := []byte(str)
	sEnc := base64.StdEncoding.EncodeToString(data)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sEncList := splitStringByLength(sEnc, 100)

	p := widgets.NewParagraph()

	p.Title = "加密结果"
	p.Text = strings.Join(sEncList, "\n")

	p.TextStyle.Fg = ui.ColorGreen
	p.TextStyle.Modifier = ui.ModifierBold
	p.BorderStyle.Fg = ui.ColorGreen

	p.WrapText = true

	p.SetRect(0, 0, 105, len(sEncList)+2)

	ui.Render(p)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		// q 或者 Ctrl+C 退出程序
		case "q", "<C-c>":
			return
			// c 复制到粘贴板
		case "c":
			_ = clipboard.WriteAll(sEnc)
			return
		}
	}
}

func (c *Cryption) DecodeString(sEnc string) {
	sDec, err := base64.StdEncoding.DecodeString(sEnc)

	// sDecStr := fmt.Sprintf("%s", sDec) 可以简洁写为 sDecStr := string(sDec)
	sDecStr := string(sDec)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	if err != nil {
		p := widgets.NewParagraph()

		p.Title = "解密结果"
		p.Text = "解密失败:不是一个正确的base64编码"

		p.TextStyle.Fg = ui.ColorGreen
		p.BorderStyle.Fg = ui.ColorRed

		p.SetRect(0, 0, 40, 3)
		ui.Render(p)

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	} else {
		p := widgets.NewParagraph()

		p.Title = "解密结果"
		p.Text = "解密成功:" + sDecStr

		p.TextStyle.Fg = ui.ColorGreen
		p.BorderStyle.Fg = ui.ColorGreen

		p.SetRect(0, 0, 30+len(sDecStr), 3)
		ui.Render(p)

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			case "c":
				_ = clipboard.WriteAll(sDecStr)
				return
			}
		}
	}
}
