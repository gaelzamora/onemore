package db

import (
	"database/sql"
	"onemore/models"
	"onemore/tools"
	"strconv"
)

func InsertMuscle(m models.Muscle) (int64, error) {
	err := DBConnect()

	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentence := "INSERT INTO muscle (Name_Muscle) VALUES ('"+tools.ScapeString(m.Name_Muscle)+"')"

	var result sql.Result

	result, err = Db.Exec(sentence)

	if err != nil {
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	return LastInsertId, nil
}

func UpdateMuscle(m models.Muscle) (error) {
	
	err := DBConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := "UPDATE Muscle SET "
	sentence = tools.Sentence(sentence, "Name_Muscle", "S", 0, 0, m.Name_Muscle)
	sentence += " WHERE Id_Muscle = "+strconv.Itoa(m.Id_Muscle)

	_, err = Db.Exec(sentence)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMuscle(id int) error {
	err := DBConnect()

	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := "DELETE from Muscle WHERE Id_Muscle = "+strconv.Itoa(id)

	_, err2 := Db.Exec(sentence)

	if err2 != nil {
		return err2
	}

	return nil
}