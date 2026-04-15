import { getRequestEvent } from "$app/server";
import { typedFetch } from "$lib/server/api/utils";
import * as Schema from "./schema"

export const getSession = async () => {
  return await typedFetch(
    '/auth/session',
    Schema.sessionResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const login = async (body: Schema.LoginRequest) => {
  const { cookies } = getRequestEvent()
  const response = await typedFetch(
    '/auth/login',
    Schema.loginResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.loginRequestSchema,
      requireAuthentication: false
    })
  cookies.set("authToken", response.data, {
    path: '/',
    httpOnly: true,
    secure: true,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
}

export const register = async (body: Schema.RegisterRequest) => {
  return await typedFetch(
    '/auth/register',
    Schema.registerResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.registerRequestSchema,
      requireAuthentication: false
    })
}

export const sso = async (body: Schema.SSORequest) => {
  const { cookies } = getRequestEvent()
  const response = await typedFetch(
    '/auth/sso',
    Schema.ssoResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.ssoRequestSchema,
      requireAuthentication: false
    })
  cookies.set("authToken", response.data, {
    path: '/',
    httpOnly: true,
    secure: true,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
}