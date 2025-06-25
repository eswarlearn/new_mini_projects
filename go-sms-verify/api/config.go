package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal("ACCOUNTSID"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envSERVICEID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_SERVICE_ID")
}
