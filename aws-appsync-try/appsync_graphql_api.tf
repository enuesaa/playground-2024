resource "aws_appsync_graphql_api" "main" {
  authentication_type = "API_KEY"
  name                = var.identifier

  schema = file("${path.module}/schema.graphql")
}
