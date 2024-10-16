package tools

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Weather struct{}

type IPInfo struct {
	Status string `json:"status"`
	Data   []IP   `json:"data"`
}

type IP struct {
	Location string `json:"location"`
}

type WeatherRes struct {
	City    string        `json:"city"`
	Weather []WeatherInfo `json:"weather"`
}
type WeatherInfo struct {
	Date    string `json:"date"`
	Weather string `json:"weather"`
	Temp    string `json:"temp"`
	W       string `json:"w"`
	Wind    string `json:"wind"`
}

func (w *Weather) GetWeather(city string) {
	res := getWeatherInfo(city)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.Title = res.City + "天气"

	table.Rows = [][]string{
		// 在初始化 table.Rows 时，可以根据上下文自动推断切片的类型，因此可以简化类型声明。
		// []string{"日期", "天气", "风向", "体感温度"} 可以简化为 {"日期", "天气", "风向", "体感温度"}
		{"日期", "天气", "风向", "体感温度"},
	}
	for _, v := range res.Weather {
		table.Rows = append(table.Rows, []string{v.Date, v.Weather, v.Wind, v.Temp})
	}

	table.BorderStyle = ui.NewStyle(ui.ColorWhite)
	table.TitleStyle = ui.NewStyle(ui.ColorGreen)
	table.TextStyle = ui.NewStyle(ui.ColorGreen)
	table.TextAlignment = ui.AlignLeft

	table.SetRect(0, 0, 100, 18)

	ui.Render(table)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		// q 或者 Ctrl+C 退出程序
		case "q", "<C-c>":
			return
		}
	}
}

func getWeatherInfo(city string) (weatherResponse *WeatherRes) {
	api := "https://api.asilu.com/weather/?city=" + city
	resp, err := http.Get(api)
	if err != nil {
		return nil
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	/*
		{
			"city":"深圳",
			"update_time":"07:30",
			"date":"10月16日",
			"weather":[
				{"date":"16日（今天）","weather":"小雨转多云","icon1":"07","icon2":"01","temp":"31/25℃","w":"3-4级转","wind":"东风转无持续风向","iconn":"151","icond":"305"},
				{"date":"17日（明天）","weather":"多云转小雨","icon1":"01","icon2":"07","temp":"31/25℃","w":"3-4级转","wind":"东南风转无持续风向","icond":"101","iconn":"305"},
				{"date":"18日（后天）","weather":"小雨转晴","icon1":"07","icon2":"00","temp":"30/25℃","w":"","wind":"无持续风向","iconn":"150","icond":"305"},
				{"date":"19日（周六）","weather":"阵雨转多云","icon1":"03","icon2":"01","temp":"31/24℃","w":"3-4级","wind":"东风","iconn":"151","icond":"300"},
				{"date":"20日（周日）","weather":"多云转晴","icon1":"01","icon2":"00","temp":"30/23℃","w":"3-4级","wind":"东风","icond":"101","iconn":"150"},
				{"date":"21日（周一）","weather":"晴","icon1":"00","icon2":"00","temp":"31/23℃","w":"3-4级","wind":"东风","icond":"100","iconn":"150"},
				{"date":"22日（周二）","weather":"晴","icon1":"00","icon2":"00","temp":"28/24℃","w":"","wind":"无持续风向","icond":"100","iconn":"150"}
			]
		}
	*/
	out, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(out))
	if err := json.Unmarshal(out, &weatherResponse); err != nil {
		return weatherResponse
	}
	return weatherResponse
}
