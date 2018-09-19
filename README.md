# nanopow

a simple nano proof of work calculation implementation based on _golang.org/x/crypto/blake2b_ in golang that takes advantage of multiple cores.



## Usage

'nanopow <<previousBlockHash>>'

1. `go get github.com/serdaroquai/nanopow` 
2. `go build nanopow.go`
3. `nanopow C08C7727AC85E6DCC26D13B2FB9083AF05C17616C4999B966C2BBCD1586398E6 FFFFFFC000000000`
