// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAttributesToRetrieve(opts ...interface{}) *opt.AttributesToRetrieveOption {
    for _, o := range opts {
        if v, ok := o.(opt.AttributesToRetrieveOption); ok {
            return &v
        }
    }
    return nil
}