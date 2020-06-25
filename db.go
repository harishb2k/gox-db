package db

const (
    DatabaseErrorRecordNotFound = "record_not_found"
    DatabaseErrorUnknown        = "unknown_db_error"
)

type IDb interface {
    InitDatabase() (err error)
    Persist(queryString string, val ...interface{}) (e error)
    FindOne(queryString string, mapper RowMapper, val ...interface{}) (result interface{}, e error)
    FindAll(queryString string, mapper RowMapper, val ...interface{}) (result []interface{}, e error)
    Execute(queryString string, val ...interface{}) (result interface{}, e error)
    ShopDatabase() (err error)
}

type RowMapper interface {
    Map(row interface{}) (result interface{}, err error)
}
