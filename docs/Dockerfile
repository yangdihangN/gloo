FROM nginx:1.13.9-alpine

ARG VERSION

# Replace existing NGINX configuration
RUN rm -rf /etc/nginx/conf.d
COPY conf /etc/nginx

# Copy over both regular an no_auth bundles
COPY ./site /usr/share/nginx/html/no_auth/$VERSION

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]

