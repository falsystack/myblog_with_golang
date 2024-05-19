provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_iam_user" "terraform_user" {
  name = "terraform_user"
}

data "aws_iam_policy_document" "policy_document" {
  statement {
    effect    = "Allow"
    actions   = [
      "iam:CreateOpenIDConnectProvider",
      "iam:GetOpenIDConnectProvider",
      "iam:DeleteOpenIDConnectProvider",
      "iam:PutRolePolicy",
      "iam:CreateRole",
      "iam:GetRole"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_user_policy" "terraform_policy" {
  name   = "terraform_policy"
  user   = aws_iam_user.terraform_user.name
  policy = data.aws_iam_policy_document.policy_document.json
}

resource "aws_ecr_repository" "toy_project_repository" {
  name                 = "toy"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_iam_openid_connect_provider" "toy_project" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
}

// github action
data "aws_iam_policy_document" "github_action_policy_document" {
  # ECR ログインに必要
  statement {
    effect    = "Allow"
    actions   = ["ecr:GetAuthorizationToken"]
    resources = ["*"]
  }

  # `docker push` に必要
  statement {
    effect = "Allow"
    actions = [
      "ecr:CompleteLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:InitiateLayerUpload",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
    ]
    resources = [aws_ecr_repository.toy_project_repository.arn]
  }
}

resource "aws_iam_policy" "github_action_policy" {
  name   = "github_action_policy"
  policy = data.aws_iam_policy_document.github_action_policy_document.json
}

resource "aws_iam_role" "github_action_role" {
  name = "github_action_role"

  assume_role_policy = data.aws_iam_policy_document.github_action_role_policy.json
}

data "aws_iam_policy_document" "github_action_role_policy" {
  statement {
    effect = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.toy_project.arn]
    }
    condition {
#       test     = "StringEquals"
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.github_action_role.name
  policy_arn = aws_iam_policy.github_action_policy.arn
}


/*

resource "aws_ecr_repository" "toy_project_repository" {
  name                 = "toy"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_iam_openid_connect_provider" "toy_project" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
}

# IAM Role Assume Policy Document
data "aws_iam_policy_document" "github_action_assume_role_policy" {
  statement {
    effect = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.toy_project.arn]
    }
    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }
  }
}

# IAM Role Policy Document
data "aws_iam_policy_document" "github_action_role_policy_document" {
  statement {
    effect    = "Allow"
    actions   = ["ecr:GetAuthorizationToken"]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ecr:CompleteLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:InitiateLayerUpload",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      //
      "ecr:CreateRepository",
      "ecr:DescribeRepositories",
      "ecr:ListTagsForResource",
      "ecr:DeleteRepository",
      "iam:CreateOpenIDConnectProvider",
      "iam:GetOpenIDConnectProvider",
      "iam:DeleteOpenIDConnectProvider",
      "iam:PutRolePolicy",
      "iam:CreateRole",
      "iam:GetRole"
    ]
    resources = [aws_ecr_repository.toy_project_repository.arn]
  }
}

resource "aws_iam_role" "github_action_role" {
  name               = "github_action_role"
  assume_role_policy = data.aws_iam_policy_document.github_action_assume_role_policy.json
}

resource "aws_iam_role_policy" "github_action_role_policy" {
  name   = "allow-ecr-push-image"
  role   = aws_iam_role.github_action_role.name
  policy = data.aws_iam_policy_document.github_action_role_policy_document.json
}

resource "aws_iam_policy_attachment" "test-attach" {
  name       = "test-attachment"
  users      = ["terraform-user"]
  roles      = [aws_iam_role.github_action_role.name]
  policy_arn = aws_iam_policy.policy.arn
}

*/