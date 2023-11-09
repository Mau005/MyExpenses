@echo off
setlocal enabledelayedexpansion

set "output_folder=distribution"
mkdir "%output_folder%"

set "files_to_copy=config.yml"

:: Compilar para Linux en arquitectura AMD64
set "GOARCH=amd64"
set "GOOS=linux"
go build -o "%output_folder%\MyExpenses_linux_amd64" main.go

:: Compilar para Windows en arquitectura AMD64
set "GOARCH=amd64"
set "GOOS=windows"
go build -o "%output_folder%\MyExpenses_windows_amd64.exe" main.go

:: Compilar para macOS en arquitectura AMD64
set "GOARCH=amd64"
set "GOOS=darwin"
go build -o "%output_folder%\MyExpenses_macos_amd64" main.go

:: Copiar archivos
for %%f in (%files_to_copy%) do (
    copy "%%f" "%output_folder%"
)

endlocal