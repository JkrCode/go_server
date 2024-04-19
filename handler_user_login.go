package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	user := cfg.DB.GetUser(params.Email)
	hashedPWerr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if hashedPWerr==nil{
		respondWithJSON(w, 200, User{
			ID: user.ID,
			Email: user.Email,
		})
		return
	} else {
		respondWithError(w, 401, "Password dont match")
	}
	return
}