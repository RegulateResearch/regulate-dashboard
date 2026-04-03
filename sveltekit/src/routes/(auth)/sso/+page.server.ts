import { fail } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"
import { message } from "valibot"

export const load: PageServerLoad = async ({ url, fetch }) => {
  const ticket = url.searchParams.get('ticket')
  const service = encodeURIComponent("http://localhost:5173/")
  const validationUrl = `https://sso.ui.ac.id/cas2/serviceValidate?ticket=${ticket}&service=${service}`
  const res = await fetch(validationUrl)
  if (res.status === 200) {
    const body = await res.text()
    console.log(body)
    return { message: body }
  }
  return {}
}