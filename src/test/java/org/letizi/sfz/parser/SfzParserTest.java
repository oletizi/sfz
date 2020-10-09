package org.letizi.sfz.parser;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.core.config.Configurator;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

import static java.util.Arrays.stream;
import static org.junit.jupiter.api.Assertions.*;


public class SfzParserTest {
  static {
    Configurator.setRootLevel(Level.INFO);
  }

  private static final Logger logger = LogManager.getLogger(SfzParserTest.class.getName());

  @Test
  public void testBasics() throws Exception {
    logger.atLevel(Level.OFF);
    final String filename = getClass().getClassLoader().getResource("test.sfz").getFile();
    final CharStream in = CharStreams.fromFileName(filename);
    final SfzLexer lexer = new SfzLexer(in);
    final SfzParser parser = new SfzParser(new CommonTokenStream(lexer));
    final TestSfzListener listener = new TestSfzListener();
    parser.addParseListener(listener);
    parser.sfz();

    // test a few headers
    stream(new String[] {
        "global",
        "group",
        "region"
    }).forEach(header -> assertTrue(listener.headers.contains(header)));

    // test a few opcodes
    stream(new String[] {
        "sample",
        "lokey",
        "hikey",
        "seq_length"
    }).forEach(opcode -> assertTrue(listener.opcodes.contains(opcode)));

    // test a few values
    stream(new String[]{
        "-3600",
        "arco\\arco_c1_pp_down.wav",
        "59"
    }).forEach(value-> assertTrue(listener.values.contains(value)));
  }

  private static class TestSfzListener extends SfzBaseListener {

    final Set<String> headers = new HashSet<>();
    final Set<String> opcodes = new HashSet<>();
    final Set<String> values = new HashSet<>();

    private final Logger logger = LogManager.getLogger(TestSfzListener.class.getName());

    public void exitHeader(SfzParser.HeaderContext ctx) {
      final String header = ctx.getText();
      //logger.info("exitHeader: " + header);
      headers.add(header);
    }

    public void exitOpcode(SfzParser.OpcodeContext ctx) {
      final String opcode = ctx.getText();
      //logger.info("exitOpcode: " + opcode);
      opcodes.add(opcode);
    }

    public void exitValue(SfzParser.ValueContext ctx) {
      final String value = ctx.getText();
      //logger.info("exitValue: " + value);
      values.add(value);
    }
  }
}