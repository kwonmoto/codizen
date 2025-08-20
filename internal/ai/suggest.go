// internal/ai/suggest.go
package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
)

func SuggestType(openAIKey string, description string, options []string) (string, error) {
	key := openAIKey
	if key == "" {
		key = os.Getenv("OPENAI_API_KEY")
	}
	if key == "" {
		return "", errors.New("OpenAI 키가 없습니다")
	}

	prompt := "다음 작업 설명에 가장 알맞은 브랜치 타입을 딱 하나만 고르세요.\n" +
		"옵션: " + strings.Join(options, ", ") + "\n" +
		"설명: " + description + "\n" +
		"정답만 소문자로 반환."

	body := map[string]any{
		"model": "gpt-4o-mini", // 가볍게
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"temperature": 0,
	}

	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	if len(res.Choices) == 0 {
		return "", errors.New("추천 결과 없음")
	}
	ans := strings.TrimSpace(res.Choices[0].Message.Content)
	// 답이 옵션 중 하나인지 확인
	for _, opt := range options {
		if strings.EqualFold(ans, opt) {
			return opt, nil
		}
	}
	return "", errors.New("추천이 옵션과 일치하지 않습니다: " + ans)
}
