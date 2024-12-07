resource "aws_codebuild_project" "main" {
  name         = "${var.identifier}"

  service_role = aws_iam_role.codebuild.arn

  source {
    type     = "GITHUB"
    location = "https://github.com/enuesaa/<repository-name-here>"
  }

  artifacts {
    type = "NO_ARTIFACTS"
  }

  environment {
    type            = "LINUX_CONTAINER"
    compute_type    = "BUILD_GENERAL1_SMALL"
    image           = "aws/codebuild/standard:7.0"
  }

  logs_config {
    cloudwatch_logs {
      group_name = aws_cloudwatch_log_group.codebuild.name
    }
  }
}
