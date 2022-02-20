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

func SetUpDatabase() *gorm.DB {
	godotenv.Load()
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
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
	return db
}
