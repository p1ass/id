resource "google_cloud_run_v2_service" "backend" {
  name     = "${local.env}-backend"
  location = local.location
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      name  = "${local.env}-backend"
      image = "${local.location}-docker.pkg.dev/${var.google_cloud_project_id}/${var.container_images_repository_id}/backend:latest"

      ports {
        name           = "h2c"
        container_port = 8080
      }
    }
    scaling {
      min_instance_count = 0
      max_instance_count = 1
    }
    timeout = "3s"
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }

  depends_on = []
}
