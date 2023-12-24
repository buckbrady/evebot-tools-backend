variable "pg_username" {
  type    = string
  default = "postgres"
}

variable "pg_password" {
  type    = string
  default = "postgres"
}

variable "pg_host" {
  type    = string
  default = "localhost"
}

variable "pg_port" {
  type    = number
  default = 5432
}

variable "pg_db" {
  type    = string
  default = "postgres"
}

lint {
  incompatible {
    error = true
  }
  naming {
    error   = true
    match   = "^[a-z]+$"
    message = "must be lowercase"
  }
  data_depend {
    error = true
  }
}

env "dev" {
  src = "file://schema.hcl"
  url = "postgres://${var.pg_username}:${var.pg_password}@${var.pg_host}:${var.pg_port}/${var.pg_db}?sslmode=disable"
  dev = "docker://postgres/16/dev"
}

env "prod" {
  src = "file://schema.hcl"
  url = "postgres://${var.pg_username}:${var.pg_password}@${var.pg_host}:${var.pg_port}/${var.pg_db}?sslmode=disable"
  dev = "docker://postgres/16/dev"
}