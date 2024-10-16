package cmd

import (
	"github.com/ocean-gao/go_ocean_tools/tools"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "字符串 base64 格式",
	Long:  `字符串 base64 格式, 默认加密, 第二个参数 -d 则是解密, 方便快捷使用`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}

		cry := tools.Cryption{}

		decode, _ := cmd.Flags().GetBool("decode")

		if decode {
			cry.DecodeString(args[0])
		} else {
			cry.EncodeToString(args[0])
		}
	},
}

func init() {
	// 注册子命令选项 --decode
	base64Cmd.Flags().BoolP("decode", "d", false, "解密")
	rootCmd.AddCommand(base64Cmd)
}
