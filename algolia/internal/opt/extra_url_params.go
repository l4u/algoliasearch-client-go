// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractExtraURLParams(opts ...interface{}) *opt.ExtraURLParamsOption {
	for _, o := range opts {
		if v, ok := o.(opt.ExtraURLParamsOption); ok {
			return &v
		}
	}
	return nil
}
