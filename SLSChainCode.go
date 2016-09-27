/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	err := CreateParticipantTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating Participants table: " + err.Error())
	}
	err = CreateLoanTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating Loan table: " + err.Error())
	}
	err = CreateLoanSharesTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating LoanShares table: " + err.Error())
	}
	err = CreateLoanRequestTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating LoanRequests table: " + err.Error())
	}
	err = CreateLoanInvitationTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating LoanInvitations table: " + err.Error())
	}
	err = CreateTransactionTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating Transactions table: " + err.Error())
	}
	err = CreateLoanReturnTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating LoanReturn table: " + err.Error())
	}
	err = CreateLoanSaleTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating LoanSale table: " + err.Error())
	}
	err = CreateLoanShareNegotiationTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating CreateLoanShareNegotiation table: " + err.Error())
	}
	err = createAccountTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating CreateLoanShareNegotiation table: " + err.Error())
	}

	t.populateInitialData(stub)

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}
	if function == "addParticipant" {
		return addParticipant(stub, args)
	}
	if function == "addLoan" {
		return addLoan(stub, args)
	}
	if function == "addLoanShare" {
		return addLoanShare(stub, args)
	}
	if function == "addLoanRequest" {
		return addLoanRequest(stub, args)
	}
	if function == "addLoanInvitation" {
		return addLoanInvitation(stub, args)
	}
	if function == "addTransaction" {
		return addTransaction(stub, args)
	}
	if function == "addLoanReturn" {
		return addLoanReturn(stub, args)
	}
	if function == "addLoanSale" {
		return addLoanSale(stub, args)
	}
	if function == "addLoanShareNegotiation" {
		return addLoanShareNegotiation(stub, args)
	}

	//Account
	if function == "addAccount" {
		return addAccount(stub, args)
	}
	if function == "updateAccountAmount" {
		return updateAccountAmount(stub, args)
	}

	// This function should not be invoked directly in the future!!!!!
	if function == "updateTableField" {
		return updateTableField(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	switch function {
	//Participants
	case "getParticipantsQuantity":
		return getParticipantsQuantity(stub, args)
	case "getParticipantsList":
		return getParticipantsList(stub, args)
	case "getParticipantsByType":
		return getParticipantsByType(stub, args)
	//Loans
	case "getLoansQuantity":
		return getLoansQuantity(stub, args)
	case "getLoansList":
		return getLoansList(stub, args)
	//LoanShares
	case "getLoanSharesQuantity":
		return getLoanSharesQuantity(stub, args)
	case "getLoanSharesList":
		return getLoanSharesList(stub, args)
	}

	//LoanRequest
	if function == "getLoanRequestsQuantity" { //read a variable
		res, err := getLoanRequestsQuantity(stub, args)
		if err != nil {
			return nil, errors.New("Error getting loan requests quantity")
		}
		return res, nil
	}
	if function == "getLoanRequestsList" {
		res, err := getLoanRequestsList(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan requests list")
		}
		return res, nil
	}

	//LoanInvitation
	if function == "getLoanInvitationsQuantity" { //read a variable
		res, err := getLoanInvitationsQuantity(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Invitations quantity")
		}
		return res, nil
	}
	if function == "getLoanInvitationsList" {
		res, err := getLoanInvitationsList(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Invitations list")
		}
		return res, nil
	}

	//Transactions
	if function == "getTransactionsQuantity" { //read a variable
		res, err := getTransactionsQuantity(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Transactions quantity")
		}
		return res, nil
	}
	if function == "getTransactionsList" {
		res, err := getTransactionsList(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Transactions list")
		}
		return res, nil
	}

	//Loan Return
	if function == "getLoanReturnsQuantity" { //read a variable
		res, err := getLoanReturnsQuantity(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Returns quantity")
		}
		return res, nil
	}
	if function == "getLoanReturnsList" {
		res, err := getLoanReturnsList(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Returns list")
		}
		return res, nil
	}

	//Loan Sale
	if function == "getLoanSalesQuantity" { //read a variable
		res, err := getLoanSalesQuantity(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Sales quantity")
		}
		return res, nil
	}
	if function == "getLoanSalesList" {
		res, err := getLoanSalesList(stub, args)
		if err != nil {
			return nil, errors.New("Error getting Loan Sales list")
		}
		return res, nil
	}

	//Loan Share Negotiation
	if function == "getLoanShareNegotiationsQuantity" { //read a variable
		return getLoanShareNegotiationsQuantity(stub, args)
	}
	if function == "getLoanShareNegotiationsList" {
		return getLoanShareNegotiationsList(stub, args)
	}

	//Account
	if function == "getAccountsQuantity" { //read a variable
		return getAccountsQuantity(stub, args)
	}
	if function == "getAccountsList" {
		return getAccountsList(stub, args)
	}

	//========================================================================
	// This function should not be invoked directly in the future!!!!!!!!!!!!!!!!!!!!
	if function == "countTableRows" {
		return countTableRows(stub, args)
	}
	if function == "filterTableByValue" {
		return filterTableByValue(stub, args)
	}
	//========================================================================

	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query")
}

func (t *SimpleChaincode) populateInitialData(stub *shim.ChaincodeStub) error {

	//Participants
	_, _ = addParticipant(stub, []string{"Bank of Associates & Companies LTD", "Bank"})
	_, _ = addParticipant(stub, []string{"Connected Colaborators Bank", "Bank"})
	_, _ = addParticipant(stub, []string{"Bank of Paper, Wilson & Bluemine LTD", "Bank"})
	_, _ = addParticipant(stub, []string{"Bill Gates", "Borrower"})
	_, _ = addParticipant(stub, []string{"Peter Froystad", "Borrower"})
	_, _ = addParticipant(stub, []string{"John Smith", "Lowyer"})

	//Accounts
	_, _ = addAccount(stub, []string{"1", "10000"})
	_, _ = addAccount(stub, []string{"2", "50000"})
	_, _ = addAccount(stub, []string{"3", "30000"})

	return nil
}
