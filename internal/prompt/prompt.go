// internal/prompt/prompt.go
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Select(label string, options []string, defaultIndex int) (string, error) {
	var picked string
	q := &survey.Select{
		Message: label,
		Options: options,
	}
	if defaultIndex >= 0 && defaultIndex < len(options) {
		q.Default = options[defaultIndex]
	}
	if err := survey.AskOne(q, &picked); err != nil {
		return "", err
	}
	return picked, nil
}

func Input(label, def string) (string, error) {
	var v string
	q := &survey.Input{
		Message: label,
		Default: def,
	}
	if err := survey.AskOne(q, &v, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return v, nil
}

func Confirm(label string, def bool) (bool, error) {
	var ok bool
	q := &survey.Confirm{
		Message: label,
		Default: def,
	}
	if err := survey.AskOne(q, &ok); err != nil {
		return false, err
	}
	return ok, nil
}

func FormatLabel(label, desc string) string {
	if desc == "" {
		return label
	}
	return fmt.Sprintf("%s — %s", label, desc)
}
