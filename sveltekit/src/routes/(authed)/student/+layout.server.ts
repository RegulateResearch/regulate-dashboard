import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  if (locals.session?.role === 'admin') throw redirect(307, '/admin');
  if (locals.userInfo?.academicRole === 'lecturer') throw redirect(307, '/lecturer');
};
