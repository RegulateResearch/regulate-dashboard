// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			userInfo?: {
				id: number,
				name?: string;
				email?: string;
				role: string;
			};
		}
		interface PageData {
			userInfo?: {
				id: number,
				name: string;
				email: string;
				role: string;
			};
			pageTitle?: string;
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
