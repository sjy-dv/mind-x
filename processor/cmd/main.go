package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sjy-dv/mind-x/processor/mxvd"
	"github.com/sjy-dv/mind-x/processor/transformers"
	"github.com/sjy-dv/mind-x/processor/transport"
)

func main() {

	dbID := ""
	vdb, err := mxvd.InitMXVD()
	if err != nil {
		log.Fatal(err)
	}
	filePath := "personal.mx"
	data, err := os.ReadFile(filePath)
	if string(data) == "" || err != nil {
		newID := bootPersonal(vdb)
		if newID == "" {
			log.Fatal("personalDB Create Error")
		}
		dbID = newID
	} else {
		dbID = string(data)
	}
	status, err := vdb.DatabaseStatus(context.Background(), []byte(dbID))
	if err != nil {
		log.Fatal("LoadDB Fatal Error ", err)
	}
	log.Println(status.Id, status.Size, status.Dimension)
	//load llm
	llama3, err := transformers.LoadLLAMA3()
	if err != nil {
		log.Fatal("LoadLLM Fatal Error ", err)
	}
	go transport.NewSocketServer().Run(llama3, vdb)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signals
	log.Printf("Received signal: %s. Shutting down...\n", sig)
}

func bootPersonal(vdb *mxvd.MXVDProxy) string {
	ids, err := vdb.CreateDatabase(context.Background(), 384, 1, 3)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	err = os.WriteFile("personal.mx", ids, 0400)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(ids)
}
