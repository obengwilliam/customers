package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	moovhttp "github.com/moov-io/base/http"

	"github.com/moov-io/customers/pkg/route"
)

func AddCustomerAddressRoutes(logger log.Logger, r *mux.Router, repo CustomerRepository) {
	r.Methods("POST").Path("/customers/{customerID}/addresses").HandlerFunc(createCustomerAddress(logger, repo))
	r.Methods("PUT").Path("/customers/{customerID}/addresses/{addressID}").HandlerFunc(updateCustomerAddress(logger, repo))
	r.Methods("DELETE").Path("/customers/{customerID}/addresses/{addressID}").HandlerFunc(deleteCustomerAddress(logger, repo))
}

func createCustomerAddress(logger log.Logger, repo CustomerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID, requestID := route.GetCustomerID(w, r), moovhttp.GetRequestID(r)
		if customerID == "" {
			return
		}

		var req address
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		if err := repo.addCustomerAddress(customerID, req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("customers", fmt.Sprintf("added address for customer=%s", customerID), "requestID", requestID)

		respondWithCustomer(logger, w, customerID, requestID, repo)
	}
}

func updateCustomerAddress(logger log.Logger, repo CustomerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = route.Responder(logger, w, r)

		customerID, addressId := route.GetCustomerID(w, r), getAddressId(w, r)
		if customerID == "" || addressId == "" {
			return
		}
		requestID := moovhttp.GetRequestID(r)
		var req updateCustomerAddressRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		if err := req.validate(); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		if err := repo.updateCustomerAddress(customerID, addressId, req); err != nil {
			logger.Log("customers", fmt.Sprintf("error updating customer's address: customer=%s address=%s: %v", customerID, addressId, err), "requestID", requestID)
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("customers", fmt.Sprintf("updating address=%s for customer=%s", addressId, customerID), "requestID", requestID)
		respondWithCustomer(logger, w, customerID, requestID, repo)
	}
}

func deleteCustomerAddress(logger log.Logger, repo CustomerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = route.Responder(logger, w, r)

		customerID, addressId := route.GetCustomerID(w, r), getAddressId(w, r)
		if customerID == "" || addressId == "" {
			return
		}
		requestID := moovhttp.GetRequestID(r)

		err := repo.deleteCustomerAddress(customerID, addressId)
		if err != nil {
			moovhttp.Problem(w, err)
			logger.Log("customers", fmt.Sprintf("error deleting customer's address: customer=%s address=%s: %v", customerID, addressId, err), "requestID", requestID)
			return
		}

		logger.Log("customers", fmt.Sprintf("successfully deleted address=%s for customer=%s", addressId, customerID), "requestID", requestID)
		w.WriteHeader(http.StatusNoContent)
	}
}

func getAddressId(w http.ResponseWriter, r *http.Request) string {
	varName := "addressID"
	v, ok := mux.Vars(r)[varName]
	if !ok || v == "" {
		moovhttp.Problem(w, fmt.Errorf("path variable %s not found in url", varName))
		return ""
	}
	return v
}

type updateCustomerAddressRequest struct {
	Type       string `json:"type"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
	Validated  bool   `json:"validated"`
}

func (req *updateCustomerAddressRequest) validate() error {
	switch strings.ToLower(req.Type) {
	case "primary", "secondary":
		return nil
	default:
		return fmt.Errorf("updateCustomerAddressRequest: unknown type: %s", req.Type)
	}
}
