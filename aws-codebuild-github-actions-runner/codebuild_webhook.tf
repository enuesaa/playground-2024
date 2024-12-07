resource "aws_codebuild_webhook" "main" {
  project_name = aws_codebuild_project.main.name

  build_type   = "BUILD"

  filter_group {
    filter {
      type    = "EVENT"
      pattern = "WORKFLOW_JOB_QUEUED"
    }
  }
}
