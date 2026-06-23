import { getSession } from '$lib/server/api/auth';
import { AuthorizationError } from '$lib/server/api/errors';
import { redirect } from '@sveltejs/kit';

export const handle = async ({ event, resolve }) => {
  try {
    const res = await getSession();
    if (res.data && res.message === 'token is valid') {
      event.locals.userInfo = { ...res.data }
    }
  } catch (error) {
    if (error instanceof AuthorizationError) {
      throw redirect(302, '/login');
    }
    event.locals.userInfo = undefined
  }

  return resolve(event);
};