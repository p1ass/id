resource "google_cloud_run_v2_service" "backend" {
  name     = "${local.env}-backend"
  location = local.location
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      name  = "${local.env}-backend"
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
    scaling {
      min_instance_count = 0
      max_instance_count = 1
    }
    timeout = "5s"
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }

  depends_on = []
}
