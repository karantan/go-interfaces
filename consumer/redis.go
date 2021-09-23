package consumer

import (
	"go-interfaces/good"
)

// ProcessAnimalsRedis ilustrates how a consumer would use the "good" version of the
// redis implementation (i.e. accept interfaces return struct approach).
func ProcessAnimalsRedis(rdb good.RedisSource, animals []Animal) error {
	for _, a := range animals {
		log.Infow("Storing our animal to redis", "animal", a.name)
		if err := good.SetKey(rdb, string(a.hash), a.name); err != nil {
			return err
		}

		value, err := good.GetKey(rdb, string(a.hash))
		if err != nil {
			return err
		}
		log.Infow("Animal saved to rdb", string(a.hash), value)
	}
	return nil
}
