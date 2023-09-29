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
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return nil
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, ok := helper.GetQueryInt(r, "id")
	if !ok {
		return nil
	}

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
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

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	deleteAccountReq := new(model.DeleteAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(deleteAccountReq); err != nil {
		return err
	}

	v := validate.Struct(deleteAccountReq)
	if !v.Validate() {
		return fmt.Errorf("missing required fields")
	}

	err := s.store.DeleteAccount(deleteAccountReq.ID)
	if err != nil {
		return err
	}

	account := &model.Account{ID: deleteAccountReq.ID}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
