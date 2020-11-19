FROM centos:7.9.2009
WORKDIR /app
ADD ./simpleecho ./simpleecho
RUN ["chmod", "+x", "/app/simpleecho"]
EXPOSE 8080
ENTRYPOINT ["./simpleecho"]
CMD ["./simpleecho"]
