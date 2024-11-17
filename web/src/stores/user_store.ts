import { writable } from "svelte/store";

export const userStore = writable<App.User | null>(null);
