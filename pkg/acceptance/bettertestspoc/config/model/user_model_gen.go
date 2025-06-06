// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type UserModel struct {
	Name                                     tfconfig.Variable `json:"name,omitempty"`
	AbortDetachedQuery                       tfconfig.Variable `json:"abort_detached_query,omitempty"`
	Autocommit                               tfconfig.Variable `json:"autocommit,omitempty"`
	BinaryInputFormat                        tfconfig.Variable `json:"binary_input_format,omitempty"`
	BinaryOutputFormat                       tfconfig.Variable `json:"binary_output_format,omitempty"`
	ClientMemoryLimit                        tfconfig.Variable `json:"client_memory_limit,omitempty"`
	ClientMetadataRequestUseConnectionCtx    tfconfig.Variable `json:"client_metadata_request_use_connection_ctx,omitempty"`
	ClientPrefetchThreads                    tfconfig.Variable `json:"client_prefetch_threads,omitempty"`
	ClientResultChunkSize                    tfconfig.Variable `json:"client_result_chunk_size,omitempty"`
	ClientResultColumnCaseInsensitive        tfconfig.Variable `json:"client_result_column_case_insensitive,omitempty"`
	ClientSessionKeepAlive                   tfconfig.Variable `json:"client_session_keep_alive,omitempty"`
	ClientSessionKeepAliveHeartbeatFrequency tfconfig.Variable `json:"client_session_keep_alive_heartbeat_frequency,omitempty"`
	ClientTimestampTypeMapping               tfconfig.Variable `json:"client_timestamp_type_mapping,omitempty"`
	Comment                                  tfconfig.Variable `json:"comment,omitempty"`
	DateInputFormat                          tfconfig.Variable `json:"date_input_format,omitempty"`
	DateOutputFormat                         tfconfig.Variable `json:"date_output_format,omitempty"`
	DaysToExpiry                             tfconfig.Variable `json:"days_to_expiry,omitempty"`
	DefaultNamespace                         tfconfig.Variable `json:"default_namespace,omitempty"`
	DefaultRole                              tfconfig.Variable `json:"default_role,omitempty"`
	DefaultSecondaryRolesOption              tfconfig.Variable `json:"default_secondary_roles_option,omitempty"`
	DefaultWarehouse                         tfconfig.Variable `json:"default_warehouse,omitempty"`
	DisableMfa                               tfconfig.Variable `json:"disable_mfa,omitempty"`
	Disabled                                 tfconfig.Variable `json:"disabled,omitempty"`
	DisplayName                              tfconfig.Variable `json:"display_name,omitempty"`
	Email                                    tfconfig.Variable `json:"email,omitempty"`
	EnableUnloadPhysicalTypeOptimization     tfconfig.Variable `json:"enable_unload_physical_type_optimization,omitempty"`
	EnableUnredactedQuerySyntaxError         tfconfig.Variable `json:"enable_unredacted_query_syntax_error,omitempty"`
	ErrorOnNondeterministicMerge             tfconfig.Variable `json:"error_on_nondeterministic_merge,omitempty"`
	ErrorOnNondeterministicUpdate            tfconfig.Variable `json:"error_on_nondeterministic_update,omitempty"`
	FirstName                                tfconfig.Variable `json:"first_name,omitempty"`
	FullyQualifiedName                       tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	GeographyOutputFormat                    tfconfig.Variable `json:"geography_output_format,omitempty"`
	GeometryOutputFormat                     tfconfig.Variable `json:"geometry_output_format,omitempty"`
	JdbcTreatDecimalAsInt                    tfconfig.Variable `json:"jdbc_treat_decimal_as_int,omitempty"`
	JdbcTreatTimestampNtzAsUtc               tfconfig.Variable `json:"jdbc_treat_timestamp_ntz_as_utc,omitempty"`
	JdbcUseSessionTimezone                   tfconfig.Variable `json:"jdbc_use_session_timezone,omitempty"`
	JsonIndent                               tfconfig.Variable `json:"json_indent,omitempty"`
	LastName                                 tfconfig.Variable `json:"last_name,omitempty"`
	LockTimeout                              tfconfig.Variable `json:"lock_timeout,omitempty"`
	LogLevel                                 tfconfig.Variable `json:"log_level,omitempty"`
	LoginName                                tfconfig.Variable `json:"login_name,omitempty"`
	MiddleName                               tfconfig.Variable `json:"middle_name,omitempty"`
	MinsToBypassMfa                          tfconfig.Variable `json:"mins_to_bypass_mfa,omitempty"`
	MinsToUnlock                             tfconfig.Variable `json:"mins_to_unlock,omitempty"`
	MultiStatementCount                      tfconfig.Variable `json:"multi_statement_count,omitempty"`
	MustChangePassword                       tfconfig.Variable `json:"must_change_password,omitempty"`
	NetworkPolicy                            tfconfig.Variable `json:"network_policy,omitempty"`
	NoorderSequenceAsDefault                 tfconfig.Variable `json:"noorder_sequence_as_default,omitempty"`
	OdbcTreatDecimalAsInt                    tfconfig.Variable `json:"odbc_treat_decimal_as_int,omitempty"`
	Password                                 tfconfig.Variable `json:"password,omitempty"`
	PreventUnloadToInternalStages            tfconfig.Variable `json:"prevent_unload_to_internal_stages,omitempty"`
	QueryTag                                 tfconfig.Variable `json:"query_tag,omitempty"`
	QuotedIdentifiersIgnoreCase              tfconfig.Variable `json:"quoted_identifiers_ignore_case,omitempty"`
	RowsPerResultset                         tfconfig.Variable `json:"rows_per_resultset,omitempty"`
	RsaPublicKey                             tfconfig.Variable `json:"rsa_public_key,omitempty"`
	RsaPublicKey2                            tfconfig.Variable `json:"rsa_public_key_2,omitempty"`
	S3StageVpceDnsName                       tfconfig.Variable `json:"s3_stage_vpce_dns_name,omitempty"`
	SearchPath                               tfconfig.Variable `json:"search_path,omitempty"`
	SimulatedDataSharingConsumer             tfconfig.Variable `json:"simulated_data_sharing_consumer,omitempty"`
	StatementQueuedTimeoutInSeconds          tfconfig.Variable `json:"statement_queued_timeout_in_seconds,omitempty"`
	StatementTimeoutInSeconds                tfconfig.Variable `json:"statement_timeout_in_seconds,omitempty"`
	StrictJsonOutput                         tfconfig.Variable `json:"strict_json_output,omitempty"`
	TimeInputFormat                          tfconfig.Variable `json:"time_input_format,omitempty"`
	TimeOutputFormat                         tfconfig.Variable `json:"time_output_format,omitempty"`
	TimestampDayIsAlways24h                  tfconfig.Variable `json:"timestamp_day_is_always_24h,omitempty"`
	TimestampInputFormat                     tfconfig.Variable `json:"timestamp_input_format,omitempty"`
	TimestampLtzOutputFormat                 tfconfig.Variable `json:"timestamp_ltz_output_format,omitempty"`
	TimestampNtzOutputFormat                 tfconfig.Variable `json:"timestamp_ntz_output_format,omitempty"`
	TimestampOutputFormat                    tfconfig.Variable `json:"timestamp_output_format,omitempty"`
	TimestampTypeMapping                     tfconfig.Variable `json:"timestamp_type_mapping,omitempty"`
	TimestampTzOutputFormat                  tfconfig.Variable `json:"timestamp_tz_output_format,omitempty"`
	Timezone                                 tfconfig.Variable `json:"timezone,omitempty"`
	TraceLevel                               tfconfig.Variable `json:"trace_level,omitempty"`
	TransactionAbortOnError                  tfconfig.Variable `json:"transaction_abort_on_error,omitempty"`
	TransactionDefaultIsolationLevel         tfconfig.Variable `json:"transaction_default_isolation_level,omitempty"`
	TwoDigitCenturyStart                     tfconfig.Variable `json:"two_digit_century_start,omitempty"`
	UnsupportedDdlAction                     tfconfig.Variable `json:"unsupported_ddl_action,omitempty"`
	UseCachedResult                          tfconfig.Variable `json:"use_cached_result,omitempty"`
	UserType                                 tfconfig.Variable `json:"user_type,omitempty"`
	WeekOfYearPolicy                         tfconfig.Variable `json:"week_of_year_policy,omitempty"`
	WeekStart                                tfconfig.Variable `json:"week_start,omitempty"`

	DynamicBlock *config.DynamicBlock `json:"dynamic,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func User(
	resourceName string,
	name string,
) *UserModel {
	u := &UserModel{ResourceModelMeta: config.Meta(resourceName, resources.User)}
	u.WithName(name)
	return u
}

func UserWithDefaultMeta(
	name string,
) *UserModel {
	u := &UserModel{ResourceModelMeta: config.DefaultMeta(resources.User)}
	u.WithName(name)
	return u
}

///////////////////////////////////////////////////////////////////////
// set proper json marshalling, handle depends on and dynamic blocks //
///////////////////////////////////////////////////////////////////////

func (u *UserModel) MarshalJSON() ([]byte, error) {
	type Alias UserModel
	return json.Marshal(&struct {
		*Alias
		DependsOn []string `json:"depends_on,omitempty"`
	}{
		Alias:     (*Alias)(u),
		DependsOn: u.DependsOn(),
	})
}

func (u *UserModel) WithDependsOn(values ...string) *UserModel {
	u.SetDependsOn(values...)
	return u
}

func (u *UserModel) WithDynamicBlock(dynamicBlock *config.DynamicBlock) *UserModel {
	u.DynamicBlock = dynamicBlock
	return u
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (u *UserModel) WithName(name string) *UserModel {
	u.Name = tfconfig.StringVariable(name)
	return u
}

func (u *UserModel) WithAbortDetachedQuery(abortDetachedQuery bool) *UserModel {
	u.AbortDetachedQuery = tfconfig.BoolVariable(abortDetachedQuery)
	return u
}

func (u *UserModel) WithAutocommit(autocommit bool) *UserModel {
	u.Autocommit = tfconfig.BoolVariable(autocommit)
	return u
}

func (u *UserModel) WithBinaryInputFormat(binaryInputFormat string) *UserModel {
	u.BinaryInputFormat = tfconfig.StringVariable(binaryInputFormat)
	return u
}

func (u *UserModel) WithBinaryOutputFormat(binaryOutputFormat string) *UserModel {
	u.BinaryOutputFormat = tfconfig.StringVariable(binaryOutputFormat)
	return u
}

func (u *UserModel) WithClientMemoryLimit(clientMemoryLimit int) *UserModel {
	u.ClientMemoryLimit = tfconfig.IntegerVariable(clientMemoryLimit)
	return u
}

func (u *UserModel) WithClientMetadataRequestUseConnectionCtx(clientMetadataRequestUseConnectionCtx bool) *UserModel {
	u.ClientMetadataRequestUseConnectionCtx = tfconfig.BoolVariable(clientMetadataRequestUseConnectionCtx)
	return u
}

func (u *UserModel) WithClientPrefetchThreads(clientPrefetchThreads int) *UserModel {
	u.ClientPrefetchThreads = tfconfig.IntegerVariable(clientPrefetchThreads)
	return u
}

func (u *UserModel) WithClientResultChunkSize(clientResultChunkSize int) *UserModel {
	u.ClientResultChunkSize = tfconfig.IntegerVariable(clientResultChunkSize)
	return u
}

func (u *UserModel) WithClientResultColumnCaseInsensitive(clientResultColumnCaseInsensitive bool) *UserModel {
	u.ClientResultColumnCaseInsensitive = tfconfig.BoolVariable(clientResultColumnCaseInsensitive)
	return u
}

func (u *UserModel) WithClientSessionKeepAlive(clientSessionKeepAlive bool) *UserModel {
	u.ClientSessionKeepAlive = tfconfig.BoolVariable(clientSessionKeepAlive)
	return u
}

func (u *UserModel) WithClientSessionKeepAliveHeartbeatFrequency(clientSessionKeepAliveHeartbeatFrequency int) *UserModel {
	u.ClientSessionKeepAliveHeartbeatFrequency = tfconfig.IntegerVariable(clientSessionKeepAliveHeartbeatFrequency)
	return u
}

func (u *UserModel) WithClientTimestampTypeMapping(clientTimestampTypeMapping string) *UserModel {
	u.ClientTimestampTypeMapping = tfconfig.StringVariable(clientTimestampTypeMapping)
	return u
}

func (u *UserModel) WithComment(comment string) *UserModel {
	u.Comment = tfconfig.StringVariable(comment)
	return u
}

func (u *UserModel) WithDateInputFormat(dateInputFormat string) *UserModel {
	u.DateInputFormat = tfconfig.StringVariable(dateInputFormat)
	return u
}

func (u *UserModel) WithDateOutputFormat(dateOutputFormat string) *UserModel {
	u.DateOutputFormat = tfconfig.StringVariable(dateOutputFormat)
	return u
}

func (u *UserModel) WithDaysToExpiry(daysToExpiry int) *UserModel {
	u.DaysToExpiry = tfconfig.IntegerVariable(daysToExpiry)
	return u
}

func (u *UserModel) WithDefaultNamespace(defaultNamespace string) *UserModel {
	u.DefaultNamespace = tfconfig.StringVariable(defaultNamespace)
	return u
}

func (u *UserModel) WithDefaultRole(defaultRole string) *UserModel {
	u.DefaultRole = tfconfig.StringVariable(defaultRole)
	return u
}

func (u *UserModel) WithDefaultSecondaryRolesOption(defaultSecondaryRolesOption string) *UserModel {
	u.DefaultSecondaryRolesOption = tfconfig.StringVariable(defaultSecondaryRolesOption)
	return u
}

func (u *UserModel) WithDefaultWarehouse(defaultWarehouse string) *UserModel {
	u.DefaultWarehouse = tfconfig.StringVariable(defaultWarehouse)
	return u
}

func (u *UserModel) WithDisableMfa(disableMfa string) *UserModel {
	u.DisableMfa = tfconfig.StringVariable(disableMfa)
	return u
}

func (u *UserModel) WithDisabled(disabled string) *UserModel {
	u.Disabled = tfconfig.StringVariable(disabled)
	return u
}

func (u *UserModel) WithDisplayName(displayName string) *UserModel {
	u.DisplayName = tfconfig.StringVariable(displayName)
	return u
}

func (u *UserModel) WithEmail(email string) *UserModel {
	u.Email = tfconfig.StringVariable(email)
	return u
}

func (u *UserModel) WithEnableUnloadPhysicalTypeOptimization(enableUnloadPhysicalTypeOptimization bool) *UserModel {
	u.EnableUnloadPhysicalTypeOptimization = tfconfig.BoolVariable(enableUnloadPhysicalTypeOptimization)
	return u
}

func (u *UserModel) WithEnableUnredactedQuerySyntaxError(enableUnredactedQuerySyntaxError bool) *UserModel {
	u.EnableUnredactedQuerySyntaxError = tfconfig.BoolVariable(enableUnredactedQuerySyntaxError)
	return u
}

func (u *UserModel) WithErrorOnNondeterministicMerge(errorOnNondeterministicMerge bool) *UserModel {
	u.ErrorOnNondeterministicMerge = tfconfig.BoolVariable(errorOnNondeterministicMerge)
	return u
}

func (u *UserModel) WithErrorOnNondeterministicUpdate(errorOnNondeterministicUpdate bool) *UserModel {
	u.ErrorOnNondeterministicUpdate = tfconfig.BoolVariable(errorOnNondeterministicUpdate)
	return u
}

func (u *UserModel) WithFirstName(firstName string) *UserModel {
	u.FirstName = tfconfig.StringVariable(firstName)
	return u
}

func (u *UserModel) WithFullyQualifiedName(fullyQualifiedName string) *UserModel {
	u.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return u
}

func (u *UserModel) WithGeographyOutputFormat(geographyOutputFormat string) *UserModel {
	u.GeographyOutputFormat = tfconfig.StringVariable(geographyOutputFormat)
	return u
}

func (u *UserModel) WithGeometryOutputFormat(geometryOutputFormat string) *UserModel {
	u.GeometryOutputFormat = tfconfig.StringVariable(geometryOutputFormat)
	return u
}

func (u *UserModel) WithJdbcTreatDecimalAsInt(jdbcTreatDecimalAsInt bool) *UserModel {
	u.JdbcTreatDecimalAsInt = tfconfig.BoolVariable(jdbcTreatDecimalAsInt)
	return u
}

func (u *UserModel) WithJdbcTreatTimestampNtzAsUtc(jdbcTreatTimestampNtzAsUtc bool) *UserModel {
	u.JdbcTreatTimestampNtzAsUtc = tfconfig.BoolVariable(jdbcTreatTimestampNtzAsUtc)
	return u
}

func (u *UserModel) WithJdbcUseSessionTimezone(jdbcUseSessionTimezone bool) *UserModel {
	u.JdbcUseSessionTimezone = tfconfig.BoolVariable(jdbcUseSessionTimezone)
	return u
}

func (u *UserModel) WithJsonIndent(jsonIndent int) *UserModel {
	u.JsonIndent = tfconfig.IntegerVariable(jsonIndent)
	return u
}

func (u *UserModel) WithLastName(lastName string) *UserModel {
	u.LastName = tfconfig.StringVariable(lastName)
	return u
}

func (u *UserModel) WithLockTimeout(lockTimeout int) *UserModel {
	u.LockTimeout = tfconfig.IntegerVariable(lockTimeout)
	return u
}

func (u *UserModel) WithLogLevel(logLevel string) *UserModel {
	u.LogLevel = tfconfig.StringVariable(logLevel)
	return u
}

func (u *UserModel) WithLoginName(loginName string) *UserModel {
	u.LoginName = tfconfig.StringVariable(loginName)
	return u
}

func (u *UserModel) WithMiddleName(middleName string) *UserModel {
	u.MiddleName = tfconfig.StringVariable(middleName)
	return u
}

func (u *UserModel) WithMinsToBypassMfa(minsToBypassMfa int) *UserModel {
	u.MinsToBypassMfa = tfconfig.IntegerVariable(minsToBypassMfa)
	return u
}

func (u *UserModel) WithMinsToUnlock(minsToUnlock int) *UserModel {
	u.MinsToUnlock = tfconfig.IntegerVariable(minsToUnlock)
	return u
}

func (u *UserModel) WithMultiStatementCount(multiStatementCount int) *UserModel {
	u.MultiStatementCount = tfconfig.IntegerVariable(multiStatementCount)
	return u
}

func (u *UserModel) WithMustChangePassword(mustChangePassword string) *UserModel {
	u.MustChangePassword = tfconfig.StringVariable(mustChangePassword)
	return u
}

func (u *UserModel) WithNetworkPolicy(networkPolicy string) *UserModel {
	u.NetworkPolicy = tfconfig.StringVariable(networkPolicy)
	return u
}

func (u *UserModel) WithNoorderSequenceAsDefault(noorderSequenceAsDefault bool) *UserModel {
	u.NoorderSequenceAsDefault = tfconfig.BoolVariable(noorderSequenceAsDefault)
	return u
}

func (u *UserModel) WithOdbcTreatDecimalAsInt(odbcTreatDecimalAsInt bool) *UserModel {
	u.OdbcTreatDecimalAsInt = tfconfig.BoolVariable(odbcTreatDecimalAsInt)
	return u
}

func (u *UserModel) WithPassword(password string) *UserModel {
	u.Password = tfconfig.StringVariable(password)
	return u
}

func (u *UserModel) WithPreventUnloadToInternalStages(preventUnloadToInternalStages bool) *UserModel {
	u.PreventUnloadToInternalStages = tfconfig.BoolVariable(preventUnloadToInternalStages)
	return u
}

func (u *UserModel) WithQueryTag(queryTag string) *UserModel {
	u.QueryTag = tfconfig.StringVariable(queryTag)
	return u
}

func (u *UserModel) WithQuotedIdentifiersIgnoreCase(quotedIdentifiersIgnoreCase bool) *UserModel {
	u.QuotedIdentifiersIgnoreCase = tfconfig.BoolVariable(quotedIdentifiersIgnoreCase)
	return u
}

func (u *UserModel) WithRowsPerResultset(rowsPerResultset int) *UserModel {
	u.RowsPerResultset = tfconfig.IntegerVariable(rowsPerResultset)
	return u
}

func (u *UserModel) WithRsaPublicKey(rsaPublicKey string) *UserModel {
	u.RsaPublicKey = config.MultilineWrapperVariable(rsaPublicKey)
	return u
}

func (u *UserModel) WithRsaPublicKey2(rsaPublicKey2 string) *UserModel {
	u.RsaPublicKey2 = config.MultilineWrapperVariable(rsaPublicKey2)
	return u
}

func (u *UserModel) WithS3StageVpceDnsName(s3StageVpceDnsName string) *UserModel {
	u.S3StageVpceDnsName = tfconfig.StringVariable(s3StageVpceDnsName)
	return u
}

func (u *UserModel) WithSearchPath(searchPath string) *UserModel {
	u.SearchPath = tfconfig.StringVariable(searchPath)
	return u
}

func (u *UserModel) WithSimulatedDataSharingConsumer(simulatedDataSharingConsumer string) *UserModel {
	u.SimulatedDataSharingConsumer = tfconfig.StringVariable(simulatedDataSharingConsumer)
	return u
}

func (u *UserModel) WithStatementQueuedTimeoutInSeconds(statementQueuedTimeoutInSeconds int) *UserModel {
	u.StatementQueuedTimeoutInSeconds = tfconfig.IntegerVariable(statementQueuedTimeoutInSeconds)
	return u
}

func (u *UserModel) WithStatementTimeoutInSeconds(statementTimeoutInSeconds int) *UserModel {
	u.StatementTimeoutInSeconds = tfconfig.IntegerVariable(statementTimeoutInSeconds)
	return u
}

func (u *UserModel) WithStrictJsonOutput(strictJsonOutput bool) *UserModel {
	u.StrictJsonOutput = tfconfig.BoolVariable(strictJsonOutput)
	return u
}

func (u *UserModel) WithTimeInputFormat(timeInputFormat string) *UserModel {
	u.TimeInputFormat = tfconfig.StringVariable(timeInputFormat)
	return u
}

func (u *UserModel) WithTimeOutputFormat(timeOutputFormat string) *UserModel {
	u.TimeOutputFormat = tfconfig.StringVariable(timeOutputFormat)
	return u
}

func (u *UserModel) WithTimestampDayIsAlways24h(timestampDayIsAlways24h bool) *UserModel {
	u.TimestampDayIsAlways24h = tfconfig.BoolVariable(timestampDayIsAlways24h)
	return u
}

func (u *UserModel) WithTimestampInputFormat(timestampInputFormat string) *UserModel {
	u.TimestampInputFormat = tfconfig.StringVariable(timestampInputFormat)
	return u
}

func (u *UserModel) WithTimestampLtzOutputFormat(timestampLtzOutputFormat string) *UserModel {
	u.TimestampLtzOutputFormat = tfconfig.StringVariable(timestampLtzOutputFormat)
	return u
}

func (u *UserModel) WithTimestampNtzOutputFormat(timestampNtzOutputFormat string) *UserModel {
	u.TimestampNtzOutputFormat = tfconfig.StringVariable(timestampNtzOutputFormat)
	return u
}

func (u *UserModel) WithTimestampOutputFormat(timestampOutputFormat string) *UserModel {
	u.TimestampOutputFormat = tfconfig.StringVariable(timestampOutputFormat)
	return u
}

func (u *UserModel) WithTimestampTypeMapping(timestampTypeMapping string) *UserModel {
	u.TimestampTypeMapping = tfconfig.StringVariable(timestampTypeMapping)
	return u
}

func (u *UserModel) WithTimestampTzOutputFormat(timestampTzOutputFormat string) *UserModel {
	u.TimestampTzOutputFormat = tfconfig.StringVariable(timestampTzOutputFormat)
	return u
}

func (u *UserModel) WithTimezone(timezone string) *UserModel {
	u.Timezone = tfconfig.StringVariable(timezone)
	return u
}

func (u *UserModel) WithTraceLevel(traceLevel string) *UserModel {
	u.TraceLevel = tfconfig.StringVariable(traceLevel)
	return u
}

func (u *UserModel) WithTransactionAbortOnError(transactionAbortOnError bool) *UserModel {
	u.TransactionAbortOnError = tfconfig.BoolVariable(transactionAbortOnError)
	return u
}

func (u *UserModel) WithTransactionDefaultIsolationLevel(transactionDefaultIsolationLevel string) *UserModel {
	u.TransactionDefaultIsolationLevel = tfconfig.StringVariable(transactionDefaultIsolationLevel)
	return u
}

func (u *UserModel) WithTwoDigitCenturyStart(twoDigitCenturyStart int) *UserModel {
	u.TwoDigitCenturyStart = tfconfig.IntegerVariable(twoDigitCenturyStart)
	return u
}

func (u *UserModel) WithUnsupportedDdlAction(unsupportedDdlAction string) *UserModel {
	u.UnsupportedDdlAction = tfconfig.StringVariable(unsupportedDdlAction)
	return u
}

func (u *UserModel) WithUseCachedResult(useCachedResult bool) *UserModel {
	u.UseCachedResult = tfconfig.BoolVariable(useCachedResult)
	return u
}

func (u *UserModel) WithUserType(userType string) *UserModel {
	u.UserType = tfconfig.StringVariable(userType)
	return u
}

func (u *UserModel) WithWeekOfYearPolicy(weekOfYearPolicy int) *UserModel {
	u.WeekOfYearPolicy = tfconfig.IntegerVariable(weekOfYearPolicy)
	return u
}

func (u *UserModel) WithWeekStart(weekStart int) *UserModel {
	u.WeekStart = tfconfig.IntegerVariable(weekStart)
	return u
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (u *UserModel) WithNameValue(value tfconfig.Variable) *UserModel {
	u.Name = value
	return u
}

func (u *UserModel) WithAbortDetachedQueryValue(value tfconfig.Variable) *UserModel {
	u.AbortDetachedQuery = value
	return u
}

func (u *UserModel) WithAutocommitValue(value tfconfig.Variable) *UserModel {
	u.Autocommit = value
	return u
}

func (u *UserModel) WithBinaryInputFormatValue(value tfconfig.Variable) *UserModel {
	u.BinaryInputFormat = value
	return u
}

func (u *UserModel) WithBinaryOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.BinaryOutputFormat = value
	return u
}

func (u *UserModel) WithClientMemoryLimitValue(value tfconfig.Variable) *UserModel {
	u.ClientMemoryLimit = value
	return u
}

func (u *UserModel) WithClientMetadataRequestUseConnectionCtxValue(value tfconfig.Variable) *UserModel {
	u.ClientMetadataRequestUseConnectionCtx = value
	return u
}

func (u *UserModel) WithClientPrefetchThreadsValue(value tfconfig.Variable) *UserModel {
	u.ClientPrefetchThreads = value
	return u
}

func (u *UserModel) WithClientResultChunkSizeValue(value tfconfig.Variable) *UserModel {
	u.ClientResultChunkSize = value
	return u
}

func (u *UserModel) WithClientResultColumnCaseInsensitiveValue(value tfconfig.Variable) *UserModel {
	u.ClientResultColumnCaseInsensitive = value
	return u
}

func (u *UserModel) WithClientSessionKeepAliveValue(value tfconfig.Variable) *UserModel {
	u.ClientSessionKeepAlive = value
	return u
}

func (u *UserModel) WithClientSessionKeepAliveHeartbeatFrequencyValue(value tfconfig.Variable) *UserModel {
	u.ClientSessionKeepAliveHeartbeatFrequency = value
	return u
}

func (u *UserModel) WithClientTimestampTypeMappingValue(value tfconfig.Variable) *UserModel {
	u.ClientTimestampTypeMapping = value
	return u
}

func (u *UserModel) WithCommentValue(value tfconfig.Variable) *UserModel {
	u.Comment = value
	return u
}

func (u *UserModel) WithDateInputFormatValue(value tfconfig.Variable) *UserModel {
	u.DateInputFormat = value
	return u
}

func (u *UserModel) WithDateOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.DateOutputFormat = value
	return u
}

func (u *UserModel) WithDaysToExpiryValue(value tfconfig.Variable) *UserModel {
	u.DaysToExpiry = value
	return u
}

func (u *UserModel) WithDefaultNamespaceValue(value tfconfig.Variable) *UserModel {
	u.DefaultNamespace = value
	return u
}

func (u *UserModel) WithDefaultRoleValue(value tfconfig.Variable) *UserModel {
	u.DefaultRole = value
	return u
}

func (u *UserModel) WithDefaultSecondaryRolesOptionValue(value tfconfig.Variable) *UserModel {
	u.DefaultSecondaryRolesOption = value
	return u
}

func (u *UserModel) WithDefaultWarehouseValue(value tfconfig.Variable) *UserModel {
	u.DefaultWarehouse = value
	return u
}

func (u *UserModel) WithDisableMfaValue(value tfconfig.Variable) *UserModel {
	u.DisableMfa = value
	return u
}

func (u *UserModel) WithDisabledValue(value tfconfig.Variable) *UserModel {
	u.Disabled = value
	return u
}

func (u *UserModel) WithDisplayNameValue(value tfconfig.Variable) *UserModel {
	u.DisplayName = value
	return u
}

func (u *UserModel) WithEmailValue(value tfconfig.Variable) *UserModel {
	u.Email = value
	return u
}

func (u *UserModel) WithEnableUnloadPhysicalTypeOptimizationValue(value tfconfig.Variable) *UserModel {
	u.EnableUnloadPhysicalTypeOptimization = value
	return u
}

func (u *UserModel) WithEnableUnredactedQuerySyntaxErrorValue(value tfconfig.Variable) *UserModel {
	u.EnableUnredactedQuerySyntaxError = value
	return u
}

func (u *UserModel) WithErrorOnNondeterministicMergeValue(value tfconfig.Variable) *UserModel {
	u.ErrorOnNondeterministicMerge = value
	return u
}

func (u *UserModel) WithErrorOnNondeterministicUpdateValue(value tfconfig.Variable) *UserModel {
	u.ErrorOnNondeterministicUpdate = value
	return u
}

func (u *UserModel) WithFirstNameValue(value tfconfig.Variable) *UserModel {
	u.FirstName = value
	return u
}

func (u *UserModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *UserModel {
	u.FullyQualifiedName = value
	return u
}

func (u *UserModel) WithGeographyOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.GeographyOutputFormat = value
	return u
}

func (u *UserModel) WithGeometryOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.GeometryOutputFormat = value
	return u
}

func (u *UserModel) WithJdbcTreatDecimalAsIntValue(value tfconfig.Variable) *UserModel {
	u.JdbcTreatDecimalAsInt = value
	return u
}

func (u *UserModel) WithJdbcTreatTimestampNtzAsUtcValue(value tfconfig.Variable) *UserModel {
	u.JdbcTreatTimestampNtzAsUtc = value
	return u
}

func (u *UserModel) WithJdbcUseSessionTimezoneValue(value tfconfig.Variable) *UserModel {
	u.JdbcUseSessionTimezone = value
	return u
}

func (u *UserModel) WithJsonIndentValue(value tfconfig.Variable) *UserModel {
	u.JsonIndent = value
	return u
}

func (u *UserModel) WithLastNameValue(value tfconfig.Variable) *UserModel {
	u.LastName = value
	return u
}

func (u *UserModel) WithLockTimeoutValue(value tfconfig.Variable) *UserModel {
	u.LockTimeout = value
	return u
}

func (u *UserModel) WithLogLevelValue(value tfconfig.Variable) *UserModel {
	u.LogLevel = value
	return u
}

func (u *UserModel) WithLoginNameValue(value tfconfig.Variable) *UserModel {
	u.LoginName = value
	return u
}

func (u *UserModel) WithMiddleNameValue(value tfconfig.Variable) *UserModel {
	u.MiddleName = value
	return u
}

func (u *UserModel) WithMinsToBypassMfaValue(value tfconfig.Variable) *UserModel {
	u.MinsToBypassMfa = value
	return u
}

func (u *UserModel) WithMinsToUnlockValue(value tfconfig.Variable) *UserModel {
	u.MinsToUnlock = value
	return u
}

func (u *UserModel) WithMultiStatementCountValue(value tfconfig.Variable) *UserModel {
	u.MultiStatementCount = value
	return u
}

func (u *UserModel) WithMustChangePasswordValue(value tfconfig.Variable) *UserModel {
	u.MustChangePassword = value
	return u
}

func (u *UserModel) WithNetworkPolicyValue(value tfconfig.Variable) *UserModel {
	u.NetworkPolicy = value
	return u
}

func (u *UserModel) WithNoorderSequenceAsDefaultValue(value tfconfig.Variable) *UserModel {
	u.NoorderSequenceAsDefault = value
	return u
}

func (u *UserModel) WithOdbcTreatDecimalAsIntValue(value tfconfig.Variable) *UserModel {
	u.OdbcTreatDecimalAsInt = value
	return u
}

func (u *UserModel) WithPasswordValue(value tfconfig.Variable) *UserModel {
	u.Password = value
	return u
}

func (u *UserModel) WithPreventUnloadToInternalStagesValue(value tfconfig.Variable) *UserModel {
	u.PreventUnloadToInternalStages = value
	return u
}

func (u *UserModel) WithQueryTagValue(value tfconfig.Variable) *UserModel {
	u.QueryTag = value
	return u
}

func (u *UserModel) WithQuotedIdentifiersIgnoreCaseValue(value tfconfig.Variable) *UserModel {
	u.QuotedIdentifiersIgnoreCase = value
	return u
}

func (u *UserModel) WithRowsPerResultsetValue(value tfconfig.Variable) *UserModel {
	u.RowsPerResultset = value
	return u
}

func (u *UserModel) WithRsaPublicKeyValue(value tfconfig.Variable) *UserModel {
	u.RsaPublicKey = value
	return u
}

func (u *UserModel) WithRsaPublicKey2Value(value tfconfig.Variable) *UserModel {
	u.RsaPublicKey2 = value
	return u
}

func (u *UserModel) WithS3StageVpceDnsNameValue(value tfconfig.Variable) *UserModel {
	u.S3StageVpceDnsName = value
	return u
}

func (u *UserModel) WithSearchPathValue(value tfconfig.Variable) *UserModel {
	u.SearchPath = value
	return u
}

func (u *UserModel) WithSimulatedDataSharingConsumerValue(value tfconfig.Variable) *UserModel {
	u.SimulatedDataSharingConsumer = value
	return u
}

func (u *UserModel) WithStatementQueuedTimeoutInSecondsValue(value tfconfig.Variable) *UserModel {
	u.StatementQueuedTimeoutInSeconds = value
	return u
}

func (u *UserModel) WithStatementTimeoutInSecondsValue(value tfconfig.Variable) *UserModel {
	u.StatementTimeoutInSeconds = value
	return u
}

func (u *UserModel) WithStrictJsonOutputValue(value tfconfig.Variable) *UserModel {
	u.StrictJsonOutput = value
	return u
}

func (u *UserModel) WithTimeInputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimeInputFormat = value
	return u
}

func (u *UserModel) WithTimeOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimeOutputFormat = value
	return u
}

func (u *UserModel) WithTimestampDayIsAlways24hValue(value tfconfig.Variable) *UserModel {
	u.TimestampDayIsAlways24h = value
	return u
}

func (u *UserModel) WithTimestampInputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimestampInputFormat = value
	return u
}

func (u *UserModel) WithTimestampLtzOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimestampLtzOutputFormat = value
	return u
}

func (u *UserModel) WithTimestampNtzOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimestampNtzOutputFormat = value
	return u
}

func (u *UserModel) WithTimestampOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimestampOutputFormat = value
	return u
}

func (u *UserModel) WithTimestampTypeMappingValue(value tfconfig.Variable) *UserModel {
	u.TimestampTypeMapping = value
	return u
}

func (u *UserModel) WithTimestampTzOutputFormatValue(value tfconfig.Variable) *UserModel {
	u.TimestampTzOutputFormat = value
	return u
}

func (u *UserModel) WithTimezoneValue(value tfconfig.Variable) *UserModel {
	u.Timezone = value
	return u
}

func (u *UserModel) WithTraceLevelValue(value tfconfig.Variable) *UserModel {
	u.TraceLevel = value
	return u
}

func (u *UserModel) WithTransactionAbortOnErrorValue(value tfconfig.Variable) *UserModel {
	u.TransactionAbortOnError = value
	return u
}

func (u *UserModel) WithTransactionDefaultIsolationLevelValue(value tfconfig.Variable) *UserModel {
	u.TransactionDefaultIsolationLevel = value
	return u
}

func (u *UserModel) WithTwoDigitCenturyStartValue(value tfconfig.Variable) *UserModel {
	u.TwoDigitCenturyStart = value
	return u
}

func (u *UserModel) WithUnsupportedDdlActionValue(value tfconfig.Variable) *UserModel {
	u.UnsupportedDdlAction = value
	return u
}

func (u *UserModel) WithUseCachedResultValue(value tfconfig.Variable) *UserModel {
	u.UseCachedResult = value
	return u
}

func (u *UserModel) WithUserTypeValue(value tfconfig.Variable) *UserModel {
	u.UserType = value
	return u
}

func (u *UserModel) WithWeekOfYearPolicyValue(value tfconfig.Variable) *UserModel {
	u.WeekOfYearPolicy = value
	return u
}

func (u *UserModel) WithWeekStartValue(value tfconfig.Variable) *UserModel {
	u.WeekStart = value
	return u
}
