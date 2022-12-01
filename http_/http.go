package http_

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lohmander/blocks/core"
)

type HTTPTransport struct{}

func parseRequestBodyBlock(r *http.Request) (*core.Block, error) {
	var body []byte
	var err error

	if r.Body != nil {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		defer r.Body.Close()
	}

	block := &core.Block{
		Value: body,
	}

	return block, nil
}

func blockHTTPRepr(block *core.Block) (map[string]interface{}, error) {
	children := []map[string]interface{}{}

	for _, child := range block.Children {
		if childJson, err := blockHTTPRepr(child); err != nil {
			return nil, err
		} else {
			children = append(children, childJson)
		}
	}

	var value interface{}

	if err := json.Unmarshal(block.Value, &value); err != nil {
		value = string(block.Value)
	}

	return map[string]interface{}{
		"id":       string(block.ID),
		"value":    value,
		"children": children,
	}, nil
}

func writeJson(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func writeJsonError(w http.ResponseWriter, err string) {
	writeJson(w, map[string]string{"error": err})
}

func (t *HTTPTransport) Serve(server *core.BlocksServer) error {
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")

			next.ServeHTTP(w, r)

			log.Println(r.Method, r.URL.Path)
		})
	})

	readHandler := func(w http.ResponseWriter, r *http.Request) {
		var id []byte

		if IDString, ok := mux.Vars(r)["id"]; ok {
			id = []byte(IDString)
		}

		block, err := server.Get(id)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		blockValue, err := blockHTTPRepr(block)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		writeJson(w, blockValue)
	}

	createHandler := func(w http.ResponseWriter, r *http.Request) {
		var parentID []byte

		if parentIDString, ok := mux.Vars(r)["parentID"]; ok {
			parentID = []byte(parentIDString)
		}

		block, err := parseRequestBodyBlock(r)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		block, err = server.Create(parentID, block)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		blockValue, err := blockHTTPRepr(block)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		writeJson(w, blockValue)
	}

	updateHandler := func(w http.ResponseWriter, r *http.Request) {
		var id []byte

		if IDString, ok := mux.Vars(r)["id"]; ok {
			id = []byte(IDString)
		}

		block, err := parseRequestBodyBlock(r)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		block, err = server.Update(id, block)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		blockValue, err := blockHTTPRepr(block)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		writeJson(w, blockValue)
	}

	r.Methods(http.MethodGet, http.MethodOptions).Path("/blocks").HandlerFunc(readHandler)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/blocks").HandlerFunc(createHandler)
	r.Methods(http.MethodGet, http.MethodOptions).Path("/blocks/{id}").HandlerFunc(readHandler)
	r.Methods(http.MethodPut, http.MethodOptions).Path("/blocks/{id}").HandlerFunc(updateHandler)
	r.Methods(http.MethodPost, http.MethodOptions).Path("/blocks/{parentID}").HandlerFunc(createHandler)

	log.Printf("HTTP transport listening on port %d", 8090)
	r.Use(mux.CORSMethodMiddleware(r))
	return http.ListenAndServe(":8090", r)
}
