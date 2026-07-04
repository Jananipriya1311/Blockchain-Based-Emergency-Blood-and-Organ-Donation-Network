# Blockchain-Based Emergency Blood and Organ Donation Network

A decentralized healthcare platform built using **Hyperledger Fabric** for secure, transparent, and real-time blood and organ donation management using a hybrid **Proof of Availability (PoAv)** and **Delegated Byzantine Fault Tolerance (DBFT)** consensus mechanism.

---

## Overview

Traditional blood and organ donation systems rely on centralized databases, which often lead to delays, lack of transparency, and manual errors. This project leverages blockchain technology to provide a secure, transparent, and immutable platform for managing donors, hospitals, and emergency donation requests.

The platform automates donor management, hospital registration, and donation request approval using smart contracts while maintaining an immutable blockchain ledger.

---

## Problem Statement

Existing blood and organ donation systems face several challenges:

- Delays in matching donors with recipients.
- Lack of transparency between hospitals and donors.
- Manual record management leading to errors and duplication.
- Increased risk of fraud in centralized systems.
- Slow coordination during emergency situations.

---

## Proposed Solution

This project provides a blockchain-based emergency blood and organ donation network where:

- Hospitals act as validator nodes.
- Donors prove their real-time availability.
- Smart contracts automate donation workflows.
- Every transaction is securely recorded on the blockchain.
- All stakeholders share a transparent and immutable ledger.

---

## Features

- Donor Registration
- Hospital Registration
- Emergency Donation Request Management
- Real-Time Donor Availability Verification
- Smart Contract-Based Approval
- Secure Blockchain Ledger
- REST API Integration
- Transparent and Traceable Transactions

---

## Technology Stack

| Component | Technology |
|-----------|------------|
| Blockchain | Hyperledger Fabric |
| Smart Contracts | Go (Chaincode) |
| Backend | Node.js + Express |
| Frontend | HTML, CSS, JavaScript |
| Consensus | PoAv + Delegated BFT |

---

## Hybrid Consensus Mechanism

### Proof of Availability (PoAv)

- Verifies donor availability in real time.
- Ensures only available donors participate in donation requests.

### Delegated Byzantine Fault Tolerance (DBFT)

- Hospitals function as validator nodes.
- Validates transactions securely.
- Maintains ledger consistency.
- Protects against malicious activities.

---

## Workflow

1. Register hospitals.
2. Register donors.
3. Create an emergency donation request.
4. Verify donor availability using PoAv.
5. Validate the transaction using Delegated BFT.
6. Execute the smart contract.
7. Store the transaction on the blockchain.
8. Update the request status in real time.

---

## Smart Contract Functions

### Donor

- AddDonor()
- ReadDonor()
- GetAllDonors()

### Hospital

- AddHospital()
- ReadHospital()
- GetAllHospitals()

### Donation Requests

- AddDonationRequestWithDBFT()
- ApproveDonationRequestWithPoAv()
- ReadDonationRequest()
- UpdateDonationRequestStatus()
- GetAllDonationRequests()

---

## REST API

### GET Endpoints

```
/donors
/hospitals
/requests
```

### POST Endpoints

```
/add-donor
/add-hospital
/add-request
/update-request
```

---

## Project Structure

```text
Blockchain-Based-Emergency-Blood-and-Organ-Donation-Network/
│
├── backend/
│   ├── package.json
│   ├── package-lock.json
│   └── server.js
│
├── chaincode/
│   └── donation.go
│
├── frontend/
│   ├── index.html
│   └── style.css
│
└── README.md
```

---

## Advantages

- Decentralized architecture
- Secure and immutable records
- Transparent donation process
- Faster emergency response
- Reduced manual intervention
- Reliable audit trail

---

## Future Enhancements

- Mobile application
- AI-based donor recommendation
- GPS-based donor search
- Government healthcare integration
- SMS and Email notifications

---

## Team Members

- **A. Gesfetha** (23I201)
- **Jananipriya N** (23I223)
- **Rokitha R** (23I252)

---

## References

- Hyperledger Fabric Documentation
- Bitcoin: A Peer-to-Peer Electronic Cash System – Satoshi Nakamoto (2008)
- Ethereum Whitepaper
- Zheng, Z., Xie, S., Dai, H., Chen, X., & Wang, H. (2017). *An Overview of Blockchain Technology: Architecture, Consensus, and Future Trends.*

---

## License

This project is developed for academic and educational purposes.
