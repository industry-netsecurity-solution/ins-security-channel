package config

type RelayConfigurations struct {
	TlsCert string
	TlsKey string
	LocalTlsServerPort  int
	RemoteTlsServerIp   string
	RemoteTlsServerPort int
}

type ClientConfigurations struct {
	RemoteTlsServerIp   string
	RemoteTlsServerPort int
}

type ServerConfigurations struct {
	TlsCert string
	TlsKey string
	LocalTlsServerPort  int
}
