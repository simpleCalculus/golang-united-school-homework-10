package web

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	w.Write([]byte(param))
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func Data(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	data := make([]byte, 0)
	temporarily := make([]byte, 1024)
	for {
		n, err := body.Read(temporarily)
		if err != nil {
			if err == io.EOF {
				data = append(data, temporarily[:n]...)
				break
			}
			w.Write([]byte("Request error: " + err.Error()))
			return
		}
		data = append(data, temporarily[:n]...)
	}
	w.Write([]byte("I got message:\n" + string(data)))

}

func Headers(w http.ResponseWriter, r *http.Request) {
	ha := r.Header.Get("a")
	hb := r.Header.Get("b")

	a, err := strconv.Atoi(ha)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(hb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sum := a + b

	w.Header().Add("a+b", strconv.Itoa(sum))
}
