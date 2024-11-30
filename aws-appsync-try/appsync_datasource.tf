resource "aws_appsync_datasource" "none" {
  api_id = aws_appsync_graphql_api.main.id

  name   = "none"
  type   = "NONE"
}
