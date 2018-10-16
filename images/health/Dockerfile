FROM scratch
ADD heatlh/ca-certificates.crt /etc/ssl/certs/
ADD heatlh/healthsrv /healthsrv
EXPOSE 8080
CMD ["/healthsrv"]
