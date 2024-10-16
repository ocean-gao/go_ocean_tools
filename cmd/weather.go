package cmd

import (
	"github.com/ocean-gao/go_ocean_tools/tools"
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "查询天气",
	Long:  `查询对应城市的天气`,
	Run: func(cmd *cobra.Command, args []string) {
		weather := tools.Weather{}
		if len(args) == 0 {
			weather.GetWeather("")
			return
		} else {
			weather.GetWeather(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
