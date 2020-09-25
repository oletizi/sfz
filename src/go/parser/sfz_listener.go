// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfz

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SfzListener is a complete listener for a parse tree produced by SfzParser.
type SfzListener interface {
	antlr.ParseTreeListener

	// EnterSfz is called when entering the sfz production.
	EnterSfz(c *SfzContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterHeader_stmt is called when entering the header_stmt production.
	EnterHeader_stmt(c *Header_stmtContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterOpcode_stmt is called when entering the opcode_stmt production.
	EnterOpcode_stmt(c *Opcode_stmtContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitSfz is called when exiting the sfz production.
	ExitSfz(c *SfzContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitHeader_stmt is called when exiting the header_stmt production.
	ExitHeader_stmt(c *Header_stmtContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitOpcode_stmt is called when exiting the opcode_stmt production.
	ExitOpcode_stmt(c *Opcode_stmtContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
