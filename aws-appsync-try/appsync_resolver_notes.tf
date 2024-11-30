resource "aws_appsync_resolver" "notes" {
  api_id = aws_appsync_graphql_api.main.id
  field  = "notes"
  type   = "Query"
  kind   = "UNIT"
  data_source = aws_appsync_datasource.none.name

  code   = file("${path.module}/appsync_resolver_notes.js")

  runtime {
    name            = "APPSYNC_JS"
    runtime_version = "1.0.0"
  }
}
