package cmd

import (
	"strconv"

	"github.com/ocean-gao/go_ocean_tools/tools"
	"github.com/spf13/cobra"
)

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "字符串加密成 md5",
	Long: `命令行加密 md5
	1. 第一个参数是要加密的字符串
	2. 第二个参数是盐
	3. 第三个参数 0 或者 1, 0-盐在左边 1-盐在右边
`,
	Run: func(cmd *cobra.Command, args []string) {
		md5Obj := tools.Md5{}
		if len(args) > 3 {
			_ = cmd.Help()
			return
		}
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		if len(args) == 1 {
			md5Obj.Encipher(args[0], tools.Options{Salt: "", Position: 0})
			return
		}
		if len(args) == 2 {
			md5Obj.Encipher(args[0], tools.Options{Salt: args[1], Position: 0})
			return
		}
		if len(args) == 3 {
			// 转为 int
			position, _ := strconv.Atoi(args[2])
			set := map[int]string{
				0: "",
				1: "",
			}
			// 判断 position 是否有效
			if _, ok := set[position]; !ok {
				_ = cmd.Help()
				return
			}
			md5Obj.Encipher(args[0], tools.Options{Salt: args[1], Position: position})
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)
}
