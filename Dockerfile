FROM ubuntu:18.04

RUN apt-get update && apt-get install -y openssh-server \
    sudo \
    vim \
    git \
    curl

COPY certificates/dev-certificate.pub /authorized_keys
RUN mkdir -p ~root/.ssh /var/run/sshd \
	&& chmod 700 ~root/.ssh \
	&& mv /authorized_keys ~root/.ssh/authorized_keys \
	&& chmod 600 ~root/.ssh/authorized_keys

# RUN echo "alias python=/usr/bin/python3" >> /etc/profile

COPY ./src/ /root/go/src/ftpserver
WORKDIR /root/go/src/ftpserver

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]