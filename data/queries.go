package data

func getCountQuery(tableName string) string {
	return "SELECT COUNT(*) as count FROM " + tableName
}

func getSelectAllQuery(tableName string) string {
	return "SELECT * FROM " + tableName
}
