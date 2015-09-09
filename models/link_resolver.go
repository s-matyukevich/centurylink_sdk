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
	err = fmt.Errorf("There is no link with rel %s in model %T", linkRel, linkModel)
	return
}

func checkVerb(link Link, verb string) bool {
	if len(link.Verbs) == 0 {
		return verb == "GET"
	}
	for _, v := range link.Verbs {
		if v == verb {
			return true
		}
	}
	return false
}

func ResolveLink(model LinkModel, linkRel string, verb string, resModel interface{}) (err error) {
	link, err := getLink(model, linkRel)
	if err != nil {
		return
	}

	if !checkVerb(link, verb) {
		err = fmt.Errorf("There is no support for verb %s in link rel %q of model %T",
				 verb, link.Rel, model)
		return
	}

	cn := model.GetConnection()
	if cn == nil {
		err = fmt.Errorf("Model connection is not initialized.")
		return
	}
	err = cn.ExecuteRequest(verb, link.Href, nil, resModel)
	return
}
