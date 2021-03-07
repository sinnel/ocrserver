package db

import "errors"

//OcrResult Model
type OcrResult struct {
	UserName    string
	ImagePath   string
	ImageResult string
}

//OcrResultOperator operator
type OcrResultOperator struct{}

//Insert OcrResult
func (operator *OcrResultOperator) Insert(or *OcrResult) error {
	if or == nil {
		return errors.New("v is nil")
	}

	stmt, err := db.Prepare("insert into ocr_result(username,image_path,image_result) values($1,$2,$3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(or.UserName, or.ImagePath, or.ImageResult)
	if err != nil {
		return err
	}

	return nil
}
