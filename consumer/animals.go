package consumer

import "crypto/sha1"

type Animal struct {
	hash []byte
	name string
}

// GetAnimal function mimics retrieving/processing some animal data
func GetAnimal(animal string) Animal {
	h := sha1.New()
	h.Write([]byte(animal))
	bs := h.Sum(nil)

	return Animal{hash: bs, name: animal}
}
