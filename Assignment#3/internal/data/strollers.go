package data

import (
	"assignment2.alikhan.net/internal/validator"
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
	return m.DB.QueryRow(query, args...).Scan(&stroller.ID, &stroller.CreatedAt, &stroller.Version)

}

func (m StrollerModel) Get(id int64) (*Stroller, error) {
	// The PostgreSQL bigserial type that we're using for the movie ID starts
	// auto-incrementing at 1 by default, so we know that no movies will have ID values
	// less than that. To avoid making an unnecessary database call, we take a shortcut
	// and return an ErrRecordNotFound error straight away.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `
	SELECT id, created_at, title, brand, price, color, ages, version
	FROM strollers
	WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var stroller Stroller
	// Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	err := m.DB.QueryRow(query, id).Scan(
		&stroller.ID,
		&stroller.CreatedAt,
		&stroller.Title,
		&stroller.Brand,
		&stroller.Price,
		&stroller.Color,
		&stroller.Ages,
		&stroller.Version,
	)
	// Handle any errors. If there was no matching movie found, Scan() will return
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound
	// error instead.
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
		WHERE id = $3
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
	return m.DB.QueryRow(query, args...).Scan(&stroller.Version)
}

func (m StrollerModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
	DELETE FROM strollers
	WHERE id = $1`
	result, err := m.DB.Exec(query, id)
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
