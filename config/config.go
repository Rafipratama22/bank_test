package config

import (
	"fmt"
	"os"
	"time"

	"github.com/Rafipratama22/bank_test/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetUpDatabase untuk menginisialisasi database
// Menggunakan GORM sebagaim ORM dan menggunakan driver postgres sebagai database
func SetUpDatabase() *gorm.DB {
	godotenv.Load()
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	dsn := "host=" + db_host + " user=" + db_user + " dbname=" + db_name + " password=" + db_pass + " port=" + db_port + " sslmode=disable"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.UserEntity{}, &entity.RekeningEntity{}, &entity.HistoryEntity{}, &entity.MerchantEntity{})
	DB = db
	return db
}

func GetDatabaseTest() *gorm.DB {
	db := SetUpDatabase()
	return db
}

func ClearTable() {
	DB.Exec("DELETE FROM user_entities")
	DB.Exec("DELETE FROM rekening_entities")
	DB.Exec("DELETE FROM history_entities")
	DB.Exec("DELETE FROM merchant_entities")
}


