package isc

type RelayConfigurations struct {
	TlsCert             string
	TlsKey              string
	LocalTlsServerPort  int
	RemoteTlsServerIp   string
	RemoteTlsServerPort int
	SourceId            string
	EventLogUrl         string
}

type ClientConfigurations struct {
	RemoteTlsServerIp   string
	RemoteTlsServerPort int
	EventLogUrl         string
}

type ServerConfigurations struct {
	TlsCert            string
	TlsKey             string
	LocalTlsServerPort int
	SourceId           string
	EventLogUrl        string
}
