// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractRemoveWordsIfNoResults(opts ...interface{}) *opt.RemoveWordsIfNoResultsOption {
    for _, o := range opts {
        if v, ok := o.(opt.RemoveWordsIfNoResultsOption); ok {
            return &v
        }
    }
    return nil
}