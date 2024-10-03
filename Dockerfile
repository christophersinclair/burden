FROM scratch

WORKDIR /app

COPY --chmod=755 ./burden .

CMD ["./burden"]