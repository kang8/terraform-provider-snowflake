// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type StreamOnExternalTableModel struct {
	Database           tfconfig.Variable `json:"database,omitempty"`
	Schema             tfconfig.Variable `json:"schema,omitempty"`
	Name               tfconfig.Variable `json:"name,omitempty"`
	At                 tfconfig.Variable `json:"at,omitempty"`
	Before             tfconfig.Variable `json:"before,omitempty"`
	Comment            tfconfig.Variable `json:"comment,omitempty"`
	CopyGrants         tfconfig.Variable `json:"copy_grants,omitempty"`
	ExternalTable      tfconfig.Variable `json:"external_table,omitempty"`
	FullyQualifiedName tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	InsertOnly         tfconfig.Variable `json:"insert_only,omitempty"`
	Stale              tfconfig.Variable `json:"stale,omitempty"`
	StreamType         tfconfig.Variable `json:"stream_type,omitempty"`

	DynamicBlock *config.DynamicBlock `json:"dynamic,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func StreamOnExternalTable(
	resourceName string,
	database string,
	schema string,
	name string,
	externalTable string,
) *StreamOnExternalTableModel {
	s := &StreamOnExternalTableModel{ResourceModelMeta: config.Meta(resourceName, resources.StreamOnExternalTable)}
	s.WithDatabase(database)
	s.WithSchema(schema)
	s.WithName(name)
	s.WithExternalTable(externalTable)
	return s
}

func StreamOnExternalTableWithDefaultMeta(
	database string,
	schema string,
	name string,
	externalTable string,
) *StreamOnExternalTableModel {
	s := &StreamOnExternalTableModel{ResourceModelMeta: config.DefaultMeta(resources.StreamOnExternalTable)}
	s.WithDatabase(database)
	s.WithSchema(schema)
	s.WithName(name)
	s.WithExternalTable(externalTable)
	return s
}

///////////////////////////////////////////////////////////////////////
// set proper json marshalling, handle depends on and dynamic blocks //
///////////////////////////////////////////////////////////////////////

func (s *StreamOnExternalTableModel) MarshalJSON() ([]byte, error) {
	type Alias StreamOnExternalTableModel
	return json.Marshal(&struct {
		*Alias
		DependsOn []string `json:"depends_on,omitempty"`
	}{
		Alias:     (*Alias)(s),
		DependsOn: s.DependsOn(),
	})
}

func (s *StreamOnExternalTableModel) WithDependsOn(values ...string) *StreamOnExternalTableModel {
	s.SetDependsOn(values...)
	return s
}

func (s *StreamOnExternalTableModel) WithDynamicBlock(dynamicBlock *config.DynamicBlock) *StreamOnExternalTableModel {
	s.DynamicBlock = dynamicBlock
	return s
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (s *StreamOnExternalTableModel) WithDatabase(database string) *StreamOnExternalTableModel {
	s.Database = tfconfig.StringVariable(database)
	return s
}

func (s *StreamOnExternalTableModel) WithSchema(schema string) *StreamOnExternalTableModel {
	s.Schema = tfconfig.StringVariable(schema)
	return s
}

func (s *StreamOnExternalTableModel) WithName(name string) *StreamOnExternalTableModel {
	s.Name = tfconfig.StringVariable(name)
	return s
}

// at attribute type is not yet supported, so WithAt can't be generated

// before attribute type is not yet supported, so WithBefore can't be generated

func (s *StreamOnExternalTableModel) WithComment(comment string) *StreamOnExternalTableModel {
	s.Comment = tfconfig.StringVariable(comment)
	return s
}

func (s *StreamOnExternalTableModel) WithCopyGrants(copyGrants bool) *StreamOnExternalTableModel {
	s.CopyGrants = tfconfig.BoolVariable(copyGrants)
	return s
}

func (s *StreamOnExternalTableModel) WithExternalTable(externalTable string) *StreamOnExternalTableModel {
	s.ExternalTable = tfconfig.StringVariable(externalTable)
	return s
}

func (s *StreamOnExternalTableModel) WithFullyQualifiedName(fullyQualifiedName string) *StreamOnExternalTableModel {
	s.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return s
}

func (s *StreamOnExternalTableModel) WithInsertOnly(insertOnly string) *StreamOnExternalTableModel {
	s.InsertOnly = tfconfig.StringVariable(insertOnly)
	return s
}

func (s *StreamOnExternalTableModel) WithStale(stale bool) *StreamOnExternalTableModel {
	s.Stale = tfconfig.BoolVariable(stale)
	return s
}

func (s *StreamOnExternalTableModel) WithStreamType(streamType string) *StreamOnExternalTableModel {
	s.StreamType = tfconfig.StringVariable(streamType)
	return s
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (s *StreamOnExternalTableModel) WithDatabaseValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Database = value
	return s
}

func (s *StreamOnExternalTableModel) WithSchemaValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Schema = value
	return s
}

func (s *StreamOnExternalTableModel) WithNameValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Name = value
	return s
}

func (s *StreamOnExternalTableModel) WithAtValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.At = value
	return s
}

func (s *StreamOnExternalTableModel) WithBeforeValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Before = value
	return s
}

func (s *StreamOnExternalTableModel) WithCommentValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Comment = value
	return s
}

func (s *StreamOnExternalTableModel) WithCopyGrantsValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.CopyGrants = value
	return s
}

func (s *StreamOnExternalTableModel) WithExternalTableValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.ExternalTable = value
	return s
}

func (s *StreamOnExternalTableModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.FullyQualifiedName = value
	return s
}

func (s *StreamOnExternalTableModel) WithInsertOnlyValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.InsertOnly = value
	return s
}

func (s *StreamOnExternalTableModel) WithStaleValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.Stale = value
	return s
}

func (s *StreamOnExternalTableModel) WithStreamTypeValue(value tfconfig.Variable) *StreamOnExternalTableModel {
	s.StreamType = value
	return s
}
