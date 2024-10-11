package sqlite

import "AnlossAPI/internal/config"

func New(cfg config.Config) {
	newRecords(cfg.RecordsPath)
	newStudents(cfg.StudentsPath)
}
