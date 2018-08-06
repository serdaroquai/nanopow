"# nanopow" 

a nano pow calculation implementation in golang. aws lambda deployable.

1. go build -o nanopow nanopow.go
2. build-lambda-zip.exe -o nanopow.zip nanopow
3. Upload .zip file using AWS lambda console