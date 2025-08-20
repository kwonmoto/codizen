package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codizen",
	Short: "codizen: branch/commit 컨벤션을 통합하는 CLI",
	Long: `codizen은 branchzen + commitzen 기능을 통합한 도구입니다.
사용 예:
  codizen branch
  codizen commit
  codizen --help`,
	// 아무 서브커맨드/플래그 없이 실행하면 사용법을 보여주도록
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help() // help 출력
	},
	SilenceUsage:  true, // 에러 시 사용법 중복 출력 방지
	SilenceErrors: true, // 에러는 우리가 제어
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}