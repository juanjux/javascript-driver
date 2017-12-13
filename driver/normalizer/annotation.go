package normalizer

import (
	"github.com/bblfsh/javascript-driver/driver/normalizer/babylon"
	"gopkg.in/bblfsh/sdk.v1/uast"
	. "gopkg.in/bblfsh/sdk.v1/uast/ann"
	"gopkg.in/bblfsh/sdk.v1/uast/transformer"
	"gopkg.in/bblfsh/sdk.v1/uast/transformer/annotatter"
)

// Transformers is the list of `transformer.Transfomer` to apply to a UAST, to
// learn more about the Transformers and the available ones take a look to:
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/transformers
var Transformers = []transformer.Tranformer{
	annotatter.NewAnnotatter(AnnotationRules),
}

// AnnotationRules describes how a UAST should be annotated with `uast.Role`.
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/ann
// AnnotationRules annotate a UAST with roles.
var AnnotationRules = On(babylon.File).Roles(uast.File).Descendants(
	On(babylon.Program).Roles(uast.Module).Descendants(
		// Identifiers
		On(babylon.Identifier).Roles(uast.Expression, uast.Identifier),

		// Literals
		On(babylon.RegExpLiteral).Roles(uast.Expression, uast.Literal, uast.Regexp),
		On(babylon.NullLiteral).Roles(uast.Expression, uast.Literal, uast.Null),
		On(babylon.StringLiteral).Roles(uast.Expression, uast.Literal, uast.String),
		On(babylon.BooleanLiteral).Roles(uast.Expression, uast.Literal, uast.Boolean),
		On(babylon.NumericLiteral).Roles(uast.Expression, uast.Literal, uast.Number),

		// Functions
		On(Or(babylon.FunctionDeclaration, babylon.ArrowFunctionExpression, babylon.FunctionExpression, babylon.ObjectMethod)).Roles(uast.Declaration, uast.Function).Children(
			On(babylon.PropertyId).Roles(uast.Function, uast.Name),
			On(babylon.PropertyParams).Roles(uast.Function, uast.Argument).Self(
				On(babylon.RestElement).Roles(uast.ArgsList),
			),
			On(babylon.PropertyBody).Roles(uast.Function, uast.Body),
		),

		// Statements
		On(babylon.ExpressionStatement).Roles(uast.Statement),
		On(babylon.BlockStatement).Roles(uast.Statement, uast.Block, uast.Scope),
		On(babylon.EmptyStatement).Roles(uast.Statement),
		On(babylon.DebuggerStatement).Roles(uast.Statement, uast.Incomplete),
		On(babylon.WithStatement).Roles(uast.Statement, uast.Scope, uast.Block, uast.Incomplete).Children(
			On(babylon.PropertyObject).Roles(uast.Incomplete),
		),

		// Control flow
		On(babylon.ReturnStatement).Roles(uast.Statement, uast.Return),
		On(babylon.LabeledStatement).Roles(uast.Statement, uast.Incomplete),
		On(babylon.BreakStatement).Roles(uast.Statement, uast.Break),
		On(babylon.ContinueStatement).Roles(uast.Statement, uast.Continue),

		// Choice
		On(babylon.IfStatement).Roles(uast.Statement, uast.If).Children(
			On(babylon.PropertyTest).Roles(uast.If, uast.Condition),
			On(babylon.PropertyConsequent).Roles(uast.If, uast.Then, uast.Body),
			On(babylon.PropertyAlternate).Roles(uast.If, uast.Else, uast.Body),
		),
		On(babylon.SwitchStatement).Roles(uast.Statement, uast.Switch).Children(
			On(babylon.PropertyDiscriminant).Roles(uast.Switch, uast.Condition),
		),
		On(babylon.SwitchCase).Roles(uast.Switch, uast.Case).Children(
			On(babylon.PropertyTest).Roles(uast.Case, uast.Condition),
		),

		// Exception
		On(babylon.ThrowStatement).Roles(uast.Statement, uast.Throw),
		On(babylon.TryStatement).Roles(uast.Statement, uast.Try).Children(
			On(babylon.PropertyFinalizer).Roles(uast.Try, uast.Finally),
		),
		On(babylon.CatchClause).Roles(uast.Try, uast.Catch),

		// Loops
		On(babylon.WhileStatement).Roles(uast.Statement, uast.While).Children(
			On(babylon.PropertyTest).Roles(uast.While, uast.Condition),
			On(babylon.PropertyBody).Roles(uast.While, uast.Body),
		),
		On(babylon.DoWhileStatement).Roles(uast.Statement, uast.DoWhile).Children(
			On(babylon.PropertyTest).Roles(uast.DoWhile, uast.Condition),
			On(babylon.PropertyBody).Roles(uast.DoWhile, uast.Body),
		),
		On(babylon.ForStatement).Roles(uast.Statement, uast.For).Children(
			On(babylon.PropertyInit).Roles(uast.For, uast.Initialization),
			On(babylon.PropertyTest).Roles(uast.For, uast.Condition),
			On(babylon.PropertyUpdate).Roles(uast.For, uast.Update),
		),
		On(Or(babylon.ForInStatement, babylon.ForOfStatement)).Roles(uast.Statement, uast.For, uast.Iterator).Children(
			On(babylon.PropertyLeft).Roles(uast.For, uast.Iterator),
			On(babylon.PropertyRight).Roles(uast.For),
			On(babylon.PropertyBody).Roles(uast.For, uast.Body),
		),

		// Declarations
		On(babylon.FunctionDeclaration).Roles(uast.Statement),
		On(babylon.VariableDeclaration).Roles(uast.Statement, uast.Declaration, uast.Variable),
		On(babylon.VariableDeclarator).Roles(uast.Declaration, uast.Variable).Children(
			On(babylon.PropertyInit).Roles(uast.Initialization),
		),

		// Misc
		On(babylon.Decorator).Roles(uast.Incomplete),
		On(babylon.Directive).Roles(uast.Incomplete),
		On(babylon.DirectiveLiteral).Roles(uast.Expression, uast.Literal, uast.Incomplete),

		// Expressions
		On(babylon.Super).Roles(uast.Expression, uast.Identifier, uast.Base),
		On(babylon.Import).Roles(uast.Expression, uast.Import),
		On(babylon.ThisExpression).Roles(uast.Expression, uast.This),
		On(babylon.ArrowFunctionExpression).Roles(uast.Expression),
		On(babylon.YieldExpression).Roles(uast.Expression, uast.Return, uast.Incomplete),
		On(babylon.AwaitExpression).Roles(uast.Expression, uast.Incomplete),
		On(babylon.ArrayExpression).Roles(uast.Expression, uast.Initialization, uast.List, uast.Literal),
		On(babylon.ObjectExpression).Roles(uast.Expression, uast.Initialization, uast.Map, uast.Literal),

		// Object properties
		On(babylon.ObjectMethod).Roles(uast.Map).Children(
			On(babylon.PropertyKey).Roles(uast.Map, uast.Key, uast.Function, uast.Name),
			On(babylon.PropertyBody).Roles(uast.Map, uast.Value),
		),
		On(babylon.ObjectProperty).Roles(uast.Map).Children(
			On(babylon.PropertyKey).Roles(uast.Map, uast.Key),
			On(babylon.PropertyValue).Roles(uast.Map, uast.Value),
		),

		// Function expressions
		On(babylon.FunctionExpression).Roles(uast.Expression),
		On(babylon.CallExpression).Roles(uast.Expression, uast.Call).Children(
			On(babylon.PropertyCallee).Roles(uast.Call, uast.Callee),
			On(babylon.PropertyArguments).Roles(uast.Call, uast.Argument),
			On(babylon.SpreadElement).Roles(uast.ArgsList),
		),

		// Unary expressions
		On(Or(babylon.UnaryExpression, babylon.UpdateExpression)).Roles(uast.Expression, uast.Unary, uast.Operator).Self(
			On(HasProperty("prefix", "false")).Roles(uast.Postfix),

			// Unary operators
			On(HasProperty("operator", "-")).Roles(uast.Arithmetic, uast.Negative),
			On(HasProperty("operator", "+")).Roles(uast.Arithmetic, uast.Positive),
			On(HasProperty("operator", "!")).Roles(uast.Boolean, uast.Not),
			On(HasProperty("operator", "~")).Roles(uast.Bitwise, uast.Not),
			On(HasProperty("operator", "typeof")).Roles(uast.Type),
			On(HasProperty("operator", "void")).Roles(uast.Null),
			On(HasProperty("operator", "delete")).Roles(uast.Incomplete),
			On(HasProperty("operator", "throw")).Roles(uast.Throw),

			// Update operators
			On(HasProperty("operator", "++")).Roles(uast.Arithmetic, uast.Increment),
			On(HasProperty("operator", "--")).Roles(uast.Arithmetic, uast.Decrement),
		),

		On(babylon.MemberExpression).Roles(uast.Qualified, uast.Expression, uast.Identifier),
		On(babylon.BinaryExpression).Roles(uast.Expression, uast.Binary),
	),
)
