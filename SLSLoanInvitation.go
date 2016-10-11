package main

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//Entity names
const LoanInvitationsTableName = "LoanInvitations"

//Column names
const LI_LoanInvitationIDColName = "LoanInvitationID"
const LI_ArrangerBankIDColName = "ArrangerBankID"
const LI_BorrowerIDColName = "BorrowerID"
const LI_LoanRequestIDColName = "LoanRequestID"
const LI_LoanTermColName = "LoanTerm"
const LI_AmountColName = "Amount"
const LI_InterestRateColName = "InterestRate"
const LI_InfoColName = "Info"
const LI_StatusColName = "Status"

var LI_ColumnNames []string

// ============================================================================================================================
//
// ============================================================================================================================

func CreateLoanInvitationTable(stub shim.ChaincodeStubInterface) error {
	LI_ColumnNames = []string{
		LI_LoanInvitationIDColName, LI_ArrangerBankIDColName, LI_BorrowerIDColName,
		LI_LoanRequestIDColName, LI_LoanTermColName, LI_AmountColName,
		LI_InterestRateColName, LI_InfoColName, LI_StatusColName}
	return createTable(stub, LoanInvitationsTableName, LI_ColumnNames)
}

func addLoanInvitation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != len(LI_ColumnNames)-1 {
		return nil, errors.New("Incorrect number of arguments in addLoanInvitation func. Expecting " + strconv.Itoa(len(LI_ColumnNames)-1))
	}

	check, err := checkRowPermissionsByBankId(stub, args[0])
	if !check {
		return nil, errors.New("Failed checking security in addLoanInvitation func or returned false: " + err.Error())
	}

	return nil, addRow(stub, LoanInvitationsTableName, args)
}

func getLoanInvitationsQuantity(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return countTableRows(stub, []string{LoanInvitationsTableName})
}

func getLoanInvitationsList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return filterTableByValue(stub, []string{LoanInvitationsTableName})
}

func updateLoanInvitationStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	loanInvitationID, newStatus := args[0], args[1]

	//Check if current user has priviledges to update Loan Negotiation Status
	arrangerBankId, err := getTableColValueByKey(stub, LoanInvitationsTableName, loanInvitationID, LI_ArrangerBankIDColName)
	if err != nil {
		return nil, errors.New("Error getting Participant Bank ID in updateLoanNegotiationStatus func: " + err.Error())
	}

	check, err := checkRowPermissionsByBankId(stub, arrangerBankId)
	if !check {
		return nil, errors.New("Failed checking security in updateLoanInvitationStatus func or returned false: " + err.Error())
	}

	return updateTableField(stub, []string{LoanInvitationsTableName, loanInvitationID, LI_StatusColName, newStatus})
}
