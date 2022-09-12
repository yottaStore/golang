package keyvalue

import (
	"log"
	"net/http"
	"yottadb/dbdriver"
	"yottadb/dbdriver/keyvalue"
)

func Handler(w http.ResponseWriter, req dbdriver.Request, kvd keyvalue.Driver) {

	switch req.Method {

	case "read":
		resp, err := kvd.Read(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("Read request failed")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write([]byte(resp.Data)); err != nil {
			log.Println("ERROR: ", err)
		}

	case "write":
		resp, err := kvd.Write(req)
		if err != nil {
			log.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("Write request failed")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write([]byte(resp.Data)); err != nil {
			log.Println("ERROR: ", err)
		}

	case "delete":
		resp, err := kvd.Delete(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("Delete request failed")); err != nil {
				log.Println("ERROR: ", err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err = w.Write([]byte(resp.Data)); err != nil {
			log.Println("ERROR: ", err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("YottaDB KeyValue Method not found")); err != nil {
			log.Println("ERROR: ", err)
		}
	}

}
