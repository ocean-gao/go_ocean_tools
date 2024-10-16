package tools

import (
	"log"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func _formatTime(timestampMillis int64) string {
	// 将 Unix 时间戳转换为以秒和纳秒为单位
	seconds := timestampMillis / 1000
	nanos := (timestampMillis % 1000) * 1e6

	// 将 Unix 时间戳转换为 time.Time 类型
	return time.Unix(seconds, nanos).Format("2006-01-02 15:04:05")
}

type Timestamp struct{}

func (t *Timestamp) GetCurrentTime() {
	// 以毫秒为单位的时间戳
	timestampMillis := time.Now().UnixMilli()
	timeUnix := strconv.FormatInt(timestampMillis, 10)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()

	l.Title = "时间信息"

	l.Rows = []string{
		"时间:" + _formatTime(timestampMillis),
		"时间戳:" + timeUnix,
	}

	l.TextStyle = ui.NewStyle(ui.ColorGreen)
	l.TitleStyle = ui.NewStyle(ui.ColorGreen)

	l.WrapText = true

	l.SetRect(0, 0, 40, 8)
	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "c":
			_ = clipboard.WriteAll(timeUnix)
			return
		}
	}
}

func (t *Timestamp) SetFormatTime(str string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	formatTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		p := widgets.NewParagraph()

		p.Title = "时间戳"
		p.Text = "时间格式错误"

		p.TextStyle.Fg = ui.ColorRed
		p.BorderStyle.Fg = ui.ColorRed

		p.SetRect(0, 0, 30, 3)

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
		// 以毫秒为单位的时间戳
		timestampMillis := formatTime.UnixMilli()
		timeUnix := strconv.FormatInt(timestampMillis, 10)

		l := widgets.NewList()

		l.Title = "时间信息"

		l.Rows = []string{
			"时间:" + _formatTime(timestampMillis),
			"时间戳:" + timeUnix,
		}

		l.TextStyle = ui.NewStyle(ui.ColorGreen)
		l.TitleStyle = ui.NewStyle(ui.ColorGreen)

		l.WrapText = true

		l.SetRect(0, 0, 40, 8)
		ui.Render(l)

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			case "c":
				_ = clipboard.WriteAll(timeUnix)
				return
			}
		}
	}
}
