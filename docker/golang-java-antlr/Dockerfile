FROM oletizi/golang-java:latest

ENV ANTLR_JAR=antlr-4.5-complete.jar
ENV ANTLR_JAR_PATH=/usr/local/lib/${ANTLR_JAR}
ENV CLASSPATH=${ANTLR_JAR_PATH}
RUN sudo curl -o ${ANTLR_JAR_PATH} https://www.antlr.org/download/${ANTLR_JAR} && \
    echo "alias antlr='java -Xmx500M org.antlr.v4.Tool'" >> ~/.bashrc
