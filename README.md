# Geo Radius Checker - Go AWS Lambda Ready

## Description
This project was made for GO and AWS Lambda study purposes.

It uses the Law of Cosines to calculate a distance between a point A and a point B coordinates and check if it is inside
a given radius (in meters).      

https://en.wikipedia.org/wiki/Law_of_cosines

## Dependencies

* AWS Lambda GO package - github.com/aws/aws-lambda-go/lambda

`go get github.com/aws/aws-lambda-go/lambda`

* AWS Lambda Deployment Package (Windows)

`go get github.com/aws/aws-lambda-go/cmd/build-lambda-zip`

## Steps to deploy

1. Install AWS Lambda GO package 
2. Install AWS Lambda Deployment Package 
3. Execute build-aws-zip.bat on PowerShell or CMD
4. Upload the zip file generated in /target folder to your AWS account
5. Just have fun =)

## Tests

`go test`
