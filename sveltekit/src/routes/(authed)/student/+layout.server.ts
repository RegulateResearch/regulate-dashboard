import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  const role = locals.userInfo?.role
  switch (role) {
    case 'admin':
      throw redirect(307, '/admin');
    case 'lecturer':
      throw redirect(307, '/lecturer');
    case 'student':
      break;
    default:
      break;
  }
};
