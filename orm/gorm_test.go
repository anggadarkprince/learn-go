package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// https://gorm.io/docs/connecting_to_the_database.html
func OpenConnection() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

var db, _ = OpenConnection()

func TestOpenConnection(t *testing.T) {
	if db == nil {
		t.Error("Failed to open database connection")
	} else {
		t.Log("Database connection opened successfully")
	}
}

func TestExecuteSQL(t *testing.T) {
	// https://gorm.io/docs/create.html
	err := db.Exec("INSERT INTO samples(id, name) VALUES(?, ?)", 1, "test").Error
	assert.Nil(t, err, "Expected no error when executing SQL")

	// https://gorm.io/docs/update.html
	err = db.Exec("UPDATE samples SET name = ? WHERE id = ?", "updated", 1).Error
	assert.Nil(t, err, "Expected no error when updating SQL")

	// https://gorm.io/docs/delete.html
	//err = db.Exec("DELETE FROM samples WHERE id = ?", 1).Error
	//assert.Nil(t, err, "Expected no error when deleting SQL")
}

type Sample struct {
	ID string
	Name string
}
func TestQuerySQL(t *testing.T) {
	// https://gorm.io/docs/query.html
	var sample Sample
	err := db.Raw("SELECT id, name FROM samples WHERE id = ?", 1).Scan(&sample).Error
	assert.Nil(t, err, "Expected no error when querying SQL")
	assert.Equal(t, "test", sample.Name, "Expected name to be 'test'")

	var samples []Sample
	err = db.Raw("SELECT id, name FROM samples").Scan(&samples).Error
	assert.Nil(t, err, "Expected no error when querying all samples")
}

func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("SELECT id, name FROM samples").Rows()
	assert.Nil(t, err, "Expected no error when getting rows")
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		assert.Nil(t, err, "Expected no error when scanning row")

		samples = append(samples, Sample{ID: id, Name: name})
	}
	assert.NotEmpty(t, samples, "Expected samples to not be empty")
}

func TestScanRows(t *testing.T) {
	rows, err := db.Raw("SELECT id, name FROM samples").Rows()
	assert.Nil(t, err, "Expected no error when getting rows")
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var sample Sample
		err := db.ScanRows(rows, &sample)
		assert.Nil(t, err, "Expected no error when scanning row")
		samples = append(samples, sample)
	}
	assert.NotEmpty(t, samples, "Expected samples to not be empty")
}

func TestSqlRowQueryBuilder(t *testing.T) {
	var sample Sample
	err := db.Table("samples").Where("id = ?", 1).First(&sample).Error
	assert.Nil(t, err, "Expected no error when querying single row")
	assert.Equal(t, "test", sample.Name, "Expected name to be 'test'")

	var samples []Sample
	err = db.Table("samples").Find(&samples).Error
	assert.Nil(t, err, "Expected no error when querying all rows")
	assert.NotEmpty(t, samples, "Expected samples to not be empty")
}