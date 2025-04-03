package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

func main() {
    version := os.Getenv("VERSION")
    if version == "" {
        version = "dev"
    }

    hostname, _ := os.Hostname()

    addrs, _ := net.InterfaceAddrs()
    var ip string
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
            ip = ipnet.IP.String()
            break
        }
    }

    htmlContent := fmt.Sprintf(`
		<html>
		<head><title>docker_lab5</title></head>
		<body>
			<h1>Info</h1>
			<p>Adres IP serwera: %s</p>
			<p>Nazwa serwera (hostname): %s</p>
			<p>Wersja aplikacji: %s</p>
		</body>
		</html>
`, ip, hostname, version)

    err := ioutil.WriteFile("index.html", []byte(htmlContent), 0644)
    if err != nil {
        fmt.Printf("Błąd zapisu pliku index.html: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Plik index.html został wygenerowany.")
}
