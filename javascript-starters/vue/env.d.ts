/// <reference types="vite/client" />

interface ImportMetaEnv {
  /** Prefix for API requests; defaults to /api (see src/api/client.ts) */
  readonly VITE_API_BASE_URL?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
