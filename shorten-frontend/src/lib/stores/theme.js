import { writable } from "svelte/store";
import { browser } from "$app/environment";

const getInitial = () => {
  if (!browser) return "light";
  return (
    localStorage.getItem("theme") ||
    (window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light")
  );
};

export const theme = writable(getInitial());

export function toggleTheme() {
  theme.update((t) => {
    const next = t === "dark" ? "light" : "dark";
    if (browser) {
      localStorage.setItem("theme", next);
      document.documentElement.setAttribute("data-theme", next);
    }
    return next;
  });
}
