package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/ocean-gao/go_ocean_tools/tools"
	"github.com/spf13/cobra"
)

var LogoTips = `
欢迎使用 go_ocean_tools, 请按照下面的指示操作。
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_ocean_tools",
	Short: "开发小工具",
	Long: `基于 go 开发的 小工具集合
   - go_ocean_tools weather 查询天气
   - go_ocean_tools md5 md5 小工具
   - go_ocean_tools base64 base64 小工具
   - go_ocean_tools time  获取当前时间戳
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// 判断是否有 --version 命令选项
		version, _ := cmd.Flags().GetBool("version")
		if version {
			color.Green(tools.GetVersion())
			return
		}

		// 默认执行 help 子命令
		if len(args) == 0 {
			// 实例化一个新的 color 对象，设置前景色为绿色，文字加粗
			colorPrint := color.New(color.Bold)
			colorPrint.Add(color.FgGreen)

			_, _ = colorPrint.Println(LogoTips)

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
	// 注册主命令选项 --versin
	rootCmd.Flags().BoolP("version", "v", false, "当前版本号")

	// 执行命令程序 (内部会自动进行命令参数和选项解析, 进而执行对应的主命令或子命令)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
