// config/types.go
package config

type BranchType struct {
	BranchTypeLabel string `json:"branchTypeLabel"`
	Description     string `json:"description"`
}

type BranchLabel struct {
	BranchLabel string `json:"branchLabel"`
	Description string `json:"description"`
}

type Config struct {
	OpenAIKey         string       `json:"open_ai_key"`
	TypeSuggestion    bool         `json:"type_suggestion"`
	BranchTypeValues  []BranchType `json:"branch_type_values"`
	BranchLabelSource string       `json:"branch_label_source"` // e.g. ./scripts/fetch-jira-labels.sh
}
