variable "google_cloud_project_id" {
  type        = string
  description = "Google Cloud Project ID"
}

variable "container_images_repository_id" {
  type        = string
  description = "Artifact Repository id for container images"
}

locals {
  env      = "dev"
  location = "asia-northeast1"
}
