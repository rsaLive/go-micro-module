package common

import (
	"github.com/oa-meeting/pkg/db"
)

func Init() {
	db.DBMigration()
}
