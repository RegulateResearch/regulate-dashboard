import { getSession } from '$lib/server/api/auth';
import type { LayoutServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load: LayoutServerLoad = async ({ url, cookies }) => {
  try {
    const session = await getSession()
    if (session && session.message) {
      return { status: 'success' }
    }
  } catch {
    cookies.delete('authToken', {
      path: '/'
    })
    throw redirect(301, `/login?redirectTo=${encodeURIComponent(url.pathname)}`)
  }
}