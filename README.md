# IT-Department

### Disable http2.bat
1. Find prefs.js in "%APPDATA%\Mozilla\Firefox\Profiles" and move it to the same directory
   <br>```cd %filename%```
3. Automate deactivation by creating "user.js" in the moved directory
   <br>```echo user_pref("network.http.http2.enabled", false); > user.js```
<p>Implemented in firefox, but can be replaced in chrome

### collection.go
This is an executable file that outputs the information obtained by the "systeminfo" command to a csv file.
It can be used to quickly gather information even if you have a large number of hosts!!

In order to eliminate unnecessary information, we have limited the conditions to "host name", "BIOS version", and "login user".
- Change the "HOST" to a name appropriate for your environment.
- The following is information on HOST1000 ~ HOST1500 and HOST5000 ~ HOST5500.
- To extract one by one, use "hostnames = append(hostnames, fmt.Sprintf("HOST3000"))".
	
```go
hostnames := make([]string, 0, 21)
for i := 1000; i <= 1500; i++ {
	hostnames = append(hostnames, fmt.Sprintf("HOST%d", i))
}
for i := 5000; i <= 5500; i++ {
	hostnames = append(hostnames, fmt.Sprintf("HOST%d", i))
}
hostnames = append(hostnames, fmt.Sprintf("HOST3000"))
hostnames = append(hostnames, fmt.Sprintf("HOST4000"))
```

- Control parallel processing by changing the following values
	
```go
// Limit the number of goroutines to be executed at the same time
// Match the specs of the host you want to run on
sem := make(chan struct{}, 100)
```
