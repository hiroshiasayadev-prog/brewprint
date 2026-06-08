@echo off
setlocal

cd /d C:\Users\imved\projects\brewprint

echo [test] go test ./drmcp/src/...
go test ./drmcp/src/...
if errorlevel 1 exit /b 1

echo [build] v01/src/cmd/brewprint
go build -o bin\brewprint.exe .\v01\src\cmd\brewprint
if errorlevel 1 exit /b 1

echo [build] drmcp/src/cmd/design-records-mcp
go build -o bin\design-records-mcp.exe .\drmcp\src\cmd\design-records-mcp
if errorlevel 1 exit /b 1

echo OK