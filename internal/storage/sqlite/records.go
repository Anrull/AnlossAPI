package sqlite

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"log"
	"time"

	"gorm.io/gorm"
)

type Records struct {
	gorm.Model
	Date     string `json:"date"`
	Name     string `json:"name"`
	Class    string `json:"classes"`
	Olimps   string `json:"olimps"`
	Stage    string `json:"stage"`
	Subjects string `json:"subjects"`
	Teachers string `json:"teachers"`
}

var RecordsDB *gorm.DB

func newRecords(recordsPath string) {
	RecordsDB, _ = gorm.Open(sqlite.Open(recordsPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err := RecordsDB.AutoMigrate(&Records{}); err != nil {
		log.Println("Не удалось создать БД")
	}
}

func AddRecord(name, class, olimp, sub, teacher, stage string, date ...string) error {
	var record Records
	result := RecordsDB.First(&record, "name = ? AND class = ? AND olimps = ? AND stage = ? AND subjects = ? AND teachers = ?", name, class, olimp, stage, sub, teacher)
	if result.Error == nil {
		return nil
	}

	newUser := Records{
		Date:     time.Now().Format("2006-01-02"),
		Name:     name,
		Class:    class,
		Olimps:   olimp,
		Stage:    stage,
		Subjects: sub,
		Teachers: teacher,
	}

	if len(date) > 0 {
		newUser.Date = date[0]
	}

	return RecordsDB.Create(&newUser).Error
}

func GetRecords(name, sub, olimp, stage, teacher string) (*[]Records, error) {
	var records []Records
	query := getQuery(name, sub, olimp, stage, teacher)

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}

func GetRecordsCount(name, sub, olimp, stage, teacher string) (int, error) {
	var count int64
	query := getQuery(name, sub, olimp, stage, teacher)

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func getQuery(name, sub, olimp, stage, teacher string) *gorm.DB {
	query := RecordsDB.Model(&Records{}).Where("name = ?", name)

	if sub != "nil" {
		query = query.Where("subjects LIKE ?", "%"+sub+"%")
	}
	if olimp != "nil" {
		query = query.Where("olimps LIKE ?", "%"+olimp+"%")
	}
	if stage != "nil" {
		query = query.Where("stage LIKE ?", "%"+stage+"%")
	}
	if teacher != "nil" {
		query = query.Where("teachers LIKE ?", "%"+teacher+"%")
	}

	return query
}

func GetAllRecords() (*[]Records, error) {
	var records []Records

	if err := RecordsDB.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}

func DeleteAllRecords(databasePath string) error {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	// SQL-запрос для удаления всех данных из таблицы
	sqlStatement := `DELETE FROM records`

	// Выполнение запроса
	_, err = db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}
