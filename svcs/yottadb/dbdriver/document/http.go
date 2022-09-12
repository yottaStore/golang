package document

import (
	"log"
	"net/http"
	"yottadb/dbdriver"
)

func HttpHandler(w http.ResponseWriter, req dbdriver.Request, dd Driver) {

	switch req.Driver {

	case "document":

		switch req.Method {

		case "create":
			resp, err := dd.Create(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "read":
			resp, err := dd.Read(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "update":
			resp, err := dd.Update(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "delete":
			resp, err := dd.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("YottaDB KeyValue Method not found")); err != nil {
				log.Println("ERROR: ", err)
			}
		}

	case "collection":

		switch req.Method {

		case "create":
			resp, err := dd.CollectionDriver.Create(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "read":
			resp, err := dd.CollectionDriver.Read(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Read request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "update":
			resp, err := dd.CollectionDriver.Update(req)
			if err != nil {
				log.Println("Error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Write request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		case "delete":
			resp, err := dd.CollectionDriver.Delete(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write([]byte("Delete request failed")); err != nil {
					log.Println("ERROR: ", err)
				}
				return
			}

			w.WriteHeader(http.StatusOK)
			if _, err = w.Write(resp.Data); err != nil {
				log.Println("ERROR: ", err)
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("YottaDB KeyValue Method not found")); err != nil {
				log.Println("ERROR: ", err)
			}
		}

	}

}
