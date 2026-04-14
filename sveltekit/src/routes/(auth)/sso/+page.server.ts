import type { PageServerLoad } from "./$types"
import { sso } from '$lib/server/api/auth';
import { redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ url }) => {
  const ticket = url.searchParams.get('ticket')
  const service = encodeURIComponent(`${url.origin}/sso`)
  try {
    if (!ticket) throw new Error('No ticket provided')
    await sso({
      ticket,
      service: service
    })
  } catch (err) {
    return {
      message: 'SSO failed: ' + (err instanceof Error ? err.message : 'Unknown error')
    }
  }
  const redirectTo = url.searchParams.get('redirectTo') || 'after-login'
  throw redirect(301, redirectTo)
}