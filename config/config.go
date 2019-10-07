package config

type ConfigParam struct {
	JWTKey        []byte
	Users         map[string]map[string]string
	AccessControl map[string]map[string]map[string]bool
}

var CONFIG ConfigParam

func SetConfig() {
	users := make(map[string]map[string]string)
	users["user1"] = map[string]string{"password": "password1", "roleid": "admin"}
	users["user2"] = map[string]string{"password": "password2", "roleid": "guest"}

	accessctrl := make(map[string]map[string]map[string]bool)
	accessctrl["employees"] = map[string]map[string]bool{"POST": map[string]bool{"admin": true},
		"GET":    map[string]bool{"admin": true, "guest": true},
		"PATCH":  map[string]bool{"admin": true},
		"DELETE": map[string]bool{"admin": false}}

	accessctrl["departments"] = map[string]map[string]bool{"POST": map[string]bool{"admin": true},
		"GET":    map[string]bool{"admin": true, "guest": true},
		"PATCH":  map[string]bool{"admin": true},
		"DELETE": map[string]bool{"admin": false}}

	CONFIG = ConfigParam{JWTKey: []byte("mysecretkey"), Users: users, AccessControl: accessctrl}
}
