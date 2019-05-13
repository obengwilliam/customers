/*
 * Customers API
 *
 * Customers ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Customer struct {
	// The unique identifier for the customer who owns the account
	CustomerId string `json:"customerId,omitempty"`
	// Given Name or First Name
	FirstName string `json:"firstName,omitempty"`
	// Middle Name
	MiddleName string `json:"middleName,omitempty"`
	// Surname or Last Name
	LastName string `json:"lastName,omitempty"`
	// Name Customer is preferred to be called
	NickName string `json:"nickName,omitempty"`
	// Customers name suffix. \"Jr\", \"PH.D.\"
	Suffix string `json:"suffix,omitempty"`
	// Legal date of birth
	BirthDate time.Time `json:"birthDate,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	// Used for localization of documents
	Culture string `json:"culture,omitempty"`
	// State of the customer
	Status string `json:"status,omitempty"`
	// Primary email address of customer name@domain.com
	Email     string    `json:"email,omitempty"`
	Phones    []Phone   `json:"phones,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Last time the object was modified
	LastModified time.Time `json:"lastModified,omitempty"`
}
