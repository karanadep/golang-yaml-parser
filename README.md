# golang-yaml-parser

yaml parser cli with below inputs and outputs.

```
inputs:
- json file of variables that needs to be replaced
- yaml file where variables needs to be replaced

output:
- generated yaml with variable substitution
```

# steps

```
# let us run go in a docker container

go version

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh

# validating code
go run cmd/main.go

# Build your application into a static binary
cd cmd
go build -o parser

# This will produce a compiled program - parser
./parser
```

# checklist of applications / microservices CI process.
- implement as cli
- add unit tests
- add code coverage
- package it as standalone binary for mac, windows
- add makefile to perform local actions
- add dockerfile and test publishing to docker local
- implement functionality to add releases and tags
- add github-actions worfklow to publish to dockerhub
- package it as helm chart & publish it to helm repository
- add sequence diagram to explain flow [release process, application flow]
- update readme for steps to follow
- create infrastructure stack to deploy this application to local kubernetes kind cluster
