forfiles /p "C:\path" /s /m *.log /d -7 /c "cmd /c del /q @file is over 7days"
::刪除在該路徑內所有超過7天的log檔