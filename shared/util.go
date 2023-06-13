package shared

import (
	"encoding/json"
	"fmt"
)

func PrintAsJson(v any) {
	if j, err := json.Marshal(v); err == nil {
		fmt.Println(string(j))
	}
}
