package config

type ConfigParam struct {
	JWTKey        []byte
	AccessControl map[string]map[string]map[string]bool
}

var CONFIG ConfigParam

func SetConfig() {
	accessctrl := make(map[string]map[string]map[string]bool)
	accessctrl["employees"] = map[string]map[string]bool{"POST": map[string]bool{"admin": true},
		"GET":    map[string]bool{"admin": true},
		"PATCH":  map[string]bool{"admin": true},
		"DELETE": map[string]bool{"admin": false}}

	accessctrl["departments"] = map[string]map[string]bool{"POST": map[string]bool{"admin": true},
		"GET":    map[string]bool{"admin": true},
		"PATCH":  map[string]bool{"admin": true},
		"DELETE": map[string]bool{"admin": false}}

	CONFIG = ConfigParam{JWTKey: []byte("mysecretkey"), AccessControl: accessctrl}
}
