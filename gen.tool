version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn : "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Etc/UTC"
  # input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
  db  : "postgres"
  # enter the required data table or leave it blank.You can input : orders,users,goods
  tables  :
  # specify a directory for output
  outPath :  "./pkg/database"
  # query code file name, default: gen.go
  outFile :  ""
  # generate unit test for query code
  withUnitTest  : false
  # generated model code's package name
  modelPkgName  : ""
  # generate with pointer when field is nullable
  fieldNullable : true
  # generate field with gorm index tag
  fieldWithIndexTag : false
  # generate field with gorm column type tag
  fieldWithTypeTag  : false