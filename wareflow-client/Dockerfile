# Стадия сборки
FROM node:18-alpine AS builder

# Установка рабочей директории
WORKDIR /wareflow-client

# Копирование package.json и package-lock.json
COPY package*.json ./

# Установка зависимостей
RUN npm install

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN npm run build

# Стадия запуска
FROM nginx:alpine

# Копирование результатов сборки из builder-стадии в папку nginx
COPY --from=builder /wareflow-client/dist /usr/share/nginx/html

# Открытие порта 80
EXPOSE 80

# Запуск nginx
CMD ["nginx", "-g", "daemon off;"]