import type { PageLoad } from "../$types"

export const load: PageLoad = () => {
  const breadcrumbs = [
    {
      name: 'Dashboard',
      href: '/student/dashboard'
    }
  ]
  return { breadcrumbs }
}