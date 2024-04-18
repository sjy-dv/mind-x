package main

import (
	"flag"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sjy-dv/mind-x/mindx-v/mxvd"
	"github.com/sjy-dv/mind-x/mindx-v/operating"
)

func main() {
	var joinNodesRaw string

	config := mxvd.NewConfig()
	flag.Uint64Var(&config.RaftNodeId, "node-id", 0, "Raft node ID")
	flag.StringVar(&config.Port, "port", operating.GetenvDefault("MXVD_PORT", "6000"), "Node port")
	flag.StringVar(&joinNodesRaw, "join", operating.GetenvDefault("MXVD_JOIN", ""), "Comma separated list of existing cluster nodes")
	flag.StringVar(&config.DataDir, "data-dir", operating.GetenvDefault("MXVD_DATA_DIR", "/tmp"), "Data directory")
	flag.StringVar(&config.TlsCertFile, "cert", operating.GetenvDefault("MXVD_CERT", ""), "TLS Certificate file")
	flag.StringVar(&config.TlsKeyFile, "key", operating.GetenvDefault("MXVD_KEY", ""), "TLS Key file")
	flag.Parse()

	if joinNodesRaw == "false" {
		config.DoNotJoinCluster = true
	} else {
		config.JoinNodes = splitCommaSeparatedValues(joinNodesRaw)
	}

	if _, err := os.Stat(config.DataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(config.DataDir, os.ModePerm); err != nil {
			log.Fatal("Failed to create data directory.")
		}
	}

	server := mxvd.NewServer(config)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	defer server.Stop()

	if !config.DoNotJoinCluster {
		if err := server.JoinCluster(); err != nil {
			log.Fatal(err)
		}
	}

	// Shutdown
	<-operating.InterruptSignal()
	log.Info("Shutdown")
}

func splitCommaSeparatedValues(raw string) []string {
	result := make([]string, 0)
	for _, s := range strings.Split(raw, ",") {
		if len(s) > 0 {
			result = append(result, strings.Trim(s, " "))
		}
	}
	return result
}
