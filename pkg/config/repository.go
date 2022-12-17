package config

type SQLDBSettings struct {
	SqlDsn              string
	SqlMaxOpenConns     int
	SqlMaxIdleConns     int
	SqlConnsMaxLifetime int
}

func (s *SQLDBSettings) DSN() string {
	return s.SqlDsn
}

func (s *SQLDBSettings) MaxOpenConns() int {
	return s.SqlMaxOpenConns
}

func (s *SQLDBSettings) MaxIdleConns() int {
	return s.SqlMaxIdleConns
}

func (s *SQLDBSettings) ConnsMaxLifetime() int {
	return s.SqlConnsMaxLifetime
}
