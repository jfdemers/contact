FROM ubuntu
RUN apt-get update ; \
    apt-get install -y apt-transport-https
COPY contact ./contact
COPY template.html ./template.html
COPY settings.yaml ./settings.yaml
CMD ["./contact"]
