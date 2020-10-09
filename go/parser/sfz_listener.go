// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfz

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SfzListener is a complete listener for a parse tree produced by SfzParser.
type SfzListener interface {
	antlr.ParseTreeListener

	// EnterSfz is called when entering the sfz production.
	EnterSfz(c *SfzContext)

	// EnterSfzObject is called when entering the sfzObject production.
	EnterSfzObject(c *SfzObjectContext)

	// EnterHeaderObject is called when entering the headerObject production.
	EnterHeaderObject(c *HeaderObjectContext)

	// EnterHeaderName is called when entering the headerName production.
	EnterHeaderName(c *HeaderNameContext)

	// EnterOpcodeStatement is called when entering the opcodeStatement production.
	EnterOpcodeStatement(c *OpcodeStatementContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitSfz is called when exiting the sfz production.
	ExitSfz(c *SfzContext)

	// ExitSfzObject is called when exiting the sfzObject production.
	ExitSfzObject(c *SfzObjectContext)

	// ExitHeaderObject is called when exiting the headerObject production.
	ExitHeaderObject(c *HeaderObjectContext)

	// ExitHeaderName is called when exiting the headerName production.
	ExitHeaderName(c *HeaderNameContext)

	// ExitOpcodeStatement is called when exiting the opcodeStatement production.
	ExitOpcodeStatement(c *OpcodeStatementContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
