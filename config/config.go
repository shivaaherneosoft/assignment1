package config

type ConfigParam struct {
	JWTKey        []byte
	AccessControl map[string]map[string]bool
}

var CONFIG ConfigParam

func SetConfig() {
	accessctrl := make(map[string]map[string]bool)
	accessctrl["E01"] = map[string]bool{"admin": true}
	accessctrl["E02"] = map[string]bool{"admin": true}
	accessctrl["E03"] = map[string]bool{"admin": true}
	accessctrl["E04"] = map[string]bool{"admin": true}

	CONFIG = ConfigParam{JWTKey: []byte("mysecretkey"), AccessControl: accessctrl}
}
