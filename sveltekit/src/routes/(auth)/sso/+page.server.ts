import { fail } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async ({ url, fetch }) => {
  const ticket = url.searchParams.get('ticket')
  const service = encodeURIComponent("http://localhost:5173/sso")
  const validationUrl = `https://sso.ui.ac.id/cas2/serviceValidate?ticket=${ticket}&service=${service}`
  console.log(validationUrl)
  const res = await fetch(validationUrl)
  if (!res.ok) {
    return {}
  }
  if (res.status === 200) {
    const body = await res.text()
    return {}
  }
  return {}
}