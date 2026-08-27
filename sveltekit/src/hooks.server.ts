import { getSession, logout } from '$lib/server/api/auth';
import { AuthorizationError } from '$lib/server/errors';

export const handle = async ({ event, resolve }) => {
  try {
    const res = await getSession();
    if (res.data && res.message === 'token is valid') {
      event.locals.session = { ...res.data }
    }
  } catch (error) {
    if (error instanceof AuthorizationError) {
      logout();
    }
    event.locals.session = undefined
  }

  return resolve(event);
};