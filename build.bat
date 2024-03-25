@echo off
set CGO_ENABLED=1
go build  -ldflags "-w -s -H=windowsgui" -trimpath  -o licenseGenerator.exe
rem fyne package -os windows -icon Icon.png