# TruSSI Health API

This project contains an API component used in [TruSSI Health](https://ontochain.ngi.eu/content/trussihealth-decentralized-trustworthy-health-information-exchange-patients-self-sovereign). Specifically, this API is used by health centers to deliver health credentials to patients and by doctors to retrieve health data securely saved.

This API exposes 2 endpoints:

* `POST /health-data` encrypts data and saves it into IPFS plus deliver an eSealed Verifiable Credential to data owner.
* `GET /health-data/<documentHash>` retrieves document from IPFS and decrypts it.

You can find API requests examples importable into Postman in `examples` directory.

## Environment variables

This project requires to have a `.env` file with all these environment variables set in the root directory of the project.
This file must contain the following input:
```
VIDCHAIN_API=http://labs.vidchain.net
APP_ENV=deployment
PORT=3011
```

plus 

```
# TrussiHealth API Authentication 
PASSWORD=

# TrussiHealth API encryption key for storing data
ENCRYPTION_KEY=

# IPFS endpoint where data is stored/retrieved
IPFS_URL=

# Authentication information towards VIDchain API
TRUSSIHEALTH_ASSERTION=

# TrussiHealth VIDchain Entity DID
ISSUER_DID=

# TrussiHealth Certificate's password to eseal issued VCs
CERTIFICATE_PASSWORD=
```

Notice that for these last we can not set the secret values on the repository for obvious security reasons. 
Please, contact [VIDchain support](mailto:support@vidchain.org) for more details.

## Run it locally

This project can be run locally or using Docker. In any case, please set the environment variables first.

### Environment variables

To run the code locally you will need to create a `.env` file in the root directory. Copy the keys of `.env.example` to your `.env` file and request the values to [Email Address](mailto:support@validatedid.com).

```
APP_ENV=local # Set this to local for running locally
PORT=3011 # You can choose in which port to run
TRUSSIHEALTH_ASSERTION= # Request value
ENCRYPTION_KEY= # Request value
VIDCHAIN_API= # Request value
IPFS_URL= # Request value
ISSUER_DID= # Request value
CERTIFICATE_PASSWORD= # Request value
PASSWORD= # Request value
```

Notice that to run this project locally using Validated ID IPFS node, you will need to request access as well.

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
