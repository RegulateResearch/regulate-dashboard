import type { PageServerLoad } from "./$types";
import { getUsers } from "$lib/server/api/admin/user";

export const load: PageServerLoad = async () => {
  const breadcrumbs = [
    {
      name: 'Kelola Pengguna',
      href: '/admin/users'
    }
  ]
  const users = await getUsers();

  return {
    breadcrumbs,
    users: users.data,
  };
};