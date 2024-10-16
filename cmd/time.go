package cmd

import (
	"github.com/ocean-gao/go_ocean_tools/tools"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "获取时间戳",
	Long: `获取当前时间或者指定时间戳, 如果传递时间则转化对应的时间为时间戳:
	1. 第一个参数 时间格式为 yyyy-MM-dd
	2. 第二个参数 时间格式为 HH:mm:ss, 不传则为 00:00:00
`,
	Run: func(cmd *cobra.Command, args []string) {
		timestamp := tools.Timestamp{}
		if len(args) == 1 {
			timestamp.SetFormatTime(args[0] + " 00:00:00")
		} else if len(args) == 2 {
			timestamp.SetFormatTime(args[0] + " " + args[1])
		} else {
			timestamp.GetCurrentTime()
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
