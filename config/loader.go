package config

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type BranchType struct {
	BranchTypeLabel string `json:"branchTypeLabel"`
	Description     string `json:"description"`
}

type Config struct {
	BranchTypeValues []BranchType `json:"branch_type_values"`
}

func LoadConfig() (*Config, error) {
	// tsx 가 설치되어 있어야 함
	out, err := exec.Command("tsx", "branchzen.config.ts").Output()
	if err != nil {
		return nil, errors.New("설정 파일 실행 실패: " + err.Error())
	}

	var cfg Config
	err = json.Unmarshal(out, &cfg)
	if err != nil {
		return nil, errors.New("설정 JSON 파싱 실패: " + err.Error())
	}

	return &cfg, nil
}