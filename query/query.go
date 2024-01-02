package query

import (
	"gym/db"

	sq "github.com/Masterminds/squirrel"
)

type Exercise struct {
	ID   int    `db: "id"`
	Name string `db:"name"`
}

type Set struct {
	Exercise int     `db:"exercise"`
	Weight   float32 `db:"weight"`
	Reps     int     `db:"reps"`
}

func GetAllExercises() ([]Exercise, error) {
	sql, args, err := sq.Select("*").From("exercises").ToSql()
	if err != nil {
		return []Exercise{}, err
	}
	rows, err := db.DB.Query(sql, args...)
	if err != nil {
		return []Exercise{}, err
	}
	exercises := []Exercise{}
	for rows.Next() {
		exercise := Exercise{}
		err = rows.Scan(&exercise.ID, &exercise.Name)
		if err != nil {
			return []Exercise{}, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func GetAllSets() ([]Set, error) {
	sql, args, err := sq.Select("exercise", "weight", "reps").From("sets").ToSql()
	if err != nil {
		return []Set{}, err
	}
	rows, err := db.DB.Query(sql, args...)
	if err != nil {
		return []Set{}, err
	}
	sets := []Set{}
	for rows.Next() {
		set := Set{}
		err = rows.Scan(&set.Exercise, &set.Weight, &set.Reps)
		if err != nil {
			return []Set{}, err
		}
		sets = append(sets, set)
	}
	return sets, nil
}
