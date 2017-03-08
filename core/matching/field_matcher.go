package matching

import (
	"bytes"

	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
	"github.com/SpectoLabs/hoverfly/core/models"
	glob "github.com/ryanuber/go-glob"
)

func FieldMatcher(field *models.RequestFieldMatchers, toMatch string) bool {
	if field == nil {
		return true
	}

	if field.ExactMatch != nil {
		return glob.Glob(*field.ExactMatch, toMatch)
	}

	if field.XpathMatch != nil {
		xpathRule, err := goxpath.Parse(*field.XpathMatch)
		if err != nil {
			return false
		}

		xTree, err := xmltree.ParseXML(bytes.NewBufferString(toMatch))
		if err != nil {
			return false
		}

		result, err := xpathRule.Exec(xTree)
		if err != nil {
			return false
		}

		return len(result.String()) > 0
	}

	return false
}