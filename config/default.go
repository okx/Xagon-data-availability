package config

import (
	"bytes"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// DefaultValues is the default configuration
const DefaultValues = `
PrivateKey = {Path = "/pk/test-member.keystore", Password = "testonly"}
PermitApiAddress = "0x0000000000000000000000000000000000000000"

[L1]
WsURL = "ws://127.0.0.1:8546"
RpcURL = "http://127.0.0.1:8545"
PolygonValidiumAddress = "0x975725832B4909Aab87D3604A0b501569dbBE7A9"
DataCommitteeAddress = "0x2f08F654B896208dD968aFdAEf733edC5FF62c03"
Timeout = "1m"
RetryPeriod = "5s"
BlockBatchSize = "64"

[Log]
Environment = "development" # "production" or "development"
Level = "info"
Outputs = ["stderr"]

[DB]
User = "committee_user"
Password = "committee_password"
Name = "committee_db"
Host = "x1-data-availability-db"
Port = "5432"
EnableLog = false
MaxConns = 200

[RPC]
Host = "0.0.0.0"
Port = 8444
ReadTimeout = "60s"
WriteTimeout = "60s"
MaxRequestsPerIPAndSecond = 500
`

// Default parses the default configuration values.
func Default() (*Config, error) {
	var cfg Config
	viper.SetConfigType("toml")

	err := viper.ReadConfig(bytes.NewBuffer([]byte(DefaultValues)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
