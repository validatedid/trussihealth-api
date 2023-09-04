# TruSSI Health API

This project contains an API component used in [TruSSI Health](https://ontochain.ngi.eu/content/trussihealth-decentralized-trustworthy-health-information-exchange-patients-self-sovereign). Specifically, this API is used by health centers to deliver health credentials to patients and by doctors to retrieve health data securely saved.

This API exposes 2 endpoints:

* `POST /health-data` encrypts data and saves it into IPFS plus deliver an eSealed Verifiable Credential to data owner.
* `GET /health-data/<documentHash>` retrieves document from IPFS and decrypts it.

You can find API requests examples importable into Postman in `examples` directory.

## Prerequisites


### 1 - Run TruSSIHealth signer

Please, clone [TruSSIHealth Signer repository](https://github.com/validatedid/trussihealth-signer) that contains independent components for verifiable credential issuance and eSeal. Follow the instructions and run it.

### 2 - Set Environment variables

This project requires to have a `.env` file with all these environment variables set in the root directory of the project. Copy paste `.env.example` into a `.env` file and place it in the root directory.

This file must contain the following input:
```
APP_ENV=deployment
PORT=3011

# TrussiHealth API Authentication 
PASSWORD=

# TrussiHealth API encryption key for storing data
ENCRYPTION_KEY=

# IPFS endpoint where data is stored/retrieved
IPFS_URL=

# TrussiHealth VIDchain Entity DID
ISSUER_DID=

# TrussiHealth Certificate's password to eseal issued VCs
CERTIFICATE_PASSWORD=

# TruSSIHealth Signer endpoint for credential issuance
VC_API=http://127.0.0.1:3000/verifiable-credential/v1/signatures

# TruSSIHealth Signer endpoint for credential eseal
ESEAL_API=http://127.0.0.1:3001/eseal/v1/signatures
```

Notice that for these two last you will have to provide the right endpoints where you decide to host TruSSIHealth Signer.

## Run TruSSIHealth API

This project can be run locally or using Docker. In any case, please set the environment variables first.

## Golang


This project is developed in Golang. To run it you will need to [install Golang](https://go.dev/doc/install) in your machine.

Once you have Golang installed, then you simply have to run the following command in the source directory:
```
go run src/main.go
```

Alternatively, if you want to build an executable to be distributed, you can run:

```
go build -o GoExecutable src/main.go
```
You can run this executable like:
```
./GoExecutable
```

**NOTE: This repository has been tested with go version 1.20.3**

## Docker

First of all, [install Docker](https://docs.docker.com/engine/install/) on your machine. Afterwards, uncomment line 12 at `Dockerfile`:
```
# Uncomment this line to run the image locally
COPY .env .
```

Now, you are ready to build the docker image and run it:
```
docker build -t trussihealth:latest .
docker run --env-file=.env -p 3011:3011 -t trussihealth:latest
```
