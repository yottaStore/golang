package yottafs

import (
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {

	conf := Config{
		"/tmp/yottafs",
		"kkk",
		"8083",
	}

	err := StartServer(conf)

	if err == nil {
		t.Fatal("Server Should have failed: ", err)
	}

	conf.Driver = "direct"

	go func() {
		err := StartServer(conf)
		if err != nil {
			log.Println("Error starting server: ", err)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	resp, err := http.Get("http://localhost:8083")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Error("Error reaching server: ", err)
	}

	log.Println("fourth")

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error reading server response: ", err)
	}
	if string(buff) != "Hello from yottafs-go v 0.0.1!" {
		t.Error("Response didn't match")
	}

}
