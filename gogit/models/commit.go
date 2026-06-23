package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Commit struct {
	ID        string
	Message   string
	Timestamp time.Time
	Files     []string
}

var msg string

func CommitFunc() {

	if len(os.Args) < 5 {
		fmt.Println("enter proper commit")
		return
	} else {
		msg = os.Args[4]
	}
	rendomId := randomID(8)
	files, err := getFilesName()

	if err != nil {
		fmt.Println(err)
	}

	utcTimeStamp := time.Now().UTC()

	cmt := Commit{ID: rendomId, Message: msg, Timestamp: utcTimeStamp, Files: files}
	data, err := json.Marshal(cmt)
	if err != nil {
		fmt.Errorf("commit error %w", err)
	}
	err = os.WriteFile(`.gogit/commits/commit.json`, data, 0644)
	if err != nil {
		fmt.Errorf("commit error %w", err)
	}
	fmt.Printf("%+v\n changes are commited")
	
}