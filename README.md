# IT-Department

### Disable http2
1. Find prefs.js in "%APPDATA%\Mozilla\Firefox\Profiles" and move it to the same directory
   <br>```cd %filename%```
3. Automate deactivation by creating "user.js" in the moved directory
   <br>```echo user_pref("network.http.http2.enabled", false); > user.js```
<p>Implemented in firefox, but can be replaced in chrome
