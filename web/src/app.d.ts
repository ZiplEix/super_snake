// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
		interface User {
			id: number;
			email: string;
			name: string;
		}

		interface GameParamsRequest {
			nbPlayerMax: number;
			mapHeight: number;
			mapWidth: number;
		}

		interface Session {
			user?: User;
		}
	}
}

export {};
