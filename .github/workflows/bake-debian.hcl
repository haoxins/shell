target "docker-metadata-action" {}

target "bake-platform" {
  inherits = ["docker-metadata-action"]
  context = "debian"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
  ]
}
