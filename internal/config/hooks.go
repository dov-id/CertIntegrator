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
					normMap[fmt.Sprint(key)] = value
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
			networks := make([]network, 0, len(v))

			for _, rawMap := range v {
				mapElem, ok := rawMap.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast map element to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))

				for key, value := range mapElem {
					normMap[fmt.Sprint(key)] = value
				}

				var data network

				err := figure.
					Out(&data).
					With(figure.BaseHooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out network data")
				}

				networks = append(networks, data)
			}

			return reflect.ValueOf(networks), nil
		default:
			return reflect.Value{}, fmt.Errorf("unexpected type to figure Config.Network[]")
		}
	},
}

var IntegratorHooks = figure.Hooks{
	"[]config.integrator": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case []interface{}:
			integrators := make([]integrator, 0, len(v))

			for _, rawMap := range v {
				mapElem, ok := rawMap.(map[interface{}]interface{})
				if !ok {
					return reflect.Value{}, errors.New("failed to cast map element to interface")
				}

				normMap := make(map[string]interface{}, len(mapElem))

				for key, value := range mapElem {
					normMap[fmt.Sprint(key)] = value
				}

				var data integrator

				err := figure.
					Out(&data).
					With(figure.BaseHooks).
					From(normMap).
					Please()
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to figure out integrators data")
				}

				integrators = append(integrators, data)
			}

			integratorsMap := make(map[string]string)

			for _, elem := range integrators {
				integratorsMap[elem.Network] = elem.Address
			}

			return reflect.ValueOf(integrators), nil
		default:
			return reflect.Value{}, fmt.Errorf("unexpected type to figure Config.Integrator[]")
		}
	},
}
