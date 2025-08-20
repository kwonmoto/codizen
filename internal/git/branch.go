// internal/git/branch.go
package git

import (
	"fmt"
	"os/exec"
)

func CheckoutNew(name string) error {
	// git checkout -b <name>
	cmd := exec.Command("git", "checkout", "-b", name)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git 실패: %v\n%s", err, string(out))
	}
	return nil
}
