// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestReferers(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.ReferersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Referers([]string{}...),
		},
		{
			opts:     []interface{}{opt.Referers("value1")},
			expected: opt.Referers("value1"),
		},
		{
			opts:     []interface{}{opt.Referers("value1", "value2", "value3")},
			expected: opt.Referers("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractReferers(c.opts...)
			out opt.ReferersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
