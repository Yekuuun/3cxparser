@echo off
REM Script de build pour Windows et macOS
echo Compilation en cours...

REM Compilation pour Windows (exécutable .exe)
SET GOOS=windows
SET GOARCH=amd64
go build -o 3cxparser.exe main.go
IF %ERRORLEVEL% NEQ 0 (
    echo Erreur lors de la compilation pour Windows.
    exit /b %ERRORLEVEL%
)
echo Windows build success : 3cxparser.exe
