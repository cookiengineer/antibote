package structs

func (cache *Cache) AddKey(key string, name string) {

	names, ok := cache.KeyMap[key]

	if ok == true {

		found := false

		for n := 0; n < len(names); n++ {

			if names[n] == name {
				found = true
			}

		}

		if found == false {
			cache.KeyMap[key] = append(cache.KeyMap[key], name)
		}

	} else {
		cache.KeyMap[key] = []string{name}
	}


}
