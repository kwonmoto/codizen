// cmd/branch.go
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kwonmoto/codizen/config"
	"github.com/kwonmoto/codizen/internal/ai"
	"github.com/kwonmoto/codizen/internal/git"
	"github.com/kwonmoto/codizen/internal/prompt"
	"github.com/kwonmoto/codizen/internal/util"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "브랜치 생성 흐름 실행",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1) 설정 로드
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("설정 로딩 실패: %w", err)
		}
		if len(cfg.BranchTypeValues) == 0 {
			return fmt.Errorf("설정에 branch_type_values가 비어 있습니다")
		}

		// 2) 라벨 후보
		labels, err := config.FetchBranchLabels(cfg) // 없으면 nil
		if err != nil {
			fmt.Fprintln(os.Stderr, "라벨 소스 실패:", err)
		}

		// 3) 타입 선택(추천 기본값 계산)
		typeOptions := make([]string, len(cfg.BranchTypeValues))
		for i, t := range cfg.BranchTypeValues {
			typeOptions[i] = t.BranchTypeLabel
		}

		defaultTypeIdx := -1
		var pickedLabel config.BranchLabel
		var labelText string

		if len(labels) > 0 {
			// 라벨 후보에서 선택
			opts := make([]string, len(labels))
			for i, l := range labels {
				opts[i] = prompt.FormatLabel(l.BranchLabel, l.Description)
			}
			chosen, err := prompt.Select("브랜치 라벨을 선택하세요", opts, 0)
			if err != nil {
				return err
			}
			// 역매핑
			for _, l := range labels {
				if strings.HasPrefix(chosen, l.BranchLabel) {
					pickedLabel = l
					break
				}
			}

			// 타입 추천 시도 (옵션)
			if cfg.TypeSuggestion && (cfg.OpenAIKey != "" || os.Getenv("OPENAI_API_KEY") != "") {
				if t, err := ai.SuggestType(cfg.OpenAIKey, pickedLabel.Description, typeOptions); err == nil {
					for i, v := range typeOptions {
						if v == t {
							defaultTypeIdx = i
							break
						}
					}
				}
			}
		} else {
			// 라벨 직접 입력
			v, err := prompt.Input("브랜치 라벨을 입력하세요 (예: abc-123-login)", "")
			if err != nil {
				return err
			}
			labelText = v
		}

		// 4) 타입 선택
		pickedType, err := prompt.Select("브랜치 타입을 선택하세요", typeOptions, defaultTypeIdx)
		if err != nil {
			return err
		}

		// 5) 최종 라벨 문자열
		if labelText == "" {
			labelText = pickedLabel.BranchLabel
			if labelText == "" {
				// 라벨 후보에서 선택했는데 branchLabel 비어있을 수는 없지만 안전망
				if pickedLabel.Description != "" {
					labelText = util.Slug(pickedLabel.Description)
				}
			}
		}
		labelText = util.Slug(labelText)

		branchName := pickedType + "/" + labelText
		fmt.Println("생성할 브랜치:", branchName)

		ok, err := prompt.Confirm("이 브랜치로 생성할까요?", true)
		if err != nil {
			return err
		}
		if !ok {
			fmt.Println("취소되었습니다.")
			return nil
		}

		// 6) git checkout -b 실행
		if err := git.CheckoutNew(branchName); err != nil {
			return err
		}
		fmt.Println("🎉 브랜치 생성 완료!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
