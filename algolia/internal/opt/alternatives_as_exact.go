// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAlternativesAsExact(opts ...interface{}) *opt.AlternativesAsExactOption {
    for _, o := range opts {
        if v, ok := o.(opt.AlternativesAsExactOption); ok {
            return &v
        }
    }
    return nil
}