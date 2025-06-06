// Code generated by config model builder generator; DO NOT EDIT.

package datasourcemodel

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/datasources"
)

type DatabaseModel struct {
	Name          tfconfig.Variable `json:"name,omitempty"`
	Comment       tfconfig.Variable `json:"comment,omitempty"`
	CreatedOn     tfconfig.Variable `json:"created_on,omitempty"`
	IsCurrent     tfconfig.Variable `json:"is_current,omitempty"`
	IsDefault     tfconfig.Variable `json:"is_default,omitempty"`
	Options       tfconfig.Variable `json:"options,omitempty"`
	Origin        tfconfig.Variable `json:"origin,omitempty"`
	Owner         tfconfig.Variable `json:"owner,omitempty"`
	RetentionTime tfconfig.Variable `json:"retention_time,omitempty"`

	*config.DatasourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func Database(
	datasourceName string,
	name string,
) *DatabaseModel {
	d := &DatabaseModel{DatasourceModelMeta: config.DatasourceMeta(datasourceName, datasources.Database)}
	d.WithName(name)
	return d
}

func DatabaseWithDefaultMeta(
	name string,
) *DatabaseModel {
	d := &DatabaseModel{DatasourceModelMeta: config.DatasourceDefaultMeta(datasources.Database)}
	d.WithName(name)
	return d
}

///////////////////////////////////////////////////////
// set proper json marshalling and handle depends on //
///////////////////////////////////////////////////////

func (d *DatabaseModel) MarshalJSON() ([]byte, error) {
	type Alias DatabaseModel
	return json.Marshal(&struct {
		*Alias
		DependsOn                 []string                      `json:"depends_on,omitempty"`
		SingleAttributeWorkaround config.ReplacementPlaceholder `json:"single_attribute_workaround,omitempty"`
	}{
		Alias:                     (*Alias)(d),
		DependsOn:                 d.DependsOn(),
		SingleAttributeWorkaround: config.SnowflakeProviderConfigSingleAttributeWorkaround,
	})
}

func (d *DatabaseModel) WithDependsOn(values ...string) *DatabaseModel {
	d.SetDependsOn(values...)
	return d
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (d *DatabaseModel) WithName(name string) *DatabaseModel {
	d.Name = tfconfig.StringVariable(name)
	return d
}

func (d *DatabaseModel) WithComment(comment string) *DatabaseModel {
	d.Comment = tfconfig.StringVariable(comment)
	return d
}

func (d *DatabaseModel) WithCreatedOn(createdOn string) *DatabaseModel {
	d.CreatedOn = tfconfig.StringVariable(createdOn)
	return d
}

func (d *DatabaseModel) WithIsCurrent(isCurrent bool) *DatabaseModel {
	d.IsCurrent = tfconfig.BoolVariable(isCurrent)
	return d
}

func (d *DatabaseModel) WithIsDefault(isDefault bool) *DatabaseModel {
	d.IsDefault = tfconfig.BoolVariable(isDefault)
	return d
}

func (d *DatabaseModel) WithOptions(options string) *DatabaseModel {
	d.Options = tfconfig.StringVariable(options)
	return d
}

func (d *DatabaseModel) WithOrigin(origin string) *DatabaseModel {
	d.Origin = tfconfig.StringVariable(origin)
	return d
}

func (d *DatabaseModel) WithOwner(owner string) *DatabaseModel {
	d.Owner = tfconfig.StringVariable(owner)
	return d
}

func (d *DatabaseModel) WithRetentionTime(retentionTime int) *DatabaseModel {
	d.RetentionTime = tfconfig.IntegerVariable(retentionTime)
	return d
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (d *DatabaseModel) WithNameValue(value tfconfig.Variable) *DatabaseModel {
	d.Name = value
	return d
}

func (d *DatabaseModel) WithCommentValue(value tfconfig.Variable) *DatabaseModel {
	d.Comment = value
	return d
}

func (d *DatabaseModel) WithCreatedOnValue(value tfconfig.Variable) *DatabaseModel {
	d.CreatedOn = value
	return d
}

func (d *DatabaseModel) WithIsCurrentValue(value tfconfig.Variable) *DatabaseModel {
	d.IsCurrent = value
	return d
}

func (d *DatabaseModel) WithIsDefaultValue(value tfconfig.Variable) *DatabaseModel {
	d.IsDefault = value
	return d
}

func (d *DatabaseModel) WithOptionsValue(value tfconfig.Variable) *DatabaseModel {
	d.Options = value
	return d
}

func (d *DatabaseModel) WithOriginValue(value tfconfig.Variable) *DatabaseModel {
	d.Origin = value
	return d
}

func (d *DatabaseModel) WithOwnerValue(value tfconfig.Variable) *DatabaseModel {
	d.Owner = value
	return d
}

func (d *DatabaseModel) WithRetentionTimeValue(value tfconfig.Variable) *DatabaseModel {
	d.RetentionTime = value
	return d
}
