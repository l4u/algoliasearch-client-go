// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSafe(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.SafeOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Safe(false),
		},
		{
			opts:     []interface{}{opt.Safe(true)},
			expected: opt.Safe(true),
		},
		{
			opts:     []interface{}{opt.Safe(false)},
			expected: opt.Safe(false),
		},
	} {
		var (
			in  = ExtractSafe(c.opts...)
			out opt.SafeOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}