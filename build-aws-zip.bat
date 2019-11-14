set GOOS=linux
go build
mkdir target
move .\awslambda-go-georadius-checker .\target\georadius-checker
build-lambda-zip --output .\target\georadius-checker.zip .\target\georadius-checker
