package models

import (
	"context"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
	cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
	lifespan, coalesce(details, '') as details, coalesce(alternate_names, '') as alternate_names, coalesce(geographic_origin, '') as geographic_origin
	from dog_breeds order by breed`

	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.AverageWeight,
			&b.LifeSpan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)

		if err != nil {
			return nil, err
		}

		breeds = append(breeds, &b)
	}

	return breeds, nil
}

func (m *mysqlRepository) GetBreedByName(breed string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
	cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
	lifespan, coalesce(details, '') as details, coalesce(alternate_names, '') as alternate_names, coalesce(geographic_origin, '') as geographic_origin
	from dog_breeds where breed = ?`

	row := m.DB.QueryRowContext(ctx, query, breed)
	var dogBreed DogBreed
	err := row.Scan(
		&dogBreed.ID,
		&dogBreed.Breed,
		&dogBreed.WeightLowLbs,
		&dogBreed.WeightHighLbs,
		&dogBreed.AverageWeight,
		&dogBreed.LifeSpan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)

	if err != nil {
		return nil, err
	}

	return &dogBreed, nil
}

func (m *mysqlRepository) GetDogOfMonthById(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, video, image from dog_of_month where id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var dogOfMonth DogOfMonth
	err := row.Scan(
		&dogOfMonth.ID,
		&dogOfMonth.Image,
		&dogOfMonth.Video,
	)

	if err != nil {
		return nil, err
	}

	return &dogOfMonth, nil
}
