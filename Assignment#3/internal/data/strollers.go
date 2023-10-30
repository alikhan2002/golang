package data

import (
	"assignment2.alikhan.net/internal/validator"
	"context"
	"database/sql"
	"errors"
	"time"
)

type Stroller struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Brand     string    `json:"brand"`
	Price     int32     `json:"price"`
	Color     string    `json:"color"`
	Ages      string    `json.:"ages"`
	Version   int32     `json:"version"`
}

func ValidateStroller(v *validator.Validator, stroller *Stroller) {
	v.Check(stroller.Title != "", "title", "must be provided")
	v.Check(len(stroller.Title) <= 500, "title", "must not be more than 500 bytes long")
	//v.Check(stroller.Price == 0, "price", "must be provided")
	//v.Check(stroller.Price < 0, "price", "must be greater than 0")
	v.Check(len(stroller.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(stroller.Brand != "", "brand", "must be provided")
	v.Check(len(stroller.Brand) <= 500, "brand", "must not be more than 500 bytes long")
	v.Check(stroller.Color != "", "color", "must be provided")
	v.Check(len(stroller.Brand) <= 100, "color", "must not be more than 100 bytes long")
	v.Check(stroller.Ages != "", "ages", "must be provided")
	v.Check(len(stroller.Ages) <= 10, "ages", "must not be more than 10 bytes long")
}

type StrollerModel struct {
	DB *sql.DB
}

func (m StrollerModel) Insert(stroller *Stroller) error {
	// Define the SQL query for inserting a new record in the movies table and returning
	// the system-generated data.
	query := `
	INSERT INTO strollers (title, brand, price, color, ages)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, version`
	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []interface{}{stroller.Title, stroller.Brand, stroller.Price, stroller.Color, stroller.Ages}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Use QueryRowContext() and pass the context as the first argument.
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&stroller.ID, &stroller.CreatedAt, &stroller.Version)
}

}

func (m StrollerModel) Get(id int64) (*Stroller, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `
	SELECT pg_sleep(10), id, created_at, title, brand, price, color, ages, version
	FROM strollers
	WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var stroller Stroller
	// Use the context.WithTimeout() function to create a context.Context which carries a
	// 3-second timeout deadline. Note that we're using the empty context.Background()
	// as the 'parent' context.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Importantly, use defer to make sure that we cancel the context before the Get()
	// method returns.
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&stroller.ID,
		&stroller.CreatedAt,
		&stroller.Title,
		&stroller.Brand,
		&stroller.Price,
		&stroller.Color,
		&stroller.Ages,
		&stroller.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &stroller, nil
}

func (m StrollerModel) Update(stroller *Stroller) error {
	// Declare the SQL query for updating the record and returning the new version
	// number.
	query := `
		UPDATE strollers
		SET title = $1, brand = $2, price = $3, color = $4, ages = $5, version = version + 1
		WHERE id = $6 and version = $7
		RETURNING version`
	// Create an args slice containing the values for the placeholder parameters.
	args := []interface{}{
		stroller.Title,
		stroller.Brand,
		stroller.Price,
		stroller.Ages,
		stroller.Color,
		stroller.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx,query, args...).Scan(&stroller.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m StrollerModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
	DELETE FROM strollers
	WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := m.DB.ExecContext(ctx,query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
