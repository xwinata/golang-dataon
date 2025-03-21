package databases

import (
	"database/sql"
	"fmt"
	"golang-dataon/configs"
	"golang-dataon/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/joho/godotenv/autoload"
)

func NewGormDB(config *configs.Config) *gorm.DB {
	gormConfig := &gorm.Config{}

	// https://gorm.io/docs/logger.html#Log-Levels
	switch config.DB_LOGMODE {
	case "silent":
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	case "info":
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	case "warn":
		gormConfig.Logger = logger.Default.LogMode(logger.Warn)
	case "error":
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	default:
		break
	}

	pg, _ := gorm.Open(postgres.Open(
		getPGconfig(
			config.DB_HOST,
			config.DB_USER,
			config.DB_PASS,
			config.DB_NAME,
			config.DB_PORT,
			config.DB_SSL,
			config.DB_TZ,
		)), gormConfig,
	)

	if config.DB_AUTOMIGRATE == "enable" {
		models := []interface{}{
			&models.Hierarchy{},
		}

		pg.AutoMigrate(models...)
	}
	return pg
}

func getPGconfig(dbHost string, dbUser string, dbPass string, dbName string, dbPort string, dbSSL string, dbTZ string) string {
	configString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	dbSSLmode := "disable"
	if dbSSL == "enable" {
		dbSSLmode = "enable"
	}

	configString += fmt.Sprintf(" sslmode=%s", dbSSLmode)

	dbTimezone := "UTC"
	if len(dbTZ) > 0 && dbTZ != "UTC" {
		dbTimezone = dbTZ
	}

	configString += fmt.Sprintf(" TimeZone=%s", dbTimezone)

	return configString
}

type GormDBMock struct {
	DB      *gorm.DB
	SqlDb   *sql.DB
	SqlMock sqlmock.Sqlmock
}

func NewGormDBMock() *GormDBMock {
	dbMock := &GormDBMock{}
	dbMock.Setup()
	return dbMock
}

func (d *GormDBMock) Setup() {
	d.SqlDb, d.SqlMock, _ = sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn:       d.SqlDb,
		DriverName: "postgres",
	})

	d.DB, _ = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func (d *GormDBMock) Close() {
	d.SqlDb.Close()
}
