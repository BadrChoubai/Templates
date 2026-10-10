FROM node:24-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

# Copy only the necessary source files
COPY tsconfig.json ./
COPY src ./src

RUN npm run build

# Create and seed the example SQLite database
RUN npm run database:seed

FROM node:24-alpine AS runner

ENV NODE_ENV=production

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci --omit=dev

COPY --from=builder /app/dist ./dist
COPY --from=builder --chown=node:node /app/.db ./.db

USER node

EXPOSE 8082

CMD ["node", "dist/index.js"]
