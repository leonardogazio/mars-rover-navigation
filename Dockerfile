FROM alpine:3

RUN apk update \
	&& apk -U upgrade \
	&& apk add --no-cache ca-certificates bash \
	&& update-ca-certificates --fresh \
	&& rm -rf /var/cache/apk/*

# adds app user and fix app folder's permission
RUN	addgroup leogazio \
	&& adduser -S leogazio -u 1000 -G leogazio

USER leogazio

COPY --chown=leogazio:leogazio app /usr/local/bin/
RUN chmod +x /usr/local/bin/app

ENTRYPOINT [ "/usr/local/bin/app" ]