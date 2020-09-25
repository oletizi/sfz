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
	 * Enter a parse tree produced by {@link SfzParser#line}.
	 * @param ctx the parse tree
	 */
	void enterLine(SfzParser.LineContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#line}.
	 * @param ctx the parse tree
	 */
	void exitLine(SfzParser.LineContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#header_stmt}.
	 * @param ctx the parse tree
	 */
	void enterHeader_stmt(SfzParser.Header_stmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#header_stmt}.
	 * @param ctx the parse tree
	 */
	void exitHeader_stmt(SfzParser.Header_stmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#header}.
	 * @param ctx the parse tree
	 */
	void enterHeader(SfzParser.HeaderContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#header}.
	 * @param ctx the parse tree
	 */
	void exitHeader(SfzParser.HeaderContext ctx);
	/**
	 * Enter a parse tree produced by {@link SfzParser#opcode_stmt}.
	 * @param ctx the parse tree
	 */
	void enterOpcode_stmt(SfzParser.Opcode_stmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SfzParser#opcode_stmt}.
	 * @param ctx the parse tree
	 */
	void exitOpcode_stmt(SfzParser.Opcode_stmtContext ctx);
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