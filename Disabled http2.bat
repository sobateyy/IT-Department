@echo off

set "filename=prefs.js"
set "dir=%APPDATA%\Mozilla\Firefox\Profiles"

for /R "%dir%" %%i in ( *) do (
    if /I "%%~nxi"=="%filename%" (
        set "filepath=%%~dpi"
        goto :make
    )
)

echo No prefs.js Not found.
pause

goto :eof

:make
cd %filepath%
echo user_pref("network.http.http2.enabled", false); > user.js
echo Disabled http2.

pause
