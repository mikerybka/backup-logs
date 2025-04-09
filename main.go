package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/mikerybka/util"
)

func main() {
	logDir := filepath.Join(util.HomeDir(), "data/log/requests")
	backupPath := filepath.Join(util.HomeDir(), "data/logs/https", time.Now().Format("2006-01-02")+".tar.gz")

	if util.Exists(backupPath) {
		fmt.Println("backup already performed today")
		return
	}

	// mkdir -p
	err := os.MkdirAll(filepath.Dir(backupPath), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	// tar -czf
	cmd := exec.Command("tar", "-czf", backupPath, "-C", filepath.Dir(logDir), filepath.Base(logDir))
	cmd.Dir = logDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(strings.TrimSpace(string(out)))
		os.Exit(3)
	}

	// rm -rf
	err = os.RemoveAll(logDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}
