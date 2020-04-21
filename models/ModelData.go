package models

import (
	_struct "TutorialGoRestApi-Firebird/struct"
	"database/sql"
)

type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetDataTableSleeph5() (getStruct []_struct.StructData, err error) {
	row, err := model.DB.Query("select ID, FIRST_NAME, LAST_NAME FROM TABLE_SLEEPH5")
	if err != nil {
		return  nil,err
	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.Id,
				&data.FirstName,
				&data.LastName)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					Id:        data.Id,
					FirstName: data.FirstName,
					LastName:  data.LastName,
				}
				_isiStruct = append(_isiStruct, _data)

			}
		}

		return _isiStruct, nil
	}
}
