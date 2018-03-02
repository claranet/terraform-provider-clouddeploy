provider "ghost" {
  user = "test"
  password = "test"
  endpoint = "localhost"
}

resource "ghost_app" "my-app" {
  name = "test"
}
