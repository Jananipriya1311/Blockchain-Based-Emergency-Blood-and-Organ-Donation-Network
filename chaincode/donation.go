package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing donors, hospitals, and donation requests
type SmartContract struct {
	contractapi.Contract
}

// Donor defines the structure for a donor
type Donor struct {
	ID               string `json:"ID"`
	Name             string `json:"Name"`
	BloodType        string `json:"BloodType"`
	OrganType        string `json:"OrganType"`
	Availability     string `json:"AvailabilityStatus"`
	ContactInfo      string `json:"ContactInfo"`
	LastDonationDate string `json:"LastDonationDate"`
}

// Hospital defines the structure for a hospital
type Hospital struct {
	ID                string   `json:"ID"`
	Name              string   `json:"Name"`
	Location          string   `json:"Location"`
	ValidatedRequests []string `json:"ValidatedRequests"`
}

// DonationRequest defines the structure for a donation request
type DonationRequest struct {
	RequestID  string `json:"RequestID"`
	DonorID    string `json:"DonorID"`
	HospitalID string `json:"HospitalID"`
	OrganType  string `json:"OrganType"`
	Status     string `json:"Status"` // Pending, Approved, Completed
}

// ==================== Donor Functions ====================

// InitLedger adds sample donors
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	donors := []Donor{
		{ID: "DONOR1", Name: "Alice", BloodType: "O+", OrganType: "", Availability: "Available", ContactInfo: "alice@example.com", LastDonationDate: "2025-10-01"},
		{ID: "DONOR2", Name: "Bob", BloodType: "B+", OrganType: "", Availability: "Available", ContactInfo: "bob@example.com", LastDonationDate: "2025-10-06"},
	}

	for _, donor := range donors {
		donorJSON, _ := json.Marshal(donor)
		err := ctx.GetStub().PutState(donor.ID, donorJSON)
		if err != nil {
			return fmt.Errorf("Failed to put donor to world state: %v", err)
		}
	}

	return nil
}

// AddDonor registers a new donor
func (s *SmartContract) AddDonor(ctx contractapi.TransactionContextInterface, id, name, bloodType, organType, availability, contact, lastDonationDate string) error {
	donor := Donor{
		ID:               id,
		Name:             name,
		BloodType:        bloodType,
		OrganType:        organType,
		Availability:     availability,
		ContactInfo:      contact,
		LastDonationDate: lastDonationDate,
	}
	donorJSON, _ := json.Marshal(donor)
	return ctx.GetStub().PutState(id, donorJSON)
}

// ReadDonor retrieves a donor by ID
func (s *SmartContract) ReadDonor(ctx contractapi.TransactionContextInterface, id string) (*Donor, error) {
	donorJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if donorJSON == nil {
		return nil, fmt.Errorf("Donor %s does not exist", id)
	}

	var donor Donor
	err = json.Unmarshal(donorJSON, &donor)
	if err != nil {
		return nil, err
	}
	return &donor, nil
}

// GetAllDonors returns all donors
func (s *SmartContract) GetAllDonors(ctx contractapi.TransactionContextInterface) ([]*Donor, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("DONOR", "DONORz")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var donors []*Donor
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var donor Donor
		json.Unmarshal(queryResponse.Value, &donor)
		donors = append(donors, &donor)
	}
	return donors, nil
}

// ==================== Hospital Functions ====================

// AddHospital registers a new hospital
func (s *SmartContract) AddHospital(ctx contractapi.TransactionContextInterface, id, name, location string) error {
	hospital := Hospital{
		ID:                id,
		Name:              name,
		Location:          location,
		ValidatedRequests: []string{},
	}
	hospitalJSON, _ := json.Marshal(hospital)
	return ctx.GetStub().PutState(id, hospitalJSON)
}

// ReadHospital retrieves a hospital by ID
func (s *SmartContract) ReadHospital(ctx contractapi.TransactionContextInterface, id string) (*Hospital, error) {
	hospitalJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read hospital from world state: %v", err)
	}
	if hospitalJSON == nil {
		return nil, fmt.Errorf("Hospital %s does not exist", id)
	}

	var hospital Hospital
	err = json.Unmarshal(hospitalJSON, &hospital)
	if err != nil {
		return nil, err
	}

	if hospital.ValidatedRequests == nil {
		hospital.ValidatedRequests = []string{}
	}

	return &hospital, nil
}

// GetAllHospitals returns all hospitals
func (s *SmartContract) GetAllHospitals(ctx contractapi.TransactionContextInterface) ([]*Hospital, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("HOSP", "HOSPz")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var hospitals []*Hospital
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var hospital Hospital
		json.Unmarshal(queryResponse.Value, &hospital)

		if hospital.ValidatedRequests == nil {
			hospital.ValidatedRequests = []string{}
		}

		hospitals = append(hospitals, &hospital)
	}
	return hospitals, nil
}

// ==================== DonationRequest Functions ====================

// AddDonationRequest registers a new donation request
func (s *SmartContract) AddDonationRequest(ctx contractapi.TransactionContextInterface, requestID, donorID, hospitalID, organType string) error {
	request := DonationRequest{
		RequestID:  requestID,
		DonorID:    donorID,
		HospitalID: hospitalID,
		OrganType:  organType,
		Status:     "Pending",
	}
	requestJSON, _ := json.Marshal(request)
	return ctx.GetStub().PutState(requestID, requestJSON)
}

// ReadDonationRequest retrieves a donation request by ID
func (s *SmartContract) ReadDonationRequest(ctx contractapi.TransactionContextInterface, requestID string) (*DonationRequest, error) {
	requestJSON, err := ctx.GetStub().GetState(requestID)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if requestJSON == nil {
		return nil, fmt.Errorf("DonationRequest %s does not exist", requestID)
	}

	var request DonationRequest
	err = json.Unmarshal(requestJSON, &request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// UpdateDonationRequestStatus updates the status of a donation request
func (s *SmartContract) UpdateDonationRequestStatus(ctx contractapi.TransactionContextInterface, requestID, status string) error {
	request, err := s.ReadDonationRequest(ctx, requestID)
	if err != nil {
		return err
	}

	request.Status = status
	requestJSON, _ := json.Marshal(request)
	return ctx.GetStub().PutState(requestID, requestJSON)
}

// GetAllDonationRequests returns all donation requests
func (s *SmartContract) GetAllDonationRequests(ctx contractapi.TransactionContextInterface) ([]*DonationRequest, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var requests []*DonationRequest
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var request DonationRequest
		json.Unmarshal(queryResponse.Value, &request)

		if request.RequestID != "" {
			requests = append(requests, &request)
		}
	}
	return requests, nil
}

// ==================== Main ====================

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating donation chaincode: %s", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting donation chaincode: %s", err)
	}
}
