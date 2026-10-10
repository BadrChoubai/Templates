FROM node:22-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

# Copy only the necessary source files
COPY tsconfig.json ./
COPY src ./src

RUN mkdir -p /app/.db
RUN npm run database:seed

RUN npm run build

RUN rm tsconfig.json && rm -rf src

FROM node:22-alpine AS runner

WORKDIR /app

COPY --from=builder /app/.db .db
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./
COPY --from=builder /app/package-lock.json ./

RUN npm ci --omit=dev

EXPOSE 3000

CMD ["node", "dist/index.js"]
