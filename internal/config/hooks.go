package config

import (
	"fmt"
	"reflect"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var ContractHooks = figure.Hooks{
	"[]config.Contract": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case []interface{}:
			contracts := make([]Contract, 0, len(v))

			for _, rawMap := range v {
				mapElem, ok := rawMap.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast map element to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))

				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data Contract

				err := figure.
					Out(&data).
					With(figure.BaseHooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out contract data")
				}

				contracts = append(contracts, data)
			}

			return reflect.ValueOf(contracts), nil
		default:
			return reflect.Value{}, fmt.Errorf("unexpected type to figure Config.Contract[]")
		}
	},
}

var NetworkHooks = figure.Hooks{
	"[]config.network": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case []interface{}:
			contracts := make([]network, 0, len(v))

			for _, rawMap := range v {
				mapElem, ok := rawMap.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast map element to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))

				for key, value := range mapElem {
					strKey := fmt.Sprintf("%v", key)
					normMap[strKey] = value
				}

				var data network

				err := figure.
					Out(&data).
					With(figure.BaseHooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out contract data")
				}

				contracts = append(contracts, data)
			}

			return reflect.ValueOf(contracts), nil
		default:
			return reflect.Value{}, fmt.Errorf("unexpected type to figure Config.Contract[]")
		}
	},
}
