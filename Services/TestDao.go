package Services

import (
	. "WebApi/Model"
	"encoding/json"
)

// InsertTest /*插入数据*/
func InsertTest(test Test) string {
	var db=Init("Test")
    db.Create(test)
	j,error:=json.MarshalIndent(test,"","")
	if error==nil{
		return string(j)
	}
	return ""
}
