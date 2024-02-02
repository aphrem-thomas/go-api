package config

type Configuartion struct {
	db     string
	router string
}

func New(db string, router string) *Configuartion {
	return &Configuartion{
		db:     db,
		router: router,
	}
}
