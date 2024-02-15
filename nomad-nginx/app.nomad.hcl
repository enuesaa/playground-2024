job "app" {
  datacenters = ["dc1"]

  group "web" {
    network {
      port "http" {
        to = 80
        static = 3000
      }
    }

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx:latest"
        ports = ["http"]
      }
    }
  }
}
