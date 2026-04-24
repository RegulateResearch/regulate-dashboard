import { getSession, logout } from '$lib/server/api/auth';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async () => {
  try {
    const session = await getSession()
    if (session && session.message === 'token is valid') {
      return { userInfo: session.data }
    }
  } catch {
    return logout()
  }
}