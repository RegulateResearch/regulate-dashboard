// See https://svelte.dev/docs/kit/types#app.d.ts

import type { UserWithId } from "$lib/server/schema";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			session?: {
				id: number,
				role: "admin" | "user";
			},
			userInfo?: UserWithId;
		}
		interface PageData {
			userInfo?: UserWithId;
			pageTitle?: string;
			breadcrumbs?: {
				name: string,
				href: string
			}[]
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
