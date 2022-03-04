forfiles /p "C:\path" /s /m *.jpg /d -1 /c "cmd /c del /q @file"
::刪除在該路徑內超過一天的所有jpg檔