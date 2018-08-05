"# nanopow" 

a nano pow calculation implementation in golang. aws lambda deployable.

1. go build -o nanopow nanopow.go
2. build-lambda-zip.exe -o nanopow.zip nanopow
3. Upload .zip file using AWS lambda console

## TODO: 
1. introduce go routines for faster pow calculation to prevent timeouts
2. can use output of last blake2b hash as the next nonce since both output and work are 8 bytes
3. add a structure for json response