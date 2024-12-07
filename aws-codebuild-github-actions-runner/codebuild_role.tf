resource "aws_iam_role" "codebuild" {
  name = "${var.identifier}-codebuild"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "codebuild.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "codebuild_logs" {
  name = "logs"
  role = aws_iam_role.codebuild.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

// CodeBuild が CodeConnection 経由でいろいろするため、この権限が必要
// GitHub Repository に Webhook の設定をしたり、Webhook Event を受け取ったり..
resource "aws_iam_role_policy" "codebuild_codeconnections" {
  name = "codeconnections"
  role = aws_iam_role.codebuild.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "codeconnections:GetConnectionToken",
          "codeconnections:GetConnection"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}
