FROM docker.io/ubuntu:20.04

RUN apt update && DEBIAN_FRONTEND=noninteractive apt install -y gcc make libtool libhwloc-dev libx11-dev \
    libxt-dev libedit-dev libical-dev ncurses-dev perl \
    postgresql-server-dev-all postgresql-contrib python3-dev tcl-dev tk-dev swig \
    libexpat-dev libssl-dev libxext-dev libxft-dev autoconf \
    automake g++ libcjson-dev wget

RUN wget https://github.com/openpbs/openpbs/archive/refs/tags/v23.06.06.tar.gz && \
    tar xvf v23.06.06.tar.gz && \
    rm v23.06.06.tar.gz

WORKDIR /openpbs-23.06.06

RUN ./autogen.sh && \
    ./configure --prefix=/opt/pbs && \
    make && \
    make install && \
    /opt/pbs/libexec/pbs_postinstall

WORKDIR /

RUN rm -rf /v23.06.06

RUN apt update && apt install -y openssh-server
RUN mkdir -p /var/run/sshd && \
    chmod 0755 /var/run/sshd

RUN echo "root:pbs" | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

COPY testpbs.entrypoint.sh /testpbs.entrypoint.sh
RUN chmod +x /testpbs.entrypoint.sh

ENTRYPOINT ["/testpbs.entrypoint.sh"]