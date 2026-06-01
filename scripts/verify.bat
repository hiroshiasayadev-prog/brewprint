@echo off
setlocal

cd /d C:\Users\imved\projects\brewprint

echo [test] go test ./...
go test ./...
if errorlevel 1 exit /b 1

echo [build] cmd/brewprint
go build -o bin\brewprint.exe .\cmd\brewprint
if errorlevel 1 exit /b 1

echo [build] cmd/design-records-mcp
go build -o bin\design-records-mcp.exe .\cmd\design-records-mcp
if errorlevel 1 exit /b 1

echo OK