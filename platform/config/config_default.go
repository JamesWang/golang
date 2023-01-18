package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (conf *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := conf.configData
	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

func (conf *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := conf.get(name)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}
	return
}

func (conf *DefaultConfig) GetString(name string) (result string, found bool) {
	value, found := conf.get(name)
	if found {
		result = value.(string)
	}
	return
}

func (conf *DefaultConfig) GetInt(name string) (result int, found bool) {
	value, found := conf.get(name)
	if found {
		result = value.(int)
	}
	return
}

func (conf *DefaultConfig) GetBool(name string) (result bool, found bool) {
	value, found := conf.get(name)
	if found {
		result = value.(bool)
	}
	return
}

func (conf *DefaultConfig) GetFloat(name string) (result float64, found bool) {
	value, found := conf.get(name)
	if found {
		result = value.(float64)
	}
	return
}

func (conf *DefaultConfig) GetStringDefault(name string, defVal string) (result string) {
	result, ok := conf.GetString(name)
	if !ok {
		result = defVal
	}
	return
}

func (conf *DefaultConfig) GetIntDefault(name string, defVal int) (result int) {
	result, ok := conf.GetInt(name)
	if !ok {
		result = defVal
	}
	return
}

func (conf *DefaultConfig) GetBoolDefault(name string, defVal bool) (result bool) {
	result, ok := conf.GetBool(name)
	if !ok {
		result = defVal
	}
	return
}

func (conf *DefaultConfig) GetFloatDefault(name string, defVal float64) (result float64) {
	result, ok := conf.GetFloat(name)
	if !ok {
		result = defVal
	}
	return
}
