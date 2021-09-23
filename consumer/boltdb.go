package consumer

import (
	"go-interfaces/good"
)

// ProcessAnimalsBolt ilustrates how a consumer would use the "good" version of the boltDB
// implementation (i.e. accept interfaces return struct approach).
func ProcessAnimalsBolt(db good.DataSource, animals []Animal) error {
	for _, a := range animals {
		log.Infow("Storing our animal to persistant boltDB", "animal", a.name)
		if err := good.Put(db, "animals", string(a.hash), a.name); err != nil {
			return err
		}

		value, err := good.Get(db, "animals", string(a.hash))
		if err != nil {
			return err
		}
		log.Infow("Animal saved to boltDB", string(a.hash), value)
	}
	return nil
}
