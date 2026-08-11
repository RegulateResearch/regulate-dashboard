import type { PageLoad } from "./$types"

export const load: PageLoad = () => {
  const breadcrumbs = [
    {
      name: 'Kelola Pengguna',
      href: '/admin/users'
    }
  ]
  return { breadcrumbs }
}