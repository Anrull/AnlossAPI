package sqlite

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var FamilyDB *gorm.DB

type Students struct {
	gorm.Model
	Name  string
	Stage string
	Snils string
}

func newStudents(studentsPath string) {
	var err error

	FamilyDB, err = gorm.Open(sqlite.Open(studentsPath),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal("don`t exists database students")
	}
}

func CheckSnils(value string, hashing bool) (bool, string, string) {
	var student Students
	if hashing {
		value = hashValue(value)
	}

	result := FamilyDB.First(&student, "snils = ?", value)
	if result.Error != nil {
		return false, "", ""
	}
	return true, student.Name, student.Stage
}

func hashValue(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}

func AddStudent(name, stage, snils string) error {
	exists, _, _ := CheckSnils(snils, false)
	if exists {
		return errors.New("студент с таким СНИЛС уже существует")
	}

	newStudent := Students{
		Name:  name,
		Stage: stage,
		Snils: hashValue(snils),
	}
	result := FamilyDB.Create(&newStudent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
