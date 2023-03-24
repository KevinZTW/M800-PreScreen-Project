# M800 PreScreen Homework

## Introduction
This is a simple line bot, which have the following features:
- Receive a message from a user and store to MongoDB
- List all messages from a user
- Push a message to a user

## Quick Start
- create `config.yaml` under the `internal/pkg/config` folder as below
```yaml
mongoDB:
  url: mongodb://localhost:27018
  name: m800
lineMessageAPI:
    channelID: 
    channelSecrete: 
    channelAccessToken: 
```

- run `make` to build and run the project which will 
  - create a folder in `~/data/m800` and spin up a mongodb container using that on port `:27018`
  - test and build the go binary
  - execute the binary, which will start the line bot server on port `:8080`
- In order to receive the callback, recommend to use `ngrok` to expose the local server to the internet
  - `ngrok http 8080`
  - update the webhook url in the line developer console
- you could test the endpoint using `test.sh` 
  - but need to revise the user id in the script to your own line user id