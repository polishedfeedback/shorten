export const ssr = false;

import { browser } from "$app/environment";
import { redirect } from "@sveltejs/kit";

export function load() {
  if (browser) {
    const t = localStorage.getItem("token");
    if (!t) redirect(302, "/");
  }
}
