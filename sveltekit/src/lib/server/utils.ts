import { getRequestEvent } from '$app/server';
import { env } from '$env/dynamic/private';
import { ApiError, AuthorizationError } from '$lib/server/errors';
import { ZodError, type z } from 'zod';

export const typedFetch = async <TResponse extends z.ZodTypeAny, TBody extends z.ZodTypeAny>(
  endpoint: string,
  responseSchema: TResponse,
  options?: Omit<RequestInit, 'body'> & {
    requireAuthentication: boolean,
    body?: z.infer<TBody>,
    bodySchema?: TBody
  }
): Promise<z.infer<TResponse>> => {
  const serverLogBaseString = `Request ${options?.method || 'GET'} to ${env.API_BASE_URL}${endpoint}`

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

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ message: response.statusText }));
      if (response.status === 401) {
        throw new AuthorizationError(errorData.message || 'Request failed')
      }
      throw new ApiError(
        errorData.message || 'Request failed',
        response.status,
        errorData
      );
    }

    const responseData: unknown = await response.json();

    serverLog({
      message: `
[FETCH] ${serverLogBaseString} returned status ${response.status}
Request Body: ${requestBody || "-"}
Response Body: ${responseData ? JSON.stringify(responseData, null, 2) : "-"}
      `,
      isError: false,
      devOnly: true
    })

    return responseSchema.parse(responseData);
  } catch (error) {
    let message = 'An unexpected error occurred';
    let devOnly = true;
    if (error instanceof ApiError) {
      message = `[ERROR] ${serverLogBaseString} ApiError: ${error.message} (Status: ${error.statusCode}) (Data: ${JSON.stringify(error.details)})`
    } else if (error instanceof AuthorizationError) {
      message = `[ERROR] ${serverLogBaseString} AuthorizationError: ${error.message}`
    } else if (error instanceof ZodError) {
      message = `[ERROR] ${serverLogBaseString} ValidationError: ${error.message}`
    } else if (error instanceof Error) {
      message = `[ERROR] ${serverLogBaseString} ${error.name}: ${error.message}`
      devOnly = false;
    } else {
      message = `[ERROR] ${serverLogBaseString} Unknown error: An unexpected error occurred ${error}`
      devOnly = false;
    }
    serverLog({
      message,
      isError: true,
      devOnly
    })
    throw error;
  }
}

export const serverLog = ({ message, isError, devOnly }: { message: string, isError: boolean, devOnly?: boolean }) => {
  if (devOnly && env.NODE_ENV === 'production') {
    return;
  }
  if (isError) {
    console.error(message);
  } else {
    console.log(message);
  }
}