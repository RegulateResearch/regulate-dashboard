import { logout } from '$lib/server/api/auth';
import type { Actions } from './$types';

export const actions: Actions = {
  default: () => {
    logout()
  }
} satisfies Actions;