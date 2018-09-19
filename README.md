# nanopow

a simple nano proof of work calculation implementation based on _golang.org/x/crypto/blake2b_ in golang that takes advantage of multiple cores.

## Usage

nanopow takes two arguments:

1. previous hash of the block. (32 bytes, hexadecimal)
2. an optional threshold value. (8 bytes, hexadecimal). Default value is 'FFFFFFC000000000' which is also threshold value for the main nano network.

outputs: a valid 8 bytes 'proof of work' for the given previos block hash. Output is represented as in nanode.co block explorer. Check out https://www.nanode.co/block/C1A200FA700013E578D82DC7A88F7666BA63E7420357A1E447B7B8CA9F1BDD23 'Work' row for an example output format.

`nanopow C1A200FA700013E578D82DC7A88F7666BA63E7420357A1E447B7B8CA9F1BDD23`

Output: `622b548e2abb112d`

## Installation

1. `go get github.com/serdaroquai/nanopow` 
2. `go build nanopow.go`

## Spread the love
xrb_3mjri5ywtysxm154k66rccf1n5opo6zs4gcn5ybq3g384a4fcpkrm5kzpngy
