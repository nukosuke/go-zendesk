package zendesk

import (
	"strings"
)

type includeBuilder struct {
	keys []string
}

func (b *includeBuilder) addKey(s string) *includeBuilder {
	b.keys = append(b.keys, s)
	return b
}

func (b *includeBuilder) toInterface() (interface{}, error) {
	var csl strings.Builder

	for i, key := range b.keys {
		_, err := csl.WriteString(key)
		if err != nil {
			return nil, err
		}

		if i < len(b.keys)-1 {
			_, err := csl.WriteString(",")
			if err != nil {
				return nil, err
			}
		}
	}

	var data struct {
		Include string `url:"include"`
	}

	data.Include = csl.String()

	return data, nil
}

func (b *includeBuilder) path(basePath string) (string, error) {
	opts, err := b.toInterface()
	if err != nil {
		return "", err
	}

	return addOptions(basePath, opts)
}
