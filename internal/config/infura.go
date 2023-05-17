package config

import (
	"encoding/json"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type InfuraCfg struct {
	Key  string `json:"key"`
	Link string `json:"link"`
}

func (c *config) Infura() *InfuraCfg {
	return c.infura.Do(func() interface{} {
		var cfg InfuraCfg

		value, ok := os.LookupEnv("infura")
		if !ok {
			panic(errors.New("no infura env variable"))
		}

		err := json.Unmarshal([]byte(value), &cfg)
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out infura params from env variable"))
		}

		err = cfg.validate()
		if err != nil {
			panic(errors.Wrap(err, "failed to validate infura config"))
		}

		return &cfg
	}).(*InfuraCfg)
}

func (ic *InfuraCfg) validate() error {
	return validation.Errors{
		"key":  validation.Validate(ic.Key, validation.Required),
		"link": validation.Validate(ic.Link, validation.Required),
	}.Filter()
}
