package normalizer

import "gopkg.in/bblfsh/sdk.v1/uast"

// ToNode is an instance of `uast.ObjectToNode`, defining how to transform an
// into a UAST (`uast.Node`).
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ObjectToNode
var ToNode = &uast.ObjectToNode{
	TopLevelIsRootNode: true,
	InternalTypeKey:    "type",
	OffsetKey:          "start",
	EndOffsetKey:       "end",
	LineKey:            "loc.start.line",
	EndLineKey:         "loc.end.line",
	ColumnKey:          "loc.start.column",
	EndColumnKey:       "loc.end.column",
}
