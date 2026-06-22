import type { PageLoad } from "./$types"

export const load: PageLoad = () => {
  const breadcrumbs = [
    {
      name: 'Kelola Kelas',
      href: '/admin/courses'
    }
  ]
  return { breadcrumbs }
}