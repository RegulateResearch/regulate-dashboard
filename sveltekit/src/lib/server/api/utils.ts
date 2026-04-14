import { getRequestEvent } from '$app/server';
import { env } from '$env/dynamic/private';
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

    const response = await fetch(`${env.API_BASE_URL}${endpoint}`, {
      ...options,
      body: requestBody,
      headers: requestHeaders
    });

    serverLog({
      message: `[FETCH] Request to ${env.API_BASE_URL}${endpoint} returned status ${response.status}`,
      isError: false
    })

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
    if (error instanceof ApiError) {
      serverLog({
        message: `[ERROR] ApiError: ${error.message} (Status: ${error.statusCode}) (Data: ${JSON.stringify(error.details)})`,
        isError: true
      })
      throw error;
    } else if (error instanceof AuthorizationError) {
      serverLog({
        message: `[ERROR] AuthorizationError: ${error.message}`,
        isError: true
      })
      throw error
    } else if (error instanceof ZodError) {
      serverLog({
        message: `[ERROR] ValidationError: ${error.message}`,
        isError: true
      })
      throw error
    } else if (error instanceof Error) {
      serverLog({
        message: `[ERROR] NetworkError: ${error.message}`,
        isError: true
      })
      throw new NetworkError(error.message);
    } else {
      serverLog({
        message: `[ERROR] Unknown error: An unexpected error occurred`,
        isError: true
      })
      throw new Error("An unexpected error occurred");
    }
  }
}

export const serverLog = ({ message, isError }: { message: string, isError: boolean }) => {
  if (isError) {
    console.error(message);
  } else {
    console.log(message);
  }
}