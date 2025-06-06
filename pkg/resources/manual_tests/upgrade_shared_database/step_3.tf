# Commands to run
# - terraform import snowflake_shared_database.from_share '"test_database"' (import shared database into state)
# - terraform plan (expect empty plan)

terraform {
  required_providers {
    snowflake = {
      source  = "snowflakedb/snowflake"
      version = ">= 0.92.0" # latest
    }
  }
}

provider "snowflake" {}

provider "snowflake" {
  profile = "secondary_test_account"
  alias   = second_account
}

resource "snowflake_share" "test" {
  provider = snowflake.second_account
  name     = "test_share"
  accounts = ["<primary_account_organization_name>.<primary_account_account_name>"] # TODO: Replace
}

resource "snowflake_database" "test" {
  provider = snowflake.second_account
  name     = "test_database"
}

resource "snowflake_grant_privileges_to_share" "test" {
  provider    = snowflake.second_account
  privileges  = ["USAGE"]
  on_database = snowflake_database.test.name
  to_share    = snowflake_share.test.name
}

# Changed old snowflake_database to the new snowflake_shared_database counter-part
resource "snowflake_shared_database" "from_share" {
  depends_on = [snowflake_grant_privileges_to_share.test]
  name       = snowflake_database.test.name
  from_share = "\"<second_account_organization_name>\".\"<second_account_account_name>\".\"${snowflake_share.test.name}\"" # TODO: Replace
}
