package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codizen",
	Short: "코드 컨벤션을 위한 브랜치/커밋 도우미",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}