import { getRequestEvent } from "$app/server";
import { typedFetch } from "$lib/server/api/utils";
import * as Schema from "./schema"

export const getSession = async () => {
  try {
    const response = typedFetch(
      '/auth/session',
      Schema.sessionResponseSchema,
      {
        method: 'GET',
        requireAuthentication: false
      })
    return response
  } catch (error) {
    throw error
  }
}

export const login = async (body: Schema.LoginRequest) => {
  try {
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
  } catch (error) {
    throw error
  }
}

export const register = async (body: Schema.RegisterRequest) => {
  try {
    return await typedFetch(
      '/auth/register',
      Schema.registerResponseSchema,
      {
        method: 'POST',
        body,
        bodySchema: Schema.registerRequestSchema,
        requireAuthentication: false
      })
  } catch (error) {
    throw error
  }
}

export const sso = async (body: Schema.SSORequest) => {
  try {
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
  } catch (error) {
    throw error
  }
}