// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractHitsPerPage(opts ...interface{}) *opt.HitsPerPageOption {
    for _, o := range opts {
        if v, ok := o.(opt.HitsPerPageOption); ok {
            return &v
        }
    }
    return nil
}