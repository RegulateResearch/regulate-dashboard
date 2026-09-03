import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  if (locals.userInfo?.role !== 'admin') {
    if (locals.userInfo?.academicRole === 'lecturer') throw redirect(307, '/lecturer');
    throw redirect(307, '/student');
  }
};
