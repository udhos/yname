package yname

import (
	"fmt"
	"strconv"
	"strings"
	//"gopkg.in/yaml.v2"
)

func splitLeft(str string, sep byte) (string, string) {
	slash := strings.IndexByte(str, sep)
	if slash < 0 {
		return str, ""
	}
	return str[:slash], str[slash+1:]
}

func splitRight(str string, sep byte) (string, string) {
	slash := strings.LastIndexByte(str, sep)
	if slash < 0 {
		return "", str
	}
	return str[:slash], str[slash+1:]
}

// Get retrieves subdoc using its path.
func Get(doc interface{}, path string) (interface{}, error) {
	head, tail := splitLeft(path, '/')

	if head == "" {
		return nil, fmt.Errorf("invalid empty path: [%s]", path)
	}

	switch i := doc.(type) {
	case map[interface{}]interface{}:
		child, found := i[head]
		if !found {
			return nil, fmt.Errorf("not found: [%s]", head)
		}
		if tail == "" {
			return child, nil
		}
		return Get(child, tail)
	case []interface{}:
		index, errConv := strconv.Atoi(head)
		if errConv != nil {
			return nil, fmt.Errorf("not and index: %s: %v", head, errConv)
		}
		if index < 0 || index >= len(i) {
			return nil, fmt.Errorf("index not found: %d", index)
		}
		child := i[index]
		if tail == "" {
			return child, nil
		}
		return Get(child, tail)
	case string:
		if i == head {
			if tail == "" {
				return nil, nil
			}
			return nil, fmt.Errorf("node has no children: [%s]: %s", head, i)
		}
		return nil, fmt.Errorf("node mismatch: [%s]: %s", head, i)
	}

	return nil, fmt.Errorf("unsupported type: [%s]: %v", head, doc)
}
