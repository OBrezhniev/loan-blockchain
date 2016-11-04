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

const isAuthenticationEnabled = false

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
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

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
	err = CreateLoanNegotiationTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating CreateLoanNegotiation table: " + err.Error())
	}
	err = createAccountTable(stub)
	if err != nil {
		return nil, errors.New("Failed creating CreateLoanNegotiation table: " + err.Error())
	}

	populateInitialData(stub, args)

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}

	//Participant
	if function == "addParticipant" {
		return addParticipant(stub, args)
	}

	//Loan
	if function == "addLoan" {
		return addLoan(stub, args)
	}

	//LoanShare
	if function == "addLoanShare" {
		return addLoanShare(stub, args)
	}

	//LoanRequest
	if function == "addLoanRequest" {
		return addLoanRequest(stub, args)
	}
	if function == "updateLoanRequest" {
		return updateLoanRequest(stub, args)
	}

	//Loan Invitation
	if function == "addLoanInvitation" {
		return addLoanInvitation(stub, args)
	}
	if function == "updateLoanInvitationStatus" {
		return updateLoanInvitationStatus(stub, args)
	}
	if function == "updateLoanInvitation" {
		return updateLoanInvitation(stub, args)
	}

	//Transaction
	if function == "addTransaction" {
		return addTransaction(stub, args)
	}

	//Loan Return
	if function == "addLoanReturn" {
		return addLoanReturn(stub, args)
	}

	//Loan Sale
	if function == "addLoanSale" {
		return addLoanSale(stub, args)
	}

	//Loan Negotiation
	if function == "addLoanNegotiation" {
		return addLoanNegotiation(stub, args)
	}
	if function == "updateLoanNegotiation" {
		return updateLoanNegotiation(stub, args)
	}
	if function == "updateLoanNegotiationStatus" {
		return updateLoanNegotiationStatus(stub, args)
	}
	if function == "updateParticipantBankComment" {
		return updateParticipantBankComment(stub, args)
	}

	//Account
	if function == "addAccount" {
		return addAccount(stub, args)
	}
	if function == "updateAccountAmount" {
		return updateAccountAmount(stub, args)
	}

	//========================================================================
	// Specific functions
	if function == "updateTableField" {
		return updateTableField(stub, args)
	}
	if function == "deleteRow" {
		return deleteRow(stub, args)
	}
	if function == "deleteRowsByColumnValue" {
		return deleteRowsByColumnValue(stub, args)
	}
	if function == "populateInitialData" {
		return populateInitialData(stub, args)
	}
	//========================================================================

	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	//Participants
	if function == "getParticipantsQuantity" {
		return getParticipantsQuantity(stub, args)
	}
	if function == "getParticipantsList" {
		return getParticipantsList(stub, args)
	}
	if function == "getParticipantsByType" {
		return getParticipantsByType(stub, args)
	}
	if function == "getParticipantsByKey" {
		return getParticipantsByKey(stub, args)
	}
	if function == "getParticipantsMaxKey" {
		return getParticipantsMaxKey(stub, args)
	}

	//Loans
	if function == "getLoansQuantity" {
		return getLoansQuantity(stub, args)
	}
	if function == "getLoansList" {
		return getLoansList(stub, args)
	}

	//LoanShares
	if function == "getLoanSharesQuantity" {
		return getLoanSharesQuantity(stub, args)
	}
	if function == "getLoanSharesList" {
		return getLoanSharesList(stub, args)
	}

	//LoanRequest
	if function == "getLoanRequestsQuantity" { //read a variable
		return getLoanRequestsQuantity(stub, args)
	}
	if function == "getLoanRequestsList" {
		return getLoanRequestsList(stub, args)
	}
	if function == "getLoanRequestByKey" {
		return getLoanRequestByKey(stub, args)
	}
	if function == "getLoanRequestsMaxKey" {
		return getLoanRequestsMaxKey(stub, args)
	}

	//LoanInvitation
	if function == "getLoanInvitationsQuantity" { //read a variable
		return getLoanInvitationsQuantity(stub, args)
	}
	if function == "getLoanInvitationsList" {
		return getLoanInvitationsList(stub, args)
	}
	if function == "getLoanInvitationByKey" {
		return getLoanInvitationByKey(stub, args)
	}
	if function == "getLoanInvitationsMaxKey" {
		return getLoanInvitationsMaxKey(stub, args)
	}

	//Transactions
	if function == "getTransactionsQuantity" { //read a variable
		return getTransactionsQuantity(stub, args)
	}
	if function == "getTransactionsList" {
		return getTransactionsList(stub, args)
	}

	//Loan Return
	if function == "getLoanReturnsQuantity" { //read a variable
		return getLoanReturnsQuantity(stub, args)
	}
	if function == "getLoanReturnsList" {
		return getLoanReturnsList(stub, args)
	}

	//Loan Sale
	if function == "getLoanSalesQuantity" { //read a variable
		return getLoanSalesQuantity(stub, args)
	}
	if function == "getLoanSalesList" {
		return getLoanSalesList(stub, args)
	}

	//Loan Negotiation
	if function == "getLoanNegotiationsQuantity" { //read a variable
		return getLoanNegotiationsQuantity(stub, args)
	}
	if function == "getLoanNegotiationsList" {
		return getLoanNegotiationsList(stub, args)
	}
	if function == "getLoanNegotiationByKey" {
		return getLoanNegotiationByKey(stub, args)
	}
	if function == "getLoanNegotiationsMaxKey" {
		return getLoanNegotiationsMaxKey(stub, args)
	}

	//Account
	if function == "getAccountsQuantity" { //read a variable
		return getAccountsQuantity(stub, args)
	}
	if function == "getAccountsList" {
		return getAccountsList(stub, args)
	}

	//========================================================================
	// Special functions
	if function == "countTableRows" {
		return countTableRows(stub, args)
	}
	if function == "filterTableByValue" {
		return filterTableByValue(stub, args)
	}
	/*if function == "printCallerCertificate" {
		return printCallerCertificate(stub)
	}*/
	if function == "getCertAttribute" {
		return getCertAttribute(stub, args)
	}
	if function == "getBankId" {
		return getBankId(stub, args)
	}
	if function == "getProjectsList" {
		return getProjectsList(stub, args)
	}
	//========================================================================

	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query")
}

