target "docker-metadata-action" {}

target "bake-platform" {
  inherits = ["docker-metadata-action"]
  context = "ubuntu"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
  ]
}
