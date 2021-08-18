package Services

import (
	. "WebApi/Model"
)

// InsertTest /*插入数据*/
func InsertTest(test Test) Test {
	var db = Init("Test")
	db.Create(test)
	return test
}
