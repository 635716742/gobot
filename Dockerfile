FROM alpine

COPY ./bot_linux /home/bot/
COPY ./script/* /home/bot/script/

WORKDIR /home/bot

# RUN go env -w GOPROXY=https://goproxy.cn

EXPOSE 8888

RUN ls
RUN pwd

ENTRYPOINT ./bot_linux $0 $@