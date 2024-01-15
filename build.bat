@echo off
go build  -ldflags "-w -s -H=windowsgui" -trimpath  -o licenseGenerator.exe
rem fyne package -os windows -icon Icon.png