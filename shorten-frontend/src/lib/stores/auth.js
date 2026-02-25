import { writable } from "svelte/store";
import { browser } from "$app/environment";

export const token = writable(browser ? localStorage.getItem("token") : null);

export function login(newToken) {
  if (browser) localStorage.setItem("token", newToken);
  token.set(newToken);
}

export function logout() {
  if (browser) localStorage.removeItem("token");
  token.set(null);
}
