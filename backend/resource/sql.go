package resource

var (
	CreateConnectionTableSql = "CREATE TABLE IF NOT EXISTS connection (id text not null primary key, name text, urls text,user text,password text);"
)
