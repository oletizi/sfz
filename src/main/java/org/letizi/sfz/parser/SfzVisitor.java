// Generated from src/main/java/org/letizi/sfz/parser/Sfz.g4 by ANTLR 4.8
package org.letizi.sfz.parser;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link SfzParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface SfzVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link SfzParser#sfz}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSfz(SfzParser.SfzContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#line}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLine(SfzParser.LineContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#header_stmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitHeader_stmt(SfzParser.Header_stmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#header}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitHeader(SfzParser.HeaderContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#opcode_stmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOpcode_stmt(SfzParser.Opcode_stmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#opcode}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOpcode(SfzParser.OpcodeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SfzParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValue(SfzParser.ValueContext ctx);
}