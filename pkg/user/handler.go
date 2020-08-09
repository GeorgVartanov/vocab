package user

import (
	"encoding/json"
	"github.com/georgvartanov/vocabProject/pkg/user/create"
	"github.com/georgvartanov/vocabProject/pkg/user/login"
	"github.com/georgvartanov/vocabProject/pkg/user/read"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Handler(c create.ServiceCreater, r read.ServiceReader, l login.ServiceLoginer) {
	router := mux.NewRouter()
	router.HandleFunc("/create/", Creating(c)).Methods("POST")
	router.HandleFunc("/readall/", ReadingAll(r)).Methods("GET")
	router.HandleFunc("/read/", Reading(r)).Methods("GET")
	router.HandleFunc("/login/", logining(l)).Methods("POST")

	http.ListenAndServe(":8080", router)

}

func logining(l login.ServiceLoginer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)
		myUser := login.User{}
		if err := decoder.Decode(&myUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := l.GetUserByEmail(myUser.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		if err =	bcrypt.CompareHashAndPassword(user.Password, []byte(myUser.Password)); err!=nil{
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		json.NewEncoder(w).Encode("Password is Good")
	}
}

func Creating(c create.ServiceCreater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var user create.User
		if err := decoder.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.Create(user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		json.NewEncoder(w).Encode("Class was added")
	}
}

func ReadingAll(rd read.ServiceReader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		users, err := rd.ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		json.NewEncoder(w).Encode(users)
	}


}
func Reading(rd read.ServiceReader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)
		myUser := read.User{}
		if err := decoder.Decode(&myUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := rd.Read(myUser.ID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

