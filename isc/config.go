package isc

import "fmt"

type RelayConfigurations struct {
	TlsCert             string
	TlsKey              string
	EnableLocalTls      bool
	LocalServerPort     int
	EnableRemoteTls     bool
	RemoteServerIp      string
	RemoteServerPort    int
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type ClientConfigurations struct {
	EnableRemoteTls     bool
	RemoteServerIp      string
	RemoteServerPort    int
	SourceId            string
	EventLogUrl         string
	DiagInterval        int
}

type ServerConfigurations struct {
	TlsCert             string
	TlsKey              string
	EnableLocalTls      bool
	LocalServerPort     int
	SourceId            string
	EventLogUrl         string
}

func (v RelayConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("TlsCert: %s", v.TlsCert))
	strings = append(strings, fmt.Sprintf("TlsKey: %s", v.TlsKey))
	strings = append(strings, fmt.Sprintf("EnableLocalTls: %t", v.EnableLocalTls))
	strings = append(strings, fmt.Sprintf("LocalServerPort: %d", v.LocalServerPort))
	strings = append(strings, fmt.Sprintf("EnableRemoteTls: %t", v.EnableRemoteTls))
	strings = append(strings, fmt.Sprintf("RemoteServerIp: %s", v.RemoteServerIp))
	strings = append(strings, fmt.Sprintf("RemoteServerPort: %d", v.RemoteServerPort))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v ClientConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("EnableRemoteTls: %t", v.EnableRemoteTls))
	strings = append(strings, fmt.Sprintf("RemoteServerIp: %s", v.RemoteServerIp))
	strings = append(strings, fmt.Sprintf("RemoteServerPort: %d", v.RemoteServerPort))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))
	strings = append(strings, fmt.Sprintf("DiagInterval: %d", v.DiagInterval))

	return strings
}

func (v ServerConfigurations) ToString() []string {
	strings := []string{}
	strings = append(strings, fmt.Sprintf("TlsCert: %s", v.TlsCert))
	strings = append(strings, fmt.Sprintf("TlsKey: %s", v.TlsKey))
	strings = append(strings, fmt.Sprintf("EnableLocalTls: %t", v.EnableLocalTls))
	strings = append(strings, fmt.Sprintf("LocalServerPort: %d", v.LocalServerPort))
	strings = append(strings, fmt.Sprintf("SourceId: %s", v.SourceId))
	strings = append(strings, fmt.Sprintf("EventLogUrl: %s", v.EventLogUrl))

	return strings
}
