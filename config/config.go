package config

type ConfigParam struct {
	JWTKey []byte
}

var CONFIG ConfigParam

func SetConfig() {
	CONFIG = ConfigParam{JWTKey: []byte("mysecretkey")}
}
