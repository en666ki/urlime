table "urls" {
  schema = schema.public
  column "surl" {
    null = false
    type = text
  }
  column "url" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.surl]
  }
}

table "urls_e2e" {
  schema = schema.public
  column "surl" {
    null = false
    type = text
  }
  column "url" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.surl]
  }
}

schema "public" {
  comment = "Default public schema"
}