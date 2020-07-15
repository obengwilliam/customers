/*
 * Customers API
 *
 * Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with US federal law and ensure authentic transfers. Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.  Customer phone numbers and addresses are stored and partially used in KYC/OFAC validation. Arbitrary key/value pairs can be stored for a Customer. Documents and Disclaimers, and their acknowledgement are also stored under a Customer as they're accepted. Bank Accounts, which can be validated with micro-deposits currently, are stored under each Customer.  ![](https://raw.githubusercontent.com/adamdecaf/customers/create-accounts/docs/images/customer.png)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// TransitAccountNumber struct for TransitAccountNumber
type TransitAccountNumber struct {
	// The base64 encoded encrypted account number for the bank account for transit across services. The symmetric key used is configured from TRANSIT_LOCAL_BASE64_KEY and callers of endpoints that respond with encrypted values need this same key to decrypt.
	AccountNumber string `json:"accountNumber,omitempty"`
}
