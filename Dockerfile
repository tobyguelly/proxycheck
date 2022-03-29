FROM alpine
WORKDIR /project
COPY proxycheck .
EXPOSE 8080
RUN chmod +x proxycheck
CMD ls
CMD ./proxycheck