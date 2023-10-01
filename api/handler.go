package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altinthaqi/jot-fusion/helper"
	"github.com/altinthaqi/jot-fusion/model"
	"github.com/gookit/validate"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		accounts, err := s.store.GetAccounts()
		if err != nil {
			return nil
		}

		return WriteJSON(w, http.StatusOK, accounts)
	}
	if r.Method == "POST" {
		createAccountReq := new(model.CreateAccountRequest)
		if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
			return err
		}

		v := validate.Struct(createAccountReq)
		if !v.Validate() {
			return fmt.Errorf("missing required fields")
		}

		account := model.NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

		if err := s.store.CreateAccount(account); err != nil {
			return err
		}

		_, err := createJWT(account)
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, account)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helper.GetQueryInt(r, "id")
	if err != nil {
		return nil
	}

	if r.Method == "GET" {
		account, err := s.store.GetAccountByID(id)
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		err := s.store.DeleteAccount(id)

		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(model.TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferReq)
}
