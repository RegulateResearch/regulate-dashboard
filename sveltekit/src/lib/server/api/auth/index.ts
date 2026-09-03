import { getRequestEvent } from "$app/server";
import { env } from "$env/dynamic/public";
import { dev } from '$app/environment';
import { serverLog, typedFetch } from "$lib/server/utils";
import { redirect } from "@sveltejs/kit";
import * as Schema from "./schema";

export const getSession = async (clientUrl?: string) => {
  return await typedFetch(
    '/auth/session',
    Schema.sessionResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true,
      clientUrl
    })
}

export const login = async (body: Schema.LoginRequest, clientUrl?: string) => {
  const { cookies } = getRequestEvent()
  const response = await typedFetch(
    '/auth/login',
    Schema.loginResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.loginRequestSchema,
      requireAuthentication: false,
      clientUrl
    })
  cookies.set("authToken", response.data, {
    path: '/',
    httpOnly: true,
    secure: !dev,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
  cookies.set("loginMethod", 'EMAIL', {
    path: '/',
    httpOnly: true,
    secure: !dev,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
}

export const register = async (body: Schema.RegisterRequest, clientUrl?: string) => {
  return await typedFetch(
    '/auth/register',
    Schema.registerResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.registerRequestSchema,
      requireAuthentication: false,
      clientUrl
    })
}

export const sso = async (body: Schema.SSORequest, clientUrl?: string) => {
  const { cookies } = getRequestEvent()
  const response = await typedFetch(
    '/auth/sso',
    Schema.ssoResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.ssoRequestSchema,
      requireAuthentication: false,
      clientUrl
    })
  cookies.set("authToken", response.data, {
    path: '/',
    httpOnly: true,
    secure: !dev,
    sameSite: 'lax',
    maxAge: 60 * 60 * 24
  })
  cookies.set("loginMethod", 'SSO_UI', {
    path: '/',
    httpOnly: true,
    secure: !dev,
    sameSite: 'strict',
    maxAge: 60 * 60 * 24
  })
}

export const logout = (clientUrl?: string) => {
  const { url, cookies } = getRequestEvent()
  const isSSO = cookies.get('loginMethod') === 'SSO_UI'
  serverLog({
    message: `
[LOGOUT] User logged out, isSSO: ${isSSO}, redirecting to ${isSSO ? `${env.PUBLIC_SSO_UI_URL}/logout?url=${encodeURIComponent(`${url.origin}/login`)}` : '/login'}
Request client URL: ${clientUrl || "-"}`,
    isError: false,
    devOnly: true
  })
  cookies.delete('authToken', { path: '/' })
  cookies.delete('loginMethod', { path: '/' })
  cookies.delete('userInfo', { path: '/' })
  if (isSSO) {
    throw redirect(303, `${env.PUBLIC_SSO_UI_URL}/logout?url=${encodeURIComponent(`${url.origin}/login`)}`)
  } else {
    throw redirect(303, '/login')
  }
}