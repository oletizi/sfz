// Generated from Sfz.g4 by ANTLR 4.8
package org.letizi.sfz.parser;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SfzParser}.
 */
public interface SfzListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SfzParser#sfz}.
	 * @param ctx the parse tree
	 */
	void enterSfz(SfzParser.SfzContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#sfz}.
	 * @param ctx the parse tree
	 */
	void exitSfz(SfzParser.SfzContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#sfzObject}.
	 * @param ctx the parse tree
	 */
	void enterSfzObject(SfzParser.SfzObjectContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#sfzObject}.
	 * @param ctx the parse tree
	 */
	void exitSfzObject(SfzParser.SfzObjectContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#headerObject}.
	 * @param ctx the parse tree
	 */
	void enterHeaderObject(SfzParser.HeaderObjectContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#headerObject}.
	 * @param ctx the parse tree
	 */
	void exitHeaderObject(SfzParser.HeaderObjectContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#headerName}.
	 * @param ctx the parse tree
	 */
	void enterHeaderName(SfzParser.HeaderNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#headerName}.
	 * @param ctx the parse tree
	 */
	void exitHeaderName(SfzParser.HeaderNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#opcodeStatement}.
	 * @param ctx the parse tree
	 */
	void enterOpcodeStatement(SfzParser.OpcodeStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#opcodeStatement}.
	 * @param ctx the parse tree
	 */
	void exitOpcodeStatement(SfzParser.OpcodeStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#opcode}.
	 * @param ctx the parse tree
	 */
	void enterOpcode(SfzParser.OpcodeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#opcode}.
	 * @param ctx the parse tree
	 */
	void exitOpcode(SfzParser.OpcodeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(SfzParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(SfzParser.ValueContext ctx);
}