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

This project can be run locally or using Docker.

### Golang

This project is developed in Golang. To run it you will need to [install Golang](https://go.dev/doc/install) in your machine.
Once you have Golang installed, you simply have to run the following command in the source directory:

```
go run
```

If you want to build an executable to be distributed, you can run:

```
go build
```

## Docker

First of all, [install Docker](https://docs.docker.com/engine/install/) on your server. Afterwards, you will be able to build the docker image and run it:
```
docker build -t trussihealth:latest .
docker run --env-file=.env -p 3011:3011 -t trussihealth:latest
```
