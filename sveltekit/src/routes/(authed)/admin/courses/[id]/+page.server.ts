import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  const breadcrumbs = [
    {
      name: 'Kelola Kelas',
      href: '/admin/courses'
    },
    {
      name: 'Kelola Detail Kelas',
      href: `/admin/courses/${params.id}`
    },
  ]

  return {
    breadcrumbs,
  };
};