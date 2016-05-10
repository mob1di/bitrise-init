package models

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
)

// ScanResultModel ...
type ScanResultModel struct {
	OptionMap  map[string]OptionModel                               `json:"options,omitempty" yaml:"options,omitempty"`
	ConfigsMap map[string]map[string]bitriseModels.BitriseDataModel `json:"configs,omitempty" yaml:"configs,omitempty"`
}

/*
- title: Project (or Workspace) path
  envkey: BITRISE_PROJECT_PATH
  valuemap:
    ~/Develop/bitrise/sample-apps/sample-apps-ios-cocoapods/SampleAppWithCocoapods/SampleAppWithCocoapods.xcodeproj:
    - title: Scheme name
      envkey: BITRISE_SCHEME
      valuemap:
        SampleAppWithCocoapods: []
*/

// OptionValueMap ...
type OptionValueMap map[string]OptionModel

// OptionModel ...
type OptionModel struct {
	Title  string `json:"title,omitempty"  yaml:"title,omitempty"`
	EnvKey string `json:"env_key,omitempty"  yaml:"env_key,omitempty"`

	ValueMap OptionValueMap `json:"value_map,omitempty"  yaml:"value_map,omitempty"`
	Config   string         `json:"config,omitempty"  yaml:"config,omitempty"`
}

// NewOptionModel ...
func NewOptionModel(title, envKey string) OptionModel {
	return OptionModel{
		Title:  title,
		EnvKey: envKey,

		ValueMap: OptionValueMap{},
	}
}

// NewEmptyOptionModel ...
func NewEmptyOptionModel() OptionModel {
	return OptionModel{
		ValueMap: OptionValueMap{},
	}
}

// GetValues ...
func (option OptionModel) GetValues() []string {
	if option.Config != "" {
		return []string{option.Config}
	}

	values := []string{}
	for value := range option.ValueMap {
		values = append(values, value)
	}
	return values
}
