include "root" {
  path = find_in_parent_folders()
}

dependency "shared" {
  config_path = "../shared"

  # Configure mock outputs for the `validate` command that are returned when there are no outputs available
  # (e.g the module hasn't been applied yet.
  mock_outputs_allowed_terraform_commands = ["validate"]
  mock_outputs = {
    container_images_repository_id = "fake_repository_id"
  }
}

inputs = {
  container_images_repository_id = dependency.shared.outputs.container_images_repository_id
}
