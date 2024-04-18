package mxvd

type Config struct {
	RaftNodeId       uint64
	DataDir          string
	Port             string
	JoinNodes        []string
	DoNotJoinCluster bool
	TlsCertFile      string
	TlsKeyFile       string
}

func NewConfig() *Config {
	return &Config{
		Port:             "50003",
		DataDir:          "/personal_data",
		DoNotJoinCluster: false,
	}
}
