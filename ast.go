package jsongoparser

import (
	"fmt"
	"strconv"
	"strings"
)

// Object represents a JSON object - a collection of key-value pairs.
type Object struct {
	// Token is the opening '{' token
	Token Token
	// Pairs are the key-value pairs in the object.
	Pairs map[string]Value
}

// TokenLiteral returns the literal value of the token that defines the object.
func (o *Object) TokenLiteral() string { return o.Token.Literal }

// String returns a simplified string representation of the object.
func (o *Object) String() string {
	var b strings.Builder

	b.WriteString("{")

	i := 0
	for k, v := range o.Pairs {
		if i > 0 {
			b.WriteString(", ")
		}

		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(v.String())

		i++
	}

	b.WriteString("}")

	return b.String()
}

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (o *Object) valueNode() {}

// Array represents a JSON array - an ordered list of values.
type Array struct {
	// Token is the opening '[' token.
	Token Token
	// Elements are the values in the array.
	Elements []Value
}

// TokenLiteral returns the literal value of the token that defines the array.
func (a *Array) TokenLiteral() string { return a.Token.Literal }

// String returns a simplified string representation of the array.
func (a *Array) String() string { return "[]" } // Simplified for now

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (a *Array) valueNode() {}

// StringLiteral represents a JSON string value.
type StringLiteral struct {
	// Token is the string token.
	Token Token
	// Value is the actual string value.
	Value string
}

// TokenLiteral returns the literal value of the token that defines the string.
func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }

// String returns the actual string value.
func (s *StringLiteral) String() string { return s.Value }

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (s *StringLiteral) valueNode() {}

// NumberLiteral represents a JSON number value.
type NumberLiteral struct {
	// Token is the number token.
	Token Token
	// Value is the number as a string (we'll parse it when needed).
	Value string
	// Float is the actual float value of the number.
	Float float64
	// Int is the actual integer value of the number.
	Int int64
	// IsInt is a flag to indicate if the number is an integer.
	IsInt bool
	// IsValid is a flag to indicate if the number is valid JSON number.
	IsValid bool
}

// NewNumberLiteral creates a new NumberLiteral with proper validation and parsing
func NewNumberLiteral(token Token) *NumberLiteral {
	n := &NumberLiteral{
		Token: token,
		Value: token.Literal,
	}

	// Try parsing as int first
	if i, err := strconv.ParseInt(token.Literal, 10, 64); err == nil {
		n.Int = i
		n.Float = float64(i)
		n.IsInt = true
		n.IsValid = true

		return n
	}

	// Try parsing as float
	if f, err := strconv.ParseFloat(token.Literal, 64); err == nil {
		n.Float = f
		n.IsInt = false
		n.IsValid = true

		return n
	}

	// If we get here, the number is not valid
	n.IsValid = false

	return n
}

// TokenLiteral returns the literal value of the token that defines the number.
func (n *NumberLiteral) TokenLiteral() string { return n.Token.Literal }

// String returns the number value as a string.
func (n *NumberLiteral) String() string {
	if n.IsInt {
		return fmt.Sprintf("%d", n.Int)
	}

	return fmt.Sprintf("%f", n.Float)
}

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (n *NumberLiteral) valueNode() {}

// IsValidNumber returns whether the number is a valid JSON number
func (n *NumberLiteral) IsValidNumber() bool {
	return n.IsValid
}

// Boolean represents a JSON boolean value (true or false).
type Boolean struct {
	// Token is the boolean token.
	Token Token
	// Value is the actual boolean value.
	Value bool
}

// TokenLiteral returns the literal value of the token that defines the boolean.
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String returns the boolean value as a string.
func (b *Boolean) String() string { return b.Token.Literal }

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (b *Boolean) valueNode() {}

// Null represents a JSON null value.
type Null struct {
	// Token is the null token.
	Token Token
}

// TokenLiteral returns the literal value of the token that defines the null value.
func (n *Null) TokenLiteral() string { return n.Token.Literal }

// String returns the string representation of the null value.
func (n *Null) String() string { return "null" }

// valueNode is a placeholder method to ensure type safety within the Value interface.
func (n *Null) valueNode() {}
