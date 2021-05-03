FROM ubuntu:20.04
RUN apt-get update ;\
    apt-get install -y tzdata
RUN apt-get update ; \
    apt-get install -y apt-transport-https ca-certificates
COPY contact ./contact
COPY template.html ./template.html
COPY settings.yaml ./settings.yaml
CMD ["./contact"]
