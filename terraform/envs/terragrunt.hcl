remote_state {
  backend = "gcs"
  generate = {
    path      = "_backend.tf"
    if_exists = "overwrite"
  }
  config = {
    project  = get_env("GOOGLE_CLOUD_PROJECT_ID")
    location = "ASIA-NORTHEAST1"
    bucket   = get_env("REMOTE_STETE_BUCKET")
    prefix   = "${path_relative_to_include()}"
  }
}

generate "provider" {
  path      = "_provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.56.0"
    }
  }
}

provider "google" {
  project = "${get_env("GOOGLE_CLOUD_PROJECT_ID")}"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
EOF
}

generate "tfvars" {
  path      = "terraform.tfvars"
  if_exists = "overwrite"
  contents  = <<EOT
google_cloud_project_id = "${get_env("GOOGLE_CLOUD_PROJECT_ID")}"
EOT
}
