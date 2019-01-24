package html

import (
	"context"

	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

// InnerText returns inner text string of a given or matched by CSS selector element
// @param doc (HTMLDocument|HTMLElement) - Parent document or element.
// @param selector (String, optional) - String of CSS selector.
// @returns (String) - Inner text if an element found, otherwise empty string.
func InnerText(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 1, 2)

	if err != nil {
		return values.EmptyString, err
	}

	err = core.ValidateType(args[0], drivers.HTMLDocumentType, drivers.HTMLElementType)

	if err != nil {
		return values.None, err
	}

	node := args[0].(drivers.HTMLNode)

	if len(args) == 1 {
		return node.InnerText(), nil
	}

	err = core.ValidateType(args[1], core.StringType)

	if err != nil {
		return values.None, err
	}

	selector := args[1].(values.String)

	return node.InnerTextBySelector(selector), nil
}
