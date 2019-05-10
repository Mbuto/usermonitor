package main

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "os"
    "strings"
    "bufio"
    "net/http"
    "net/url"
)


func main() {
u := ""
p := ""
/***************************************
args := os.Args[1:]
if len(args) == 0 {
fmt.Printf("ARGOMENTO MANCANTE\n")
//return
}
allarg := "%" + args[0] + "%"
allarg2 := ""
mono := true
if len(args) > 1 {
allarg2 = "%" + args[1] + "%"
mono = false
}
****************************************/
fmt.Printf("INFO: Checking .conf...\n")
if fileExists("ffmon.conf") {
f, err := os.Open("ffmon.conf")
if err != nil {
fmt.Printf("ERR-OPE: %v\n", err)
}
fmt.Fscanf(f, "%s\n%s\n", &u, &p)
fmt.Printf("INFO: .conf loaded.\n")
} else {
fmt.Printf("INFO: .conf does not exist: creating...\n")
reader := bufio.NewReader(os.Stdin)
fmt.Printf("Username: ")
u, _ = reader.ReadString('\n')
u = strings.Replace(u, "\r\n", "", -1)
u = strings.Replace(u, "\n", "", -1)

fmt.Printf("Password: ")
p, _ = reader.ReadString('\n')
p = strings.Replace(p, "\r\n", "", -1)
p = strings.Replace(p, "\n", "", -1)
f, err := os.Create("ffmon.conf")
if err != nil {
fmt.Printf("ERR-CRE: %v\n", err)
}
defer f.Close()
fmt.Fprintf(f, "%s\n%s\n",u,p)
fmt.Printf("INFO: .conf created.\n")
}
//fmt.Printf("/%s/%s/\n",u,p)
a := ""
fmt.Printf("INFO: Checking .auth...\n")
if fileExists("ffmon.auth") {
f, err := os.Open("ffmon.auth")
if err != nil {
fmt.Printf("ERR-OPE: %v\n", err)
}
fmt.Fscanf(f, "%s\n", &a)
fmt.Printf("INFO: .auth loaded.\n")
} else {
fmt.Printf("INFO: .auth does not exist: creating...\n")
//curl --data-urlencode "username=$U" --data-urlencode "password=$P" https://freefeed.net/v1/session
client := &http.Client{ }
data := url.Values{}
data.Set("username", u)
data.Add("password", p)
req, err := http.NewRequest("POST", "https://freefeed.net/v1/session", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		fmt.Printf("ERR-POST: %v\n", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERR-DO: %v\n", err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERR-READ: %v\n", err)
	}
	resp.Body.Close()
	if err != nil {
		fmt.Printf("ERR-CLOSE: %v\n", err)
	}
//	fmt.Println(string(f))
aa := strings.SplitAfter(string(f), "authToken\":\"")
aaa := aa[1]
a = aaa[:len(aaa)-2]
g, err := os.Create("ffmon.auth")
if err != nil {
fmt.Printf("ERR-CREAUTH: %v\n", err)
}
fmt.Fprintf(g, "%s\n",a)
g.Close()
fmt.Printf("INFO: .auth created.\n")

}
//fmt.Printf("authToken/%s/\n",a)
// curl -H "X-Authentication-Token: $T" https://freefeed.net/v1/users/$U/subscribers
fmt.Printf("INFO: Getting data from server...\n")
client2 := &http.Client{ }
ur := fmt.Sprintf("https://freefeed.net/v1/users/%s/subscribers", u)
req2, err2 := http.NewRequest("GET", ur, nil)
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	req2.Header.Set("X-Authentication-Token", a)
	if err2 != nil {
		fmt.Printf("ERR-GET: %v\n", err2)
	}
	resp2, err3 := client2.Do(req2)
	if err3 != nil {
		fmt.Printf("ERR-DO: %v\n", err3)
		return
	}

	h, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Printf("ERR-READ: %v\n", err)
	}
	resp2.Body.Close()
	if err != nil {
		fmt.Printf("ERR-CLOSE: %v\n", err)
	}
//	fmt.Println(string(h))
tut := ""
bb := strings.SplitAfter(string(h), "username\":\"")
for k := 1 ; k < len(bb); k++ {
	bbb:=strings.Index(bb[k], "\"")
	tut = tut + ":" + bb[k][:bbb]
//	fmt.Printf("%d %s\n", k, bb[k][:bbb])
}
//fmt.Printf("INFO: All data retrieved.\n%s\n", tut)
fmt.Printf("INFO: All data retrieved.\n")
if fileExists("followers") {
fmt.Printf("INFO: followers file already exists\n")
// something to do...
f, err := os.Open("followers")
if err != nil {
fmt.Printf("ERR-OPE: %v\n", err)
}
tut2 := ""
fmt.Fscanf(f, "%s\n", &tut2)
fmt.Printf("INFO: followers loaded.\n")
if tut == tut2 {
fmt.Printf("RESULT: No differences found.\n")
} else {
fmt.Printf("RESULT: Differences found.\n")
//fmt.Printf("OLD followers list: %s\n", tut)
//fmt.Printf("NEW followers list: %s\n", tut2)
findiff(tut, tut2)
fmt.Printf("INFO: Updating followers file...\n")
f, err := os.Create("followers")
if err != nil {
fmt.Printf("ERR-CRE: %v\n", err)
}
fmt.Fprintf(f, "%s\n",tut)
f.Close()
fmt.Printf("INFO: followers updated.\n")
}
} else {
fmt.Printf("INFO: Creating followers file...\n")
f, err := os.Create("followers")
if err != nil {
fmt.Printf("ERR-CRE: %v\n", err)
}
fmt.Fprintf(f, "%s\n",tut)
f.Close()
fmt.Printf("INFO: followers created.\n")
}
}

/***************************** fileExists ***************/
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

/***************************** findiff ***************/
func findiff(new string, old string) {
// search for NEW
bb := strings.SplitAfter(new, ":")
for k := 1 ; k < len(bb); k++ {
	bbb:=strings.Index(bb[k], ":")
	s := ""
	if bbb > 0 {
		s = bb[k][:bbb]
	} else {
		s = bb[k]
	}
	if lost(s, old) {
		fmt.Printf("NEW: %s\n", s)
	}
}
// search for LOST
cc := strings.SplitAfter(old, ":")
for k := 1 ; k < len(cc); k++ {
	ccc:=strings.Index(cc[k], ":")
	s := ""
	if ccc > 0 {
		s = cc[k][:ccc]
	} else {
		s = cc[k]
	}
	if lost(s, new) {
		fmt.Printf("LOST: %s\n", s)
	}
}
}

/***************************** lost ***************/
func lost(x string, y string) bool {
return (strings.Index(y, x) < 0)
}
