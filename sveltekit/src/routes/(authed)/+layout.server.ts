import type { LayoutServerLoad } from './$types';
import { redirect, type Actions } from '@sveltejs/kit';

export const load: LayoutServerLoad = async ({ url, cookies }) => {
  const authToken = cookies.get('authToken')
  if (!authToken) throw redirect(301, `/login?redirectTo=${encodeURIComponent(url.pathname)}`)
  return { status: 'success' }
}