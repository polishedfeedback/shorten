import { k as store_get, d as attr, f as escape_html, m as ensure_array_like, n as stringify, o as attr_style, l as unsubscribe_stores } from './index-8mcHi-l6.js';
import './root-CU6M0b66.js';
import './state.svelte-ahRM4L_a.js';
import { t as theme } from './theme-DFvMdRXf.js';

const BASE_URL = "https://shorten.krishnasaimarupudi.com";
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    var $$store_subs;
    let copied = null;
    let links = [];
    let newURL = "";
    let creating = false;
    $$renderer2.push(`<div class="min-h-screen" style="background: var(--bg); color: var(--text);"><nav class="sticky top-0 z-10" style="background: var(--bg); border-bottom: 2.5px solid var(--border);"><div class="max-w-5xl mx-auto px-6 h-16 flex items-center justify-between"><span class="font-mono text-xl font-bold">KSM<span class="text-orange-500">.</span>shorten</span> <div class="flex items-center gap-3"><button class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all" style="border: 2px solid var(--border); color: var(--text);">`);
    if (store_get($$store_subs ??= {}, "$theme", theme) === "dark") {
      $$renderer2.push("<!--[-->");
      $$renderer2.push(`☀️`);
    } else {
      $$renderer2.push("<!--[!-->");
      $$renderer2.push(`🌙`);
    }
    $$renderer2.push(`<!--]--></button> <button class="font-mono text-xs font-bold uppercase tracking-widest px-4 py-2 transition-all" style="border: 2px solid var(--border); color: var(--text);">Logout</button></div></div></nav> <div class="max-w-5xl mx-auto px-6 py-12"><div class="mb-10"><p class="font-mono text-xs font-bold tracking-widest uppercase text-orange-500 mb-2">// Dashboard</p> <h1 class="font-mono text-5xl font-bold leading-none">Your<br/> <span style="-webkit-text-stroke: 2px var(--text); color: transparent;">Links</span></h1></div> <div class="p-6 mb-10" style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"><h2 class="font-mono text-sm font-bold uppercase tracking-widest mb-4" style="color: var(--text);">// New Link</h2> <div class="flex gap-3"><input type="url"${attr("value", newURL)} placeholder="https://example.com" class="flex-1 p-3 font-sans text-sm outline-none transition-all" style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"/> <button${attr("disabled", creating, true)} class="bg-orange-500 text-white font-mono font-bold text-sm uppercase tracking-widest py-3 px-6 transition-all disabled:opacity-50" style="border: 2.5px solid var(--border); box-shadow: 4px 4px 0 var(--border);">${escape_html("Shorten →")}</button></div> `);
    {
      $$renderer2.push("<!--[!-->");
    }
    $$renderer2.push(`<!--]--></div> `);
    if (links.length === 0) {
      $$renderer2.push("<!--[1-->");
      $$renderer2.push(`<div class="p-12 text-center" style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"><p class="font-mono text-sm font-bold uppercase tracking-widest" style="color: var(--text-muted);">// No links yet</p> <p class="font-mono text-xs mt-2" style="color: var(--text-muted);">Create your first short link above</p></div>`);
    } else {
      $$renderer2.push("<!--[!-->");
      $$renderer2.push(`<div class="flex flex-col gap-3"><!--[-->`);
      const each_array = ensure_array_like(links);
      for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
        let link = each_array[$$index];
        $$renderer2.push(`<div class="p-5 flex items-center justify-between gap-4 transition-all" style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"><div class="flex flex-col gap-1 min-w-0"><a${attr("href", `${stringify(BASE_URL)}/${stringify(link.short_code)}`)} target="_blank" rel="noopener noreferrer" class="font-mono text-sm font-bold text-orange-500 hover:underline">shorten.krishnasaimarupudi.com/${escape_html(link.short_code)}</a> <p class="text-sm truncate" style="color: var(--text-muted);">${escape_html(link.original_url)}</p> <p class="font-mono text-xs" style="color: var(--text-muted);">${escape_html(new Date(link.created_at).toLocaleDateString())}</p></div> <div class="flex items-center gap-4 shrink-0"><div class="text-center"><p class="font-mono text-xs uppercase tracking-widest" style="color: var(--text-muted);">Clicks</p> <p class="font-mono text-lg font-bold">${escape_html(link.clicks ?? 0)}</p></div> <button class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all"${attr_style(`border: 2px solid var(--border); color: ${stringify(copied === link.id ? "#22c55e" : "var(--text)")}; border-color: ${stringify(copied === link.id ? "#22c55e" : "var(--border)")};`)}>${escape_html(copied === link.id ? "✓ Copied!" : "Copy")}</button> <button class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 hover:bg-red-500 hover:text-white transition-all" style="border: 2px solid var(--border);">Delete</button></div></div>`);
      }
      $$renderer2.push(`<!--]--></div>`);
    }
    $$renderer2.push(`<!--]--></div></div>`);
    if ($$store_subs) unsubscribe_stores($$store_subs);
  });
}

export { _page as default };
//# sourceMappingURL=_page.svelte-DbuiqJ86.js.map
