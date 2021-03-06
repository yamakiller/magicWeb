@echo off

@echo installing golang web frame development environment

echo "Start installing development library"

set curPath=%cd%
git clone https://github.com/gin-gonic/gin  %GOPATH%/src/github.com/gin-gonic/gin
::git clone https://github.com/swaggo/files %GOPATH%/src/github.com/swaggo/files
git clone https://github.com/swaggo/swag  %GOPATH%/src/github.com/swaggo/swag
git clone https://github.com/mojocn/base64Captcha %GOPATH%/src/github.com/mojocn/base64Captcha
git clone -b v8 https://github.com/go-playground/validator.git %GOPATH%/src/gopkg.in/go-playground/validator.v8
git clone -b v1 https://github.com/go-sourcemap/sourcemap %GOPATH%/src/gopkg.in/sourcemap.v1

cd %GOPATH%/src 
go get -v github.com/swaggo/gin-swagger
go get -v github.com/swaggo/files

go install github.com/swaggo/swag/cmd/swag

echo "installation is complete"

@echo installed golang development environment complate