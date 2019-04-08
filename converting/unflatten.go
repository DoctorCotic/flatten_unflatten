package converting

import "strings"

func Unflatten(nested map[string]interface{}) (map[string]interface{}, error) {
	var tree = make(map[string]interface{})
	for k, v := range nested {
		keyParts := strings.Split(k, ".")
		tr := tree
		for _, tk := range keyParts[:len(keyParts)-1] {
			newTree, ok := tr[tk]
			if !ok {
				newTree = make(map[string]interface{})
				tr[tk] = newTree
			}
			tr = newTree.(map[string]interface{})
		}
		tr[keyParts[len(keyParts)-1]] = v
	}
	return tree, nil
}