@echo off
REM Windows batch script for Go project setup
REM Usage: new.bat projectname

if "%~1"=="" (
    echo Error: Project name is required
    echo Usage: new.bat projectname
    echo Example: new.bat day01
    exit /b 1
)

set NAME=%~1

if not exist "%NAME%" mkdir "%NAME%"
copy boilerplate.go "%NAME%\main.go"
type nul > "%NAME%\sample.txt"
type nul > "%NAME%\input.txt"

echo Project %NAME% created successfully!
