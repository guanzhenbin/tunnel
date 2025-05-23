//go:build darwin || linux

package libbox

import (
	"os"
	"runtime"

	cb "github.com/sagernet/sing-box/experimental/commonbox"
	"golang.org/x/sys/unix"
)

var stderrFile *os.File

func RedirectStderr(path string) error {
	if stats, err := os.Stat(path); err == nil && stats.Size() > 0 {
		_ = os.Rename(path, path+".old")
	}
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	if runtime.GOOS != "android" {
		err = outputFile.Chown(cb.SUserID, cb.SGroupID)
		if err != nil {
			outputFile.Close()
			os.Remove(outputFile.Name())
			return err
		}
	}
	err = unix.Dup2(int(outputFile.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		outputFile.Close()
		os.Remove(outputFile.Name())
		return err
	}
	stderrFile = outputFile
	return nil
}
