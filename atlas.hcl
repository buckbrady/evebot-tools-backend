env "local" {
  src = "file://schema.hcl"
  url = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  dev = "docker://postgres/16/dev"
}

env "prod" {

}