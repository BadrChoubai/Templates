/**
 * Minimal fetch wrapper for the API. Requests go to /api by default, which the
 * dev server (vite.config.ts) and the Docker image (nginx.conf.template)
 * forward to the API. Set VITE_API_BASE_URL at build time to call the API
 * directly instead; it then has to allow this site's origin (CORS).
 */
const baseUrl = (import.meta.env.VITE_API_BASE_URL ?? "/api").replace(
  /\/$/,
  "",
);

/** An RFC 9457 Problem Details body, as both API starters return on errors */
export interface ProblemDetails {
  type?: string;
  title?: string;
  status?: number;
  detail?: string;
  /** Validation messages keyed by request field, e.g. { name: [...] } */
  errors?: Record<string, string[]>;
}

/** Thrown for any non-2xx response, or when the API can't be reached */
export class ApiError extends Error {
  readonly status: number;
  readonly problem: ProblemDetails;

  constructor(status: number, problem: ProblemDetails) {
    super(problem.detail ?? problem.title ?? `Request failed (${status})`);
    this.name = "ApiError";
    this.status = status;
    this.problem = problem;
  }

  /** Validation messages for one request field, if there are any */
  fieldErrors(field: string): string[] {
    return this.problem.errors?.[field] ?? [];
  }
}

export async function request<T>(
  path: string,
  init: RequestInit = {},
): Promise<T> {
  const headers = new Headers(init.headers);
  headers.set("Accept", "application/json");
  if (init.body !== undefined) headers.set("Content-Type", "application/json");

  let response: Response;
  try {
    response = await fetch(`${baseUrl}${path}`, { ...init, headers });
  } catch {
    throw new ApiError(0, {
      title: "Network error",
      detail: "Couldn't reach the API. Is it running?",
    });
  }

  if (!response.ok) {
    throw new ApiError(response.status, await readProblem(response));
  }
  return (await response.json()) as T;
}

async function readProblem(response: Response): Promise<ProblemDetails> {
  try {
    return (await response.json()) as ProblemDetails;
  } catch {
    // Not JSON, e.g. an HTML error page from a proxy
    return { status: response.status, title: response.statusText };
  }
}
