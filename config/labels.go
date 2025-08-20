// config/labels.go
package config

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func FetchBranchLabels(cfg *Config) ([]BranchLabel, error) {
	if cfg.BranchLabelSource == "" {
		return nil, nil
	}
	cmd := exec.Command("bash", "-lc", cfg.BranchLabelSource)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("branch_label_source 실행 실패: %w", err)
	}
	var labels []BranchLabel
	if err := json.Unmarshal(out, &labels); err != nil {
		return nil, fmt.Errorf("branch_label_source JSON 파싱 실패: %w", err)
	}
	return labels, nil
}
