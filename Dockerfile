FROM alpine
WORKDIR /project
COPY proxycheck .
EXPOSE 8080
RUN ls -la
RUN chmod +x proxycheck
CMD ["./proxycheck"]