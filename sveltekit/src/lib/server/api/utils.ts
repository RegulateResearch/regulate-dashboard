import { getRequestEvent } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import { ApiError, AuthorizationError, NetworkError } from '$lib/server/api/errors';
import { ZodError, type z } from 'zod';

export const typedFetch = async <TResponse extends z.ZodTypeAny, TBody extends z.ZodTypeAny>(
  endpoint: string,
  responseSchema: TResponse,
  options?: Omit<RequestInit, 'body'> & {
    requireAuthentication: boolean,
    body?: z.infer<TBody>,
    bodySchema?: TBody,
    authFallback?: string
  }
): Promise<z.infer<TResponse>> => {
  const { fetch, cookies } = getRequestEvent();
  try {
    let requestBody: string | undefined;

    if (options?.body && options?.bodySchema) {
      const parsedBody = options.bodySchema.parse(options.body);
      requestBody = JSON.stringify(parsedBody);
    }

    let requestHeaders: RequestInit["headers"] = {
      'Content-Type': 'application/json',
      ...options?.headers
    }
    if (requestHeaders && options?.requireAuthentication) {
      const authToken = cookies.get('authToken')
      if (!authToken) throw AuthorizationError
      requestHeaders = {
        ...requestHeaders,
        'Authorization': `Bearer ${authToken}`,
      }
    }

    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      body: requestBody,
      headers: requestHeaders
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ message: response.statusText }));
      throw new ApiError(
        errorData.message || 'Request failed',
        response.status,
        errorData
      );
    }

    const responseData: unknown = await response.json();
    return responseSchema.parse(responseData);
  } catch (error) {
    if (error instanceof ApiError || error instanceof ZodError) {
      throw error;
    } else if (error instanceof Error) {
      throw new NetworkError(error.message);
    } else {
      throw new Error("An unexpected error occurred");
    }
  }
}