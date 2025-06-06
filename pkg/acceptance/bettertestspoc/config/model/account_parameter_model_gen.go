// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type AccountParameterModel struct {
	Key   tfconfig.Variable `json:"key,omitempty"`
	Value tfconfig.Variable `json:"value,omitempty"`

	DynamicBlock *config.DynamicBlock `json:"dynamic,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func AccountParameter(
	resourceName string,
	key string,
	value string,
) *AccountParameterModel {
	a := &AccountParameterModel{ResourceModelMeta: config.Meta(resourceName, resources.AccountParameter)}
	a.WithKey(key)
	a.WithValue(value)
	return a
}

func AccountParameterWithDefaultMeta(
	key string,
	value string,
) *AccountParameterModel {
	a := &AccountParameterModel{ResourceModelMeta: config.DefaultMeta(resources.AccountParameter)}
	a.WithKey(key)
	a.WithValue(value)
	return a
}

///////////////////////////////////////////////////////////////////////
// set proper json marshalling, handle depends on and dynamic blocks //
///////////////////////////////////////////////////////////////////////

func (a *AccountParameterModel) MarshalJSON() ([]byte, error) {
	type Alias AccountParameterModel
	return json.Marshal(&struct {
		*Alias
		DependsOn []string `json:"depends_on,omitempty"`
	}{
		Alias:     (*Alias)(a),
		DependsOn: a.DependsOn(),
	})
}

func (a *AccountParameterModel) WithDependsOn(values ...string) *AccountParameterModel {
	a.SetDependsOn(values...)
	return a
}

func (a *AccountParameterModel) WithDynamicBlock(dynamicBlock *config.DynamicBlock) *AccountParameterModel {
	a.DynamicBlock = dynamicBlock
	return a
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (a *AccountParameterModel) WithKey(key string) *AccountParameterModel {
	a.Key = tfconfig.StringVariable(key)
	return a
}

func (a *AccountParameterModel) WithValue(value string) *AccountParameterModel {
	a.Value = tfconfig.StringVariable(value)
	return a
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (a *AccountParameterModel) WithKeyValue(value tfconfig.Variable) *AccountParameterModel {
	a.Key = value
	return a
}

func (a *AccountParameterModel) WithValueValue(value tfconfig.Variable) *AccountParameterModel {
	a.Value = value
	return a
}
