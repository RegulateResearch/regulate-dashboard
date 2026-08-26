import { getRequestEvent } from "$app/server";
import { env } from "$env/dynamic/public";
import { typedFetch } from "$lib/server/utils";
import { redirect } from "@sveltejs/kit";
import * as Schema from "./schema";

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
  cookies.set("loginMethod", 'EMAIL', {
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
    sameSite: 'lax',
    maxAge: 60 * 60 * 24
  })
  cookies.set("loginMethod", 'SSO_UI', {
    path: '/',
    httpOnly: true,
    secure: true,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
}

export const logout = () => {
  const { url, cookies } = getRequestEvent()
  const isSSO = cookies.get('loginMethod') === 'SSO_UI'
  cookies.delete('authToken', { path: '/' })
  cookies.delete('loginMethod', { path: '/' })
  if (isSSO) {
    throw redirect(303, `${env.PUBLIC_SSO_UI_URL}/logout?url=${encodeURIComponent(`${url.origin}/login`)}`)
  } else {
    throw redirect(303, '/login')
  }
}