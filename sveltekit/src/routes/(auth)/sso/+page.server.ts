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
      pageTitle: 'Single Sign-on Universitas Indonesia',
      error: err instanceof Error ? err.message : 'Unknown error'
    }
  }
  throw redirect(303, '/')
}