package main

import (
	"gopkg.in/ini.v1"
	"os"
)

/*
The purpose of program to convert ENV variables to config.ini
Here is config.ini example:
```ini
[general]
domain=demo.cloudwebrtc.com
cert=configs/certs/cert.pem
key=configs/certs/key.pem
bind=0.0.0.0
port=8086
html_root=web

[turn]
public_ip=127.0.0.1
port=19302
realm=flutter-webrtc
```
We can convert ENV variables to config.ini by running:
```bash
export DOMAIN=demo.cloudwebrtc.com
export CERT=configs/certs/cert.pem
export KEY=configs/certs/key.pem
export BIND=0.0.0.0
export HTML_ROOT=web
export PUBLIC_IP=127.0.0.1
export REALM=flutter-webrtc
go run configs/main.go
*/
func main() {
	domain, domainExists := os.LookupEnv("DOMAIN")

	cert, certExists := os.LookupEnv("CERT")

	key, keyExists := os.LookupEnv("KEY")

	bind, bindExists := os.LookupEnv("BIND")
	if !bindExists {
		bind = "0.0.0.0"
	}

	port, portExists := os.LookupEnv("PORT")
	if !portExists {
		port = "8086"
	}

	htmlRoot, htmlRootExists := os.LookupEnv("HTML_ROOT")
	if !htmlRootExists {
		htmlRoot = "web"
	}

	publicIP, publicIPExists := os.LookupEnv("PUBLIC_IP")
	if !publicIPExists {
		publicIP = "127.0.0.1"
	}

	turnPort, turnPortExists := os.LookupEnv("TURN_PORT")
	if !turnPortExists {
		turnPort = "19302"
	}

	realm, realmExists := os.LookupEnv("REALM")
	if !realmExists {
		realm = "flutter-webrtc"
	}

	cfg := ini.Empty()

	//region General Section
	if _, err := cfg.NewSection("general"); err != nil {
		panic(err)
	}

	if domainExists {
		cfg.Section("general").Key("domain").SetValue(domain)
	}

	if certExists {
		cfg.Section("general").Key("cert").SetValue(cert)
	}

	if keyExists {
		cfg.Section("general").Key("key").SetValue(key)
	}

	cfg.Section("general").Key("bind").SetValue(bind)
	cfg.Section("general").Key("port").SetValue(port)
	cfg.Section("general").Key("html_root").SetValue(htmlRoot)

	//endregion

	//region Turn Section
	if _, err := cfg.NewSection("turn"); err != nil {
		panic(err)
	}

	cfg.Section("turn").Key("public_ip").SetValue(publicIP)
	cfg.Section("turn").Key("port").SetValue(turnPort)
	cfg.Section("turn").Key("realm").SetValue(realm)

	//endregion

	if err := cfg.SaveTo("configs/config.ini"); err != nil {
		panic(err)
	}
}
