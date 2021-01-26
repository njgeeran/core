package log

// GormLogger struct
type GormLogger struct{}

// Print - Log Formatter
func (*GormLogger) Print (v ...interface{}) {
	switch v[0] {
	case "sql":
		GetLoger().Debugw("sql",
			"module","gorm",
			"type","sql",
			"src",v[1],
			"duration",v[2],
			"sql",v[3],
			"values",v[4],
			"rows_returned",v[5])
	case "log":
		GetLoger().Debugw("gorm_log",
			"desc",v[2])
	}
}
