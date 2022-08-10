package main

type (
	Module struct {
		Config ModuleConfig
		Sparks []*Spark
		Readme []byte
	}

	ModuleConfig struct {
		Package     string   `yaml:"package"`
		Label       string   `yaml:"label"`
		Repository  string   `yaml:"repository"`
		Description string   `yaml:"description"`
		Icon        string   `yaml:"icon"`
		Tags        []string `yaml:"tags"`
	}

	Spark struct {
		Config      SparkConfig
		Readme      []byte
		InputSchema []byte
	}

	SparkConfig struct {
		Label       string `yaml:"label"`
		Description string `yaml:"description"`
		Icon        string `yaml:"icon"`
	}
)

//func NewConfigFromInputs(action *ga.Action) (*Config, error) {
//	config := &Config{}
//	t := reflect.TypeOf(config).Elem()
//	v := reflect.ValueOf(config).Elem()
//
//	for i := 0; i < v.NumField(); i++ {
//		v.Field(i).SetString(action.GetInput(t.Field(i).Tag.Get("input")))
//	}
//	return config, nil
//}
