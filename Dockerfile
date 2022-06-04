# Build Webpags
FROM node:17.12.0-alpine as build_web
ADD web .
RUN npm run typecheck && cross-env VITE_ENV_TYPE=prod vite build
# Release Packages
FROM nginx:1.17.10-alpine as release
ADD web /usr/share/nginx/html
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
