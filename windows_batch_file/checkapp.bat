@echo off
::tasklist /fi "<篩選器> <運算子> <值>"
TASKKILL /F /IM APP_NAME.exe /FI "STATUS eq NOT RESPONDING"
::taskkill /im <映像名稱>
tasklist /FI "IMAGENAME eq APP_NAME.exe" 2>NUL | find /I /N "APP_NAME.exe" >NUL
if "%ERRORLEVEL%"=="0" (
exit 
) else (
::PAUSE 會暫停，cmd 視窗就會停住，等你按任意鍵繼續
PUSHD
cd /d "C:\your APP_NAME.exe path\"
start APP_NAME.exe
POPD
)