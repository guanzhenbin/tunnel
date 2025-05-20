package rocketbox

import (
	"context"

	box "github.com/sagernet/sing-box"
	"github.com/sagernet/sing-box/include"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/json"
)

func parseConfig(ctx context.Context, configContent string) (option.Options, error) {
	options, err := json.UnmarshalExtendedContext[option.Options](ctx, []byte(configContent))
	if err != nil {
		return option.Options{}, E.Cause(err, "decode config")
	}
	return options, nil
}

func BaseContext(platformInterface PlatformInterface) context.Context {
	return box.Context(
		context.Background(),
		include.InboundRegistry(),
		include.OutboundRegistry(),
		include.EndpointRegistry(),
		include.DNSTransportRegistry(),
	)
}

type PlatformInterface interface {
	WriteLog(message string)
	UsePlatformAutoDetectInterfaceControl() bool
	AutoDetectInterfaceControl(fd int32) error
	OpenTun(options TunOptions) (int32, error)
	UseProcFS() bool
	FindConnectionOwner(ipProtocol int32, sourceAddress string, sourcePort int32, destinationAddress string, destinationPort int32) (int32, error)
	PackageNameByUid(uid int32) (string, error)
	UIDByPackageName(packageName string) (int32, error)
	StartDefaultInterfaceMonitor(listener InterfaceUpdateListener) error
	CloseDefaultInterfaceMonitor(listener InterfaceUpdateListener) error
	GetInterfaces() (NetworkInterfaceIterator, error)
	UnderNetworkExtension() bool
	IncludeAllNetworks() bool
	ReadWIFIState() *WIFIState
	SystemCertificates() StringIterator
	ClearDNSCache()
	SendNotification(notification *Notification) error
}

type TunOptions interface {
	GetMTU() int32
	GetAutoRoute() bool
	GetStrictRoute() bool
}

type InterfaceUpdateListener interface {
	UpdateDefaultInterface(interfaceName string, interfaceIndex int32, isExpensive bool, isConstrained bool)
}

type WIFIState struct {
	SSID  string
	BSSID string
}

type NetworkInterfaceIterator interface {
	Next() *NetworkInterface
	HasNext() bool
}

type NetworkInterface struct {
	Index     int32
	MTU       int32
	Name      string
	Addresses StringIterator
	Flags     int32

	Type      int32
	DNSServer StringIterator
	Metered   bool
}

type Notification struct {
	Identifier string
	TypeName   string
	TypeID     int32
	Title      string
	Subtitle   string
	Body       string
	OpenURL    string
}
