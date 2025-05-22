package commonbox

import (
	"os"
	"os/user"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/sagernet/sing-box/common/humanize"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/experimental/locale"
	"github.com/sagernet/sing-box/log"
)

var (
	SBasePath        string
	SWorkingPath     string
	STempPath        string
	SUserID          int
	SGroupID         int
	STVOS            bool
	SFixAndroidStack bool
)

func init() {
	debug.SetPanicOnFault(true)
}

type SetupOptions struct {
	BasePath        string
	WorkingPath     string
	TempPath        string
	Username        string
	IsTVOS          bool
	FixAndroidStack bool
}

func Setup(options *SetupOptions) error {
	SBasePath = options.BasePath
	SWorkingPath = options.WorkingPath
	STempPath = options.TempPath
	if options.Username != "" {
		sUser, err := user.Lookup(options.Username)
		if err != nil {
			return err
		}
		SUserID, _ = strconv.Atoi(sUser.Uid)
		SGroupID, _ = strconv.Atoi(sUser.Gid)
	} else {
		SUserID = os.Getuid()
		SGroupID = os.Getgid()
	}
	STVOS = options.IsTVOS

	// TODO: remove after fixed
	// https://github.com/golang/go/issues/68760
	SFixAndroidStack = options.FixAndroidStack

	os.MkdirAll(SWorkingPath, 0o777)
	os.MkdirAll(STempPath, 0o777)
	if options.Username != "" {
		os.Chown(SWorkingPath, SUserID, SGroupID)
		os.Chown(STempPath, SUserID, SGroupID)
	}
	return nil
}

func SetLocale(localeId string) {
	locale.Set(localeId)
}

func Version() string {
	return C.Version
}

func FormatBytes(length int64) string {
	return humanize.Bytes(uint64(length))
}

func FormatMemoryBytes(length int64) string {
	return humanize.MemoryBytes(uint64(length))
}

func FormatDuration(duration int64) string {
	return log.FormatDuration(time.Duration(duration) * time.Millisecond)
}

func ProxyDisplayType(proxyType string) string {
	return C.ProxyDisplayName(proxyType)
}
