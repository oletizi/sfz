FROM oletizi/golang-java:latest

ENV ANTLR_JAR=antlr-4.8-complete.jar
ENV ANTLR_JAR_PATH=/usr/local/lib/${ANTLR_JAR}
ENV CLASSPATH=${ANTLR_JAR_PATH}
ENV GOLANG_VERSION=1.15.2
ENV GOLANG_URL=https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
COPY antlr /usr/local/bin/
RUN cd /usr/local && \
    sudo rm -rf /usr/local/go && \
    curl ${GOLANG_URL} | sudo tar xfvz - && \
    sudo chmod +x /usr/local/bin/antlr && \
    sudo curl -o ${ANTLR_JAR_PATH} https://www.antlr.org/download/${ANTLR_JAR}
    #echo "alias antlr='java -Xmx500M org.antlr.v4.Tool'" >> ~/.bashrc
