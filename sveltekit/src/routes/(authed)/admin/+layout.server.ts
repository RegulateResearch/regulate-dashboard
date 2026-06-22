import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  const role = locals.userInfo?.role
  switch (role) {
    case 'admin':
      break;
    case 'lecturer':
      throw redirect(307, '/lecturer');
    case 'student':
      throw redirect(307, '/student');
    default:
      throw redirect(307, '/student');
  }
};
