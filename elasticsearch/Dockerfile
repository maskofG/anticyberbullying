FROM ubuntu:latest

ENV ES_PKG_NAME elasticsearch-1.5.0

RUN \
  apt-get update && \
  apt-get install -y openjdk-7-jdk && \
  rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME /usr/lib/jvm/java-7-openjdk-amd64

#RUN apt-get install wget
RUN apt-get install tar
# Install Elasticsearch.
RUN cd /
ADD https://download.elasticsearch.org/elasticsearch/elasticsearch/$ES_PKG_NAME.tar.gz /
RUN tar xvzf $ES_PKG_NAME.tar.gz && \
  rm -f $ES_PKG_NAME.tar.gz && \
  mv /$ES_PKG_NAME /elasticsearch

# Define mountable directories.
VOLUME ["/data"]

# Mount elasticsearch.yml config
ADD config/elasticsearch.yml /elasticsearch/config/elasticsearch.yml

# Define working directory.
WORKDIR /data

# Define default command.
#CMD ["/elasticsearch/bin/elasticsearch"]

# Expose ports.
#   - 9200: HTTP
#   - 9300: transport
EXPOSE 9200
EXPOSE 9300

ENTRYPOINT /elasticsearch/bin/elasticsearch
