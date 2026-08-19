import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  const breadcrumbs = [
    {
      name: 'Kelola Pengguna',
      href: '/admin/users'
    },
    {
      name: 'Kelola Detail Pengguna',
      href: `/admin/users/${params.id}`
    },
  ]

  return {
    breadcrumbs,
  };
};