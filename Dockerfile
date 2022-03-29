FROM alpine
WORKDIR /project
COPY . .
EXPOSE 8080
RUN chmod a+rwx proxycheck
CMD ./proxycheck