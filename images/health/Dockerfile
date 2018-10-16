FROM scratch
ADD health/ca-certificates.crt /etc/ssl/certs/
ADD health/healthsrv /healthsrv
EXPOSE 8080
CMD ["/healthsrv"]
