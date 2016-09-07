package main
 
import (
    "log"
    "net/http"
    "os/exec"
    "github.com/gorilla/mux"
)
 
func SensuNtpHandler(w http.ResponseWriter, r *http.Request) {
    var (
      cmdOut []byte
      err    error
    )
    cmdName := "/usr/local/bin/check-ntp.rb"
    if cmdOut, err = exec.Command(cmdName).Output(); err != nil {
      http.Error(w, string(cmdOut), 503)
    }
    w.Write([]byte(cmdOut))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/ntp", SensuNtpHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8080", r))
}
