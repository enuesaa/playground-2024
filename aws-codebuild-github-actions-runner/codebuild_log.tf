resource "aws_cloudwatch_log_group" "codebuild" {
  name = "${var.identifier}-codebuild"
}
