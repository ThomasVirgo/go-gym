package query

import (
	"fmt"
	"gym/db"
	"sort"
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
type WeightAndReps struct {
	Weight float32
	Reps   int
}

type GymSession struct {
	Date       time.Time
	DateString string
	Exercises  map[string][]WeightAndReps
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

func GetGymSessions() ([]GymSession, error) {
	sql, args, err := sq.Select("exercises.name", "weight", "reps", "created_at").From("sets").Join("exercises ON sets.exercise = exercises.id").ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	sets := []Set{}
	date_to_exercise_map := make(map[time.Time]map[string][]WeightAndReps)
	for rows.Next() {
		set := Set{}
		err = rows.Scan(&set.ExerciseName, &set.Weight, &set.Reps, &set.CreatedAt)
		if err != nil {
			return nil, err
		}
		year, month, day := set.CreatedAt.Date()
		date_key := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
		set.FormattedCreatedAt = fmt.Sprintf("%d %s %d", day, month.String(), year)
		set_to_add := WeightAndReps{set.Weight, set.Reps}
		exercises_on_date, found := date_to_exercise_map[date_key]
		if found {
			sets_for_exercise, found_exercise := exercises_on_date[set.ExerciseName]
			if found_exercise {
				exercises_on_date[set.ExerciseName] = append(sets_for_exercise, set_to_add)
			} else {
				exercises_on_date[set.ExerciseName] = []WeightAndReps{set_to_add}
			}
		} else {
			date_to_exercise_map[date_key] = map[string][]WeightAndReps{set.ExerciseName: {set_to_add}}
		}
		sets = append(sets, set)
	}
	gym_sessions := []GymSession{}
	for date, exercises := range date_to_exercise_map {
		year, month, day := date.Date()
		new_session := GymSession{date, fmt.Sprintf("%d %s %d", day, month.String(), year), exercises}
		gym_sessions = append(gym_sessions, new_session)
	}
	sort.Slice(gym_sessions, func(i, j int) bool {
		return gym_sessions[i].Date.After(gym_sessions[j].Date)
	})

	return gym_sessions, nil
}
