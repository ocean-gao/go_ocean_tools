package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/ocean-gao/go_ocean_tools/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_ocean_tools",
	Short: "开发小工具",
	Long: `基于 go 开发的 小工具集合
   - go_ocean_tools weather 查询天气
   - go_ocean_tools mobile 手机归属地查询
   - go_ocean_tools md5 md5 小工具
   - go_ocean_tools base64 base64 小工具
   - go_ocean_tools time  获取当前时间戳
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			color.Green("当前版本号: " + utils.GetVersion())
			return
		}

		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().BoolP("version", "v", false, "当前版本号")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
