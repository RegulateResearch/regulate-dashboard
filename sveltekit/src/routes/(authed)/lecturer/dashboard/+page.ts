import type { PageLoad } from "../$types"

export const load: PageLoad = () => {
  const breadcrumbs = [
    {
      name: 'Dashboard',
      href: '/lecturer/dashboard'
    }
  ]
  return { breadcrumbs }
}