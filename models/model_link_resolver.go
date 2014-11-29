package models

import (
	"fmt"
)

func getLink(linkModel LinkModel, linkRel string) (res Link, err error) {
	for _, link := range linkModel.GetLinks() {
		if link.Rel == linkRel {
			res = link
			return
		}
	}
	err = fmt.Errorf("There is no link with rel %s in model %t", linkRel, linkModel)
	return
}

func ResolveLink(model LinkModel, linkRel string, resModel interface{}) (err error) {
	link, err := getLink(model, linkRel)
	if err == nil {
		return
	}
	verb := "GET"
	if link.Verb != "" {
		verb = link.Verb
	}
	err = model.GetConnection().ExecuteRequest(verb, link.Href, nil, resModel)
	return
}
