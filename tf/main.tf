terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "5.0.1"
    }
  }
}

provider "heroku" {
  # source the heroku secrets file before running terraform
}

resource "heroku_app" "bingo" {
  name   = "gbingo"
  region = "us"
  stack  = "container"

  config_vars = {
    SECRET = var.token_secret
  }

  buildpacks = [
    "heroku/go"
  ]
}

resource "heroku_build" "bingo" {
  app_id = heroku_app.bingo.id

  source {
    path = "../"
  }
}

resource "heroku_formation" "bingo" {
  app_id     = heroku_app.bingo.id
  type       = "web"
  quantity   = 1
  size       = "Standard-1x"
  depends_on = [heroku_build.bingo]
}

resource "heroku_addon" "postgres" {
  app_id = heroku_app.bingo.id
  plan   = "heroku-postgresql:hobby-dev"
}

output "bingo_app_url" {
  value = "https://${heroku_app.bingo.name}.herokuapp.com"
}
