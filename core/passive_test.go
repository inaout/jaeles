package core

//func TestPassiveCheck(t *testing.T) {
//	var record libs.Record
//	//record.Response.Beautify = "SQLite3::SQLException foo"
//	record.Response.Beautify = "Warning: file_exists(a"
//	var passive libs.Passive
//	passive = defaultPassive()
//
//	for _, rule := range passive.Rules {
//		for _, detectionString := range rule.Detections {
//			fmt.Println(detectionString)
//			_, result := RunDetector(record, detectionString)
//			fmt.Println(result)
//			if !result {
//				t.Errorf("Error resolve variable")
//			}
//		}
//	}
//}
//
//func TestRegexSearch(t *testing.T) {
//	raw := "SQLite3::SQLException foo"
//	regex := "(Exception (condition )?\\d+\\. Transaction rollback|com\\.frontbase\\.jdbc|org\\.h2\\.jdbc|Unexpected end of command in statement \\[\"|Unexpected token.*?in statement \\[|org\\.hsqldb\\.jdbc|CLI Driver.*?DB2|DB2 SQL error|\\bdb2_\\w+\\(|SQLSTATE.+SQLCODE|com\\.ibm\\.db2\\.jcc|Zend_Db_(Adapter|Statement)_Db2_Exception|Pdo[./_\\\\]Ibm|DB2Exception|Warning.*?\\Wifx_|Exception.*?Informix|Informix ODBC Driver|ODBC Informix driver|com\\.informix\\.jdbc|weblogic\\.jdbc\\.informix|Pdo[./_\\\\]Informix|IfxException|Warning.*?\\Wingres_|Ingres SQLSTATE|Ingres\\W.*?Driver|com\\.ingres\\.gcf\\.jdbc|Dynamic SQL Error|Warning.*?\\Wibase_|org\\.firebirdsql\\.jdbc|Pdo[./_\\\\]Firebird|Microsoft Access (\\d+ )?Driver|JET Database Engine|Access Database Engine|ODBC Microsoft Access|Syntax error \\(missing operator\\) in query expression|Driver.*? SQL[\\-\\_\\ ]*Server|OLE DB.*? SQL Server|\\bSQL Server[^&lt;&quot;]+Driver|Warning.*?\\W(mssql|sqlsrv)_|\\bSQL Server[^&lt;&quot;]+[0-9a-fA-F]{8}|System\\.Data\\.SqlClient\\.SqlException|(?s)Exception.*?\\bRoadhouse\\.Cms\\.|Microsoft SQL Native Client error '[0-9a-fA-F]{8}|\\[SQL Server\\]|ODBC SQL Server Driver|ODBC Driver \\d+ for SQL Server|SQLServer JDBC Driver|com\\.jnetdirect\\.jsql|macromedia\\.jdbc\\.sqlserver|Zend_Db_(Adapter|Statement)_Sqlsrv_Exception|com\\.microsoft\\.sqlserver\\.jdbc|Pdo[./_\\\\](Mssql|SqlSrv)|SQL(Srv|Server)Exception|SQL syntax.*?MySQL|Warning.*?\\Wmysqli?_|MySQLSyntaxErrorException|valid MySQL result|check the manual that corresponds to your (MySQL|MariaDB) server version|Unknown column '[^ ]+' in 'field list'|MySqlClient\\.|com\\.mysql\\.jdbc|Zend_Db_(Adapter|Statement)_Mysqli_Exception|Pdo[./_\\\\]Mysql|MySqlException|\\bORA-\\d{5}|Oracle error|Oracle.*?Driver|Warning.*?\\W(oci|ora)_|quoted string not properly terminated|SQL command not properly ended|macromedia\\.jdbc\\.oracle|oracle\\.jdbc|Zend_Db_(Adapter|Statement)_Oracle_Exception|Pdo[./_\\\\](Oracle|OCI)|OracleException|PostgreSQL.*?ERROR|Warning.*?\\Wpg_|valid PostgreSQL result|Npgsql\\.|PG::SyntaxError:|org\\.postgresql\\.util\\.PSQLException|ERROR:\\s\\ssyntax error at or near|ERROR: parser: parse error at or near|PostgreSQL query failed|org\\.postgresql\\.jdbc|Pdo[./_\\\\]Pgsql|PSQLException|SQL error.*?POS([0-9]+)|Warning.*?\\Wmaxdb_|DriverSapDB|com\\.sap\\.dbtech\\.jdbc|SQLite/JDBCDriver|SQLite\\.Exception|(Microsoft|System)\\.Data\\.SQLite\\.SQLiteException|Warning.*?\\W(sqlite_|SQLite3::)|\\[SQLITE_ERROR\\]|SQLite error \\d+:|sqlite3.OperationalError:|SQLite3::SQLException|org\\.sqlite\\.JDBC|Pdo[./_\\\\]Sqlite|SQLiteException|Warning.*?\\Wsybase_|Sybase message|Sybase.*?Server message|SybSQLException|Sybase\\.Data\\.AseClient|com\\.sybase\\.jdbc)"
//
//	_, result := RegexSearch(raw, regex)
//	if !result {
//		t.Errorf("Error resolve variable")
//	}
//}
