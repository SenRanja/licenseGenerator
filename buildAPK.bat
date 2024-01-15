@echo off
set ANDROID_HOME=C:\Users\ranja\AppData\Local\Android\Sdk
set ANDROID_NDK_HOME=C:\Users\ranja\AppData\Local\Android\Sdk\ndk\25.1.8937393
fyne package -os android -appID com.github.SenRanja -name "LicenseGenerator"
