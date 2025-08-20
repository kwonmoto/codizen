// config/loader.go
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var candidates = []string{
	"branchzen.config.js",
	"branchzen.config.cjs",
	"branchzen.config.mjs",
	"branchzen.config.json",
	// 나중에 TS 지원 재추가 시:
	// "branchzen.config.ts",
}

func Load() (*Config, error) {
	cfgFile := ""
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			cfgFile = c
			break
		}
	}
	if cfgFile == "" {
		return nil, errors.New("설정 파일(branchzen.config.*)을 찾지 못했습니다")
	}

	switch strings.ToLower(filepath.Ext(cfgFile)) {
	case ".js", ".cjs", ".mjs":
		return runAndParse(exec.Command("node", cfgFile))
	case ".json":
		b, err := os.ReadFile(cfgFile)
		if err != nil {
			return nil, err
		}
		var cfg Config
		return &cfg, json.Unmarshal(b, &cfg)
	default:
		return nil, fmt.Errorf("알 수 없는 설정 확장자: %s", cfgFile)
	}
}

func runAndParse(cmd *exec.Cmd) (*Config, error) {
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("설정 파일 실행 실패: %w", err)
	}
	var cfg Config
	if err := json.Unmarshal(out, &cfg); err != nil {
		return nil, fmt.Errorf("설정 JSON 파싱 실패: %w", err)
	}
	return &cfg, nil
}
