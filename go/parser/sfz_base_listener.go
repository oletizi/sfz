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

// EnterSfzObject is called when production sfzObject is entered.
func (s *BaseSfzListener) EnterSfzObject(ctx *SfzObjectContext) {}

// ExitSfzObject is called when production sfzObject is exited.
func (s *BaseSfzListener) ExitSfzObject(ctx *SfzObjectContext) {}

// EnterHeaderObject is called when production headerObject is entered.
func (s *BaseSfzListener) EnterHeaderObject(ctx *HeaderObjectContext) {}

// ExitHeaderObject is called when production headerObject is exited.
func (s *BaseSfzListener) ExitHeaderObject(ctx *HeaderObjectContext) {}

// EnterHeaderName is called when production headerName is entered.
func (s *BaseSfzListener) EnterHeaderName(ctx *HeaderNameContext) {}

// ExitHeaderName is called when production headerName is exited.
func (s *BaseSfzListener) ExitHeaderName(ctx *HeaderNameContext) {}

// EnterOpcodeStatement is called when production opcodeStatement is entered.
func (s *BaseSfzListener) EnterOpcodeStatement(ctx *OpcodeStatementContext) {}

// ExitOpcodeStatement is called when production opcodeStatement is exited.
func (s *BaseSfzListener) ExitOpcodeStatement(ctx *OpcodeStatementContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseSfzListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseSfzListener) ExitOpcode(ctx *OpcodeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSfzListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSfzListener) ExitValue(ctx *ValueContext) {}
