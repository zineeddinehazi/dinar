package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func DecodeJSON() ([]Item, error) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Item{}, nil
		}
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func EncodeJSON(items []Item) []byte {
	out, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding JSON (%v)", err)
	}
	return out
}
