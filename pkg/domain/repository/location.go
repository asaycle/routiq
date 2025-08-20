package repository

import "github.com/asaycle/routiq/pkg/domain/entity"

type Pref struct {
	ID          string `db:"id"`
	DisplayName string `db:"display_name"`
}

func fromPrefEntity(pref *entity.Pref) (*Pref, error) {
	return &Pref{
		ID:          pref.ID,
		DisplayName: pref.DisplayName,
	}, nil
}

func toPrefEntity(model *Pref) *entity.Pref {
	return entity.NewPref(model.ID, model.DisplayName)
}

type City struct {
	ID          string `db:"id"`
	PrefID      string `db:"pref_id"`
	DisplayName string `db:"display_name"`
}

func fromCityEntity(city *entity.City) (*City, error) {
	return &City{
		ID:          city.ID,
		PrefID:      city.PrefID,
		DisplayName: city.DisplayName,
	}, nil
}

func toCityEntity(model *City) *entity.City {
	return entity.NewCity(model.ID, model.PrefID, model.DisplayName)
}
