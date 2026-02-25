import { k as store_get, d as attr, f as escape_html, l as unsubscribe_stores } from './index-8mcHi-l6.js';
import { t as theme } from './theme-DFvMdRXf.js';
import './root-CU6M0b66.js';
import './state.svelte-ahRM4L_a.js';

function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    var $$store_subs;
    let username = "";
    let password = "";
    let loading = false;
    $$renderer2.push(`<div class="min-h-screen flex items-center justify-center p-4" style="background: var(--bg); color: var(--text);"><div class="fixed top-4 right-4"><button class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all" style="border: 2px solid var(--border); color: var(--text);">`);
    if (store_get($$store_subs ??= {}, "$theme", theme) === "dark") {
      $$renderer2.push("<!--[-->");
      $$renderer2.push(`☀️`);
    } else {
      $$renderer2.push("<!--[!-->");
      $$renderer2.push(`🌙`);
    }
    $$renderer2.push(`<!--]--></button></div> <div class="w-full max-w-md"><div class="mb-8"><p class="font-mono text-xs font-bold tracking-widest uppercase text-orange-500 mb-2">// Link Shortener</p> <h1 class="font-mono text-4xl font-bold leading-none">Sign<br/> <span style="-webkit-text-stroke: 2px var(--text); color: transparent;">In</span></h1></div> <div class="p-8" style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);">`);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--> <form class="flex flex-col gap-5"><div class="flex flex-col gap-1.5"><label for="username" class="font-mono text-xs font-bold uppercase tracking-widest" style="color: var(--text-muted);">Username</label> <input id="username" type="text"${attr("value", username)} placeholder="your username" required="" class="p-3 font-sans text-sm outline-none transition-all" style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"/></div> <div class="flex flex-col gap-1.5"><label for="password" class="font-mono text-xs font-bold uppercase tracking-widest" style="color: var(--text-muted);">Password</label> <input id="password" type="password"${attr("value", password)} placeholder="your password" required="" class="p-3 font-sans text-sm outline-none transition-all" style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"/></div> <button type="submit"${attr("disabled", loading, true)} class="bg-orange-500 text-white font-mono font-bold text-sm uppercase tracking-widest py-3 px-6 transition-all disabled:opacity-50 disabled:cursor-not-allowed" style="border: 2.5px solid var(--border); box-shadow: 4px 4px 0 var(--border);">${escape_html("Sign In →")}</button></form></div> <p class="font-mono text-xs text-center mt-4" style="color: var(--text-muted);">shorten.krishnasaimarupudi.com</p></div></div>`);
    if ($$store_subs) unsubscribe_stores($$store_subs);
  });
}

export { _page as default };
//# sourceMappingURL=_page.svelte-C1QqFwjb.js.map
