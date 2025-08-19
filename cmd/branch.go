package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kwonmoto/codizen/config"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "브랜치를 생성합니다",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("❌ 설정 파일 로딩 실패:", err)
			os.Exit(1)
		}

		fmt.Println("🧱 브랜치 타입 목록:")
		for i, t := range cfg.BranchTypeValues {
			fmt.Printf("[%d] %s - %s\n", i+1, t.BranchTypeLabel, t.Description)
		}

		fmt.Print("선택할 번호를 입력하세요: ")
		var idx int
		fmt.Scanln(&idx)

		if idx < 1 || idx > len(cfg.BranchTypeValues) {
			fmt.Println("❌ 유효하지 않은 번호입니다.")
			os.Exit(1)
		}

		branchType := cfg.BranchTypeValues[idx-1].BranchTypeLabel

		fmt.Print("브랜치 이름을 입력하세요: ")
		var branchLabel string
		fmt.Scanln(&branchLabel)

		branchName := fmt.Sprintf("%s/%s", branchType, branchLabel)

		err = exec.Command("git", "checkout", "-b", branchName).Run()
		if err != nil {
			fmt.Println("❌ Git 브랜치 생성 실패:", err)
			os.Exit(1)
		}

		fmt.Println("🎉 브랜치 생성 완료:", branchName)
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
