package healthz

import (
	"database/sql"
)

type Dependency struct {
	databaseStats sql.DBStats
}
