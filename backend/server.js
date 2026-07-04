const express = require('express');
const { exec } = require('child_process');
const path = require('path');
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.static('public'));

// Base path for TLS certificates
const ORG1_TLS = path.join(__dirname, '../organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt');
const ORG2_TLS = path.join(__dirname, '../organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt');
const ORDERER_CA = path.join(__dirname, '../organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem');

function runCommand(cmd, res) {
  exec(cmd, { env: process.env }, (error, stdout, stderr) => {
    if (error) {
      console.error(stderr);
      res.status(500).send(stderr);
    } else {
      res.send(stdout);
    }
  });
}

// ==================== GET Endpoints ====================

// Get all donors
app.get('/donors', (req, res) => {
  const cmd = `peer chaincode query -C mychannel -n donation -c '{"Args":["GetAllDonors"]}'`;
  runCommand(cmd, res);
});

// Get all hospitals
app.get('/hospitals', (req, res) => {
  const cmd = `peer chaincode query -C mychannel -n donation -c '{"Args":["GetAllHospitals"]}'`;
  runCommand(cmd, res);
});

// Get all donation requests
app.get('/requests', (req, res) => {
  const cmd = `peer chaincode query -C mychannel -n donation -c '{"Args":["GetAllDonationRequests"]}'`;
  runCommand(cmd, res);
});

// ==================== POST Endpoints ====================

// Add a donor
app.post('/add-donor', (req, res) => {
  const { id, name, bloodType, organType, availability, contact, lastDonationDate } = req.body;
  const cmd = `peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com --tls --cafile "${ORDERER_CA}" \
-C mychannel -n donation \
--peerAddresses localhost:7051 --tlsRootCertFiles "${ORG1_TLS}" \
--peerAddresses localhost:9051 --tlsRootCertFiles "${ORG2_TLS}" \
-c '{"function":"AddDonor","Args":["${id}","${name}","${bloodType}","${organType}","${availability}","${contact}","${lastDonationDate}"]}' --waitForEvent`;
  runCommand(cmd, res);
});

// Add a hospital
app.post('/add-hospital', (req, res) => {
  const { id, name, location } = req.body;
  const cmd = `peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com --tls --cafile "${ORDERER_CA}" \
-C mychannel -n donation \
--peerAddresses localhost:7051 --tlsRootCertFiles "${ORG1_TLS}" \
--peerAddresses localhost:9051 --tlsRootCertFiles "${ORG2_TLS}" \
-c '{"function":"AddHospital","Args":["${id}","${name}","${location}"]}' --waitForEvent`;
  runCommand(cmd, res);
});

// Add a donation request
app.post('/add-request', (req, res) => {
  const { requestID, donorID, hospitalID, organType } = req.body;
  const cmd = `peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com --tls --cafile "${ORDERER_CA}" \
-C mychannel -n donation \
--peerAddresses localhost:7051 --tlsRootCertFiles "${ORG1_TLS}" \
--peerAddresses localhost:9051 --tlsRootCertFiles "${ORG2_TLS}" \
-c '{"function":"AddDonationRequest","Args":["${requestID}","${donorID}","${hospitalID}","${organType}"]}' --waitForEvent`;
  runCommand(cmd, res);
});

// Update donation request status
app.post('/update-request', (req, res) => {
  const { requestID, status } = req.body;
  const cmd = `peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com --tls --cafile "${ORDERER_CA}" \
-C mychannel -n donation \
--peerAddresses localhost:7051 --tlsRootCertFiles "${ORG1_TLS}" \
--peerAddresses localhost:9051 --tlsRootCertFiles "${ORG2_TLS}" \
-c '{"function":"UpdateDonationRequestStatus","Args":["${requestID}","${status}"]}' --waitForEvent`;
  runCommand(cmd, res);
});

app.listen(PORT, () => console.log(`🌍 Donation Web running at http://localhost:${PORT}`));
