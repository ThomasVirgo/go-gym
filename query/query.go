package query

import (
	"gym/db"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type Exercise struct {
	ID   int    `db: "id"`
	Name string `db:"name"`
}

type Set struct {
	ExerciseName       string    `db:"name"`
	Weight             float32   `db:"weight"`
	Reps               int       `db:"reps"`
	CreatedAt          time.Time `db:"created_at"`
	FormattedCreatedAt string
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
	sql, args, err := sq.Select("exercises.name", "weight", "reps", "created_at").From("sets").Join("exercises ON sets.exercise = exercises.id").ToSql()
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
		err = rows.Scan(&set.ExerciseName, &set.Weight, &set.Reps, &set.CreatedAt)
		if err != nil {
			return []Set{}, err
		}
		set.FormattedCreatedAt = set.CreatedAt.Format("2nd Jan")
		sets = append(sets, set)
	}
	return sets, nil
}
