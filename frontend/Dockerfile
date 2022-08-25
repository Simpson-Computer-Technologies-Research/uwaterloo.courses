# RUNNING SVELTE FRONTEND
FROM node:16-alpine AS frontend
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

# RUN THE APPLICATION
EXPOSE 8080
CMD ["npm", "start"]