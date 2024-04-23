package main

import (
	"context"
	"log"
	"os"

	"github.com/sjy-dv/mind-x/processor/mxvd"
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
