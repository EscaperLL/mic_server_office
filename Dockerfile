FROM alpine
ADD mic_srv_office-service /mic_srv_office-service
ENTRYPOINT [ "/mic_srv_office-service" ]