func getCertAttribute(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments in getCertAttribute func. Expecting 1")
	}

	attrName := args[0]
	attribute, err := stub.ReadCertAttribute(attrName)
	if err != nil {
		return nil, errors.New("Failed retrieving Certificate Attribute '" + attrName + "' in getCertAttribute func: " + err.Error())
	}

	return []byte("Attribute '" + attrName + "': " + string(attribute)), nil
}

func getBankId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments in getBankId func. Expecting 0")
	}

	attrName := "bankid"
	attribute, err := stub.ReadCertAttribute(attrName)
	if err != nil {
		return nil, errors.New("Failed retrieving Certificate Attribute '" + attrName + "' in getBankId func: " + err.Error())
	}

	return []byte(string(attribute)), nil
}

func checkAttribute(stub shim.ChaincodeStubInterface, attrName, attrValue string) (bool, error) {
	if !isAuthenticationEnabled {
		return true, nil
	}
	// Why stub.VerifyAttribute is not used here?????? Consider using it.
	attribute, err := stub.ReadCertAttribute(attrName)
	if err != nil {
		return false, errors.New("Error checking role: " + err.Error())
	}
	if string(attribute) != attrValue {
		return false, errors.New("Current user attribute '" + attrName + "' value is '" + string(attribute) + "' but not '" + attrValue + "'")
	}
	return true, nil
}

