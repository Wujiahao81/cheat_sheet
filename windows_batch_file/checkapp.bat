@echo off
TASKKILL /F /IM APP_NAME.exe /FI "STATUS eq NOT RESPONDING"
tasklist /FI "IMAGENAME eq APP_NAME.exe" 2>NUL | find /I /N "APP_NAME.exe" >NUL
if "%ERRORLEVEL%"=="0" (
exit 
) else (
PUSHD
cd /d "C:\your APP_NAME.exe path\"
start APP_NAME.exe
POPD
)