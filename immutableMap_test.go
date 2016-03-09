package immutableMap

import (
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	immutableMap := NewImmutableMap()
	go func() {
		for i := 0; i < 1000; i++ {
			immutableMap.Set(i, i+1)
			if value, err := immutableMap.Get(i); err != nil {
				log.Printf("Setting (%d: %d) into immutable map\n", i, value)
			} else {
				log.Println(value)
			}
		}
	}()
	for i := 0; i < 1000; i++ {
		if err := immutableMap.Delete(i); err != nil {
			log.Println(err)
		} else {
			log.Printf("Deleting (%d) from immutable map\n", i)
		}
	}
}
