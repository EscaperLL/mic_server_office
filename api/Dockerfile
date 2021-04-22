FROM alpine
COPY  /bin/api_server /opt/mic_src/api/api_server
WORKDIR /opt/mic_src/api/
RUN chmod +x api_server
EXPOSE 8080
ENTRYPOINT [ "./api_server" ,"--registry=etcd","--registry_address=etcd:2379"]

