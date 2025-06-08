package simple

type Database struct {
	Name string
}

type DatabaseMySQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{Name: "MySQL"})
}

type DatabaseRepository struct {
	DatabaseMySQL *DatabaseMySQL
	DatabaseMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(databaseMySQL *DatabaseMySQL, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMySQL: databaseMySQL,
		DatabaseMongoDB: databaseMongoDB,
	}
}