// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/example.proto

package example

import options "github.com/infobloxopen/protoc-gen-atlas-query-perm/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"

// Reference imports to suppress errors if they are not otherwise used.

var exampleMessagesRequiredValidation = map[string]map[string]options.FilteringOption{
	"User": {
		"first_name":  options.FilteringOption{Deny: []string{"LT", "LE", "EQ", "GT", "GE"}},
		"weight":      options.FilteringOption{Deny: []string{"LE"}},
		"on_vacation": options.FilteringOption{DisableSorting: true},
		"speciality":  options.FilteringOption{DisableSorting: true, Deny: []string{"EQ", "GT", "GE", "LT", "LE"}},
		"comment":     options.FilteringOption{},
	},
	"ListRequest": {
		"filter":   options.FilteringOption{},
		"order_by": options.FilteringOption{},
		"fields":   options.FilteringOption{},
		"paging":   options.FilteringOption{},
	},
	"ReadRequest": {
		"order_by": options.FilteringOption{},
		"fields":   options.FilteringOption{},
		"paging":   options.FilteringOption{},
	},
	"UserResponse": {
		"data": options.FilteringOption{},
	},
}
var exampleMethodsRequiredFilteringValidation = map[string]string{
	"/example.TestService/List": "User",
}
var exampleMethodsRequiredSortingValidation = map[string]string{
	"/example.TestService/List": "User",
	"/example.TestService/Read": "User",
}

func Validate(f *query.Filtering, p *query.Sorting, methodName string) error {
	perm, ok := exampleMethodsRequiredFilteringValidation[methodName]
	if !ok {
		return nil
	}
	res := options.ValidateFilteringPermissions(f, perm, exampleMessagesRequiredValidation)
	if res != nil {
		return res
	}
	perm, ok = exampleMethodsRequiredSortingValidation[methodName]
	if !ok {
		return nil
	}
	res = options.ValidateSortingPermissions(p, perm, exampleMessagesRequiredValidation)
	if res != nil {
		return res
	}
	return nil
}
