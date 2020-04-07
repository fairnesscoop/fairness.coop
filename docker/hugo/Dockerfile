FROM alpine:3.11

LABEL Fairness <equipe@fairness.coop>

USER root

# Packages
RUN apk add --no-cache \
# System
      sudo su-exec \
# Command
      bash bash-completion \
# Tools
      curl make git dumb-init


# Entrypoint
COPY entrypoint.sh /usr/local/bin/
RUN chmod 777 /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]

# Shell
CMD ["/bin/bash"]

# Working directory
WORKDIR /srv

########
# User #
########

ARG USER_ID
ARG GROUP_ID

ENV USER_DEFAULT="fairness" \
    USER_ID="${USER_ID:-1000}" \
    USER_SUDO="1" \
    GROUP_DEFAULT="fairness" \
    GROUP_ID="${GROUP_ID:-1000}"

RUN addgroup -g ${GROUP_ID} ${GROUP_DEFAULT} && \
    adduser -D -s /bin/bash -g ${USER_DEFAULT} -u ${USER_ID} -G ${GROUP_DEFAULT} ${USER_DEFAULT}

##########
# Custom #
##########

# Hugo
ENV HUGO_VERSION="0.66.0"
RUN curl -L https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz \
      | tar -zxvf - -C /tmp && \
    mv /tmp/hugo /usr/local/bin/hugo && \
    rm -Rf /tmp/hugo_${HUGO_VERSION}_linux_amd64

# Expose default hugo port
EXPOSE 1313
