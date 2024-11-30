resource "aws_appsync_api_key" "main" {
  api_id  = aws_appsync_graphql_api.main.id
  expires = "2024-12-30T00:00:00Z"
}
