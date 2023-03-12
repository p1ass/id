resource "google_artifact_registry_repository" "container_images" {
  location      = local.location
  repository_id = "container-images"
  description   = "Docker container images"
  format        = "DOCKER"
}