func populateInitialData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//Participants
	_, _ = deleteRowsByColumnValue(stub, []string{ParticipantsTableName})
	//Adding banks with keys
	_, _ = addParticipant(stub, []string{"6", "SpareBank 1 SR-BANK", "Bank"})
	_, _ = addParticipant(stub, []string{"7", "DNB ASA", "Bank"})
	_, _ = addParticipant(stub, []string{"8", "Nationwide Building Society", "Bank"})
	_, _ = addParticipant(stub, []string{"9", "JPMorgan Chase & Co", "Bank"})
	_, _ = addParticipant(stub, []string{"10", "Barclays", "Bank"})
	_, _ = addParticipant(stub, []string{"11", "Mizuho Bank, Ltd.", "Bank"})

	//Accounts
	/*_, _ = deleteRowsByColumnValue(stub, []string{AccountsTableName})
	_, _ = addAccount(stub, []string{"1", "10000"})
	_, _ = addAccount(stub, []string{"2", "50000"})
	_, _ = addAccount(stub, []string{"3", "30000"})*/

	//Loan Request
	// "BorrowerID", "ArrangerBankID", "LoanSharesAmount", "ProjectRevenue", "ProjectName", "ProjectInformation",
	//"Company", "Website", "ContactPersonName", "ContactPersonSurname", "RequestDate",
	//"Status", "MarketAndIndustry"
	_, _ = deleteRowsByColumnValue(stub, []string{LoanRequestsTableName})
	_, _ = addLoanRequest(stub, []string{"Statoil ASA", "6", "1M", "1M", "Statoil ASA project", "Statoil ASA project info", "Statoil ASA", "www.statoil.com", "John", "Smith", "10-01-2016", "Draft", "Oil industry"})
	_, _ = addLoanRequest(stub, []string{"BP Global", "7", "1M", "1M", "BP Global project", "BP Global project info", "BP Global", "www.bp.com", "Peter", "Froystad", "10-01-2016", "Draft", "Oil industry"})

	//Loan Invitation
	//"ArrangerBankID","BorrowerID","LoanRequestID","LoanTerm","Amount","InterestRate","Info",
	//"Status", "Assets", "Convenants"
	_, _ = deleteRowsByColumnValue(stub, []string{LoanInvitationsTableName})
	_, _ = addLoanInvitation(stub, []string{"6", "Statoil ASA", "1", "2 years", "400M USD", "3%", "Statoil ASA loan invitation info", "Pending", "Assets Statoil ASA", "Convenats Statoil ASA"})
	_, _ = addLoanInvitation(stub, []string{"7", "BP Global", "2", "3 years", "750M USD", "5%", "BP Global loan invitation info", "Pending", "Assets BP Global", "Convenats BP Global"})

	//Loan Share Negotiation
	//"InvitationID","ParticipantBankID","Amount","NegotiationStatus", "ParticipantBankComment", "Date"
	_, _ = deleteRowsByColumnValue(stub, []string{LoanNegotiationsTableName})
	_, _ = addLoanNegotiation(stub, []string{"1", "6", "200 M USD", "INVITED", "Comment of SpareBank", "11-01-2016"})
	_, _ = addLoanNegotiation(stub, []string{"1", "9", "100 M USD", "INVITED", "Comment of JPMorgan", "12-01-2016"})
	_, _ = addLoanNegotiation(stub, []string{"1", "10", "100 M USD", "INVITED", "Comment of Barclays", "12-01-2016"})
	_, _ = addLoanNegotiation(stub, []string{"2", "7", "250 M USD", "INVITED", "Comment of Nationwide Building Society", "21-01-2016"})
	_, _ = addLoanNegotiation(stub, []string{"2", "9", "200 M USD", "INVITED", "Comment of JPMorgan", "22-01-2016"})
	_, _ = addLoanNegotiation(stub, []string{"2", "11", "300 M USD", "INVITED", "Comment of Mizuho Bank, Ltd.", "22-01-2016"})

	return nil, nil
}
