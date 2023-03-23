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
}

resource "google_cloud_run_domain_mapping" "default" {
  location = local.location
  name     = local.domain

  metadata {
    namespace = var.google_cloud_project_id
  }

  spec {
    route_name = google_cloud_run_v2_service.backend.name
  }
}

resource "google_service_account" "vercel" {
  account_id   = "${local.env}-vercel-service-account"
  display_name = "${local.env} Service Account for Vercel"
}

resource "google_cloud_run_v2_service_iam_binding" "backend" {
  location = google_cloud_run_v2_service.backend.location
  name  = google_cloud_run_v2_service.backend.name
  role     = "roles/run.invoker"
  members = [
    google_service_account.vercel.member
  ]
}
