# Стадия сборки
FROM node:18-alpine AS builder

# Установка рабочей директории
WORKDIR /wareflow-client

# Копирование package.json и package-lock.json
COPY package*.json ./

# Установка зависимостей
RUN npm install --legacy-peer-deps

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN npm run build

# Стадия запуска
FROM nginx:alpine

# Открытие порта 8080 для контейнера
EXPOSE 80

# Копирование результатов сборки из builder-стадии в папку nginx
COPY --from=builder /wareflow-client/dist /usr/share/nginx/html

# Запуск nginx с настройкой для порта 8080
CMD ["nginx", "-g", "daemon off;"]
