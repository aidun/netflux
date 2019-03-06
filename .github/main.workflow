workflow "netflux Build" {
  on = "push"
  resolves = [
    "GitHub Push",
    "cedrickring/golang-action/go1.12@1.1.1",
  ]
}

action "GitHub Action for Docker" {
  uses = "actions/docker/cli@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  args = "build -t netflux ."
}

action "Docker Registry" {
  uses = "actions/docker/login@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  needs = [
    "GitHub Action for Docker",
    "cedrickring/golang-action/go1.12@1.1.1",
  ]
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Docker Tag" {
  uses = "actions/docker/tag@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  needs = ["Docker Registry"]
  args = "netflux aidun/netflux"
}

action "GitHub Push" {
  uses = "actions/docker/cli@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  needs = ["Docker Tag"]
  args = "push aidun/netflux"
}

action "cedrickring/golang-action/go1.12@1.1.1" {
  uses = "cedrickring/golang-action@1.1.1"
}
