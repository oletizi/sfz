// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfz

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSfzListener is a complete listener for a parse tree produced by SfzParser.
type BaseSfzListener struct{}

var _ SfzListener = &BaseSfzListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSfzListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSfzListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSfzListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSfzListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSfz is called when production sfz is entered.
func (s *BaseSfzListener) EnterSfz(ctx *SfzContext) {}

// ExitSfz is called when production sfz is exited.
func (s *BaseSfzListener) ExitSfz(ctx *SfzContext) {}

// EnterLine is called when production line is entered.
func (s *BaseSfzListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseSfzListener) ExitLine(ctx *LineContext) {}

// EnterHeader_stmt is called when production header_stmt is entered.
func (s *BaseSfzListener) EnterHeader_stmt(ctx *Header_stmtContext) {}

// ExitHeader_stmt is called when production header_stmt is exited.
func (s *BaseSfzListener) ExitHeader_stmt(ctx *Header_stmtContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseSfzListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseSfzListener) ExitHeader(ctx *HeaderContext) {}

// EnterOpcode_stmt is called when production opcode_stmt is entered.
func (s *BaseSfzListener) EnterOpcode_stmt(ctx *Opcode_stmtContext) {}

// ExitOpcode_stmt is called when production opcode_stmt is exited.
func (s *BaseSfzListener) ExitOpcode_stmt(ctx *Opcode_stmtContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseSfzListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseSfzListener) ExitOpcode(ctx *OpcodeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSfzListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSfzListener) ExitValue(ctx *ValueContext) {}
