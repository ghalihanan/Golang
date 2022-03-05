package siswa

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	NIM   int
	Prodi string
}

var data = []student{
	{"K001", "Eshan", 1092, "TI"},
}

func mahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
