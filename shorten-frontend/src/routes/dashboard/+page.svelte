<script>
  import { BASE_URL } from "$lib/config.js";
  import { token, logout } from "$lib/stores/auth.js";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { get } from "svelte/store";
  import { theme, toggleTheme } from "$lib/stores/theme.js";
  let copied = $state(null);

  async function copyToClipboard(id, shortCode) {
    await navigator.clipboard.writeText(`${BASE_URL}/${shortCode}`);
    copied = id;
    setTimeout(() => (copied = null), 2000);
  }

  let links = $state([]);
  let newURL = $state("");
  let error = $state("");
  let loading = $state(false);
  let creating = $state(false);

  function handleLogout() {
    logout();
    goto("/");
  }

  async function fetchLinks() {
    loading = true;
    try {
      const response = await fetch(`${BASE_URL}/api/links`, {
        headers: {
          Authorization: `Bearer ${get(token)}`,
        },
      });
      if (!response.ok) throw new Error("Failed to fetch links");
      const data = await response.json();
      links = data.data ?? [];
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }

  async function createLink() {
    if (!newURL) return;
    creating = true;
    error = "";
    try {
      const response = await fetch(`${BASE_URL}/api/links`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${get(token)}`,
        },
        body: JSON.stringify({ url: newURL }),
      });
      if (!response.ok) throw new Error("Failed to create link");
      newURL = "";
      await fetchLinks();
    } catch (err) {
      error = err.message;
    } finally {
      creating = false;
    }
  }

  async function deleteLink(id) {
    try {
      const response = await fetch(`${BASE_URL}/api/links/${id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${get(token)}`,
        },
      });
      if (!response.ok) throw new Error("Failed to delete link");
      await fetchLinks();
    } catch (err) {
      error = err.message;
    }
  }

  onMount(() => {
    const t = localStorage.getItem("token");
    if (!t) {
      goto("/");
      return;
    }
    theme.subscribe((t) => {
      document.documentElement.setAttribute("data-theme", t);
    });
    fetchLinks();
  });
</script>

<div class="min-h-screen" style="background: var(--bg); color: var(--text);">
  <nav
    class="sticky top-0 z-10"
    style="background: var(--bg); border-bottom: 2.5px solid var(--border);"
  >
    <div class="max-w-5xl mx-auto px-6 h-16 flex items-center justify-between">
      <span class="font-mono text-xl font-bold"
        >KSM<span class="text-orange-500">.</span>shorten</span
      >
      <div class="flex items-center gap-3">
        <button
          onclick={toggleTheme}
          class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all"
          style="border: 2px solid var(--border); color: var(--text);"
        >
          {#if $theme === "dark"}☀️{:else}🌙{/if}
        </button>
        <button
          onclick={handleLogout}
          class="font-mono text-xs font-bold uppercase tracking-widest px-4 py-2 transition-all"
          style="border: 2px solid var(--border); color: var(--text);"
        >
          Logout
        </button>
      </div>
    </div>
  </nav>

  <div class="max-w-5xl mx-auto px-6 py-12">
    <div class="mb-10">
      <p
        class="font-mono text-xs font-bold tracking-widest uppercase text-orange-500 mb-2"
      >
        // Dashboard
      </p>
      <h1 class="font-mono text-5xl font-bold leading-none">
        Your<br />
        <span style="-webkit-text-stroke: 2px var(--text); color: transparent;"
          >Links</span
        >
      </h1>
    </div>

    <div
      class="p-6 mb-10"
      style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"
    >
      <h2
        class="font-mono text-sm font-bold uppercase tracking-widest mb-4"
        style="color: var(--text);"
      >
        // New Link
      </h2>
      <div class="flex gap-3">
        <input
          type="url"
          bind:value={newURL}
          placeholder="https://example.com"
          class="flex-1 p-3 font-sans text-sm outline-none transition-all"
          style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"
        />
        <button
          onclick={createLink}
          disabled={creating}
          class="bg-orange-500 text-white font-mono font-bold text-sm uppercase tracking-widest py-3 px-6 transition-all disabled:opacity-50"
          style="border: 2.5px solid var(--border); box-shadow: 4px 4px 0 var(--border);"
        >
          {creating ? "Creating..." : "Shorten →"}
        </button>
      </div>
      {#if error}
        <p class="font-mono text-xs text-red-500 font-bold mt-3">✗ {error}</p>
      {/if}
    </div>

    {#if loading}
      <div
        class="font-mono text-sm text-center py-12"
        style="color: var(--text-muted);"
      >
        Loading...
      </div>
    {:else if links.length === 0}
      <div
        class="p-12 text-center"
        style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"
      >
        <p
          class="font-mono text-sm font-bold uppercase tracking-widest"
          style="color: var(--text-muted);"
        >
          // No links yet
        </p>
        <p class="font-mono text-xs mt-2" style="color: var(--text-muted);">
          Create your first short link above
        </p>
      </div>
    {:else}
      <div class="flex flex-col gap-3">
        {#each links as link}
          <div
            class="p-5 flex items-center justify-between gap-4 transition-all"
            style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"
          >
            <div class="flex flex-col gap-1 min-w-0">
              <a
                href="{BASE_URL}/{link.short_code}"
                target="_blank"
                rel="noopener noreferrer"
                class="font-mono text-sm font-bold text-orange-500 hover:underline"
              >
                shorten.krishnasaimarupudi.com/{link.short_code}
              </a>
              <p class="text-sm truncate" style="color: var(--text-muted);">
                {link.original_url}
              </p>
              <p class="font-mono text-xs" style="color: var(--text-muted);">
                {new Date(link.created_at).toLocaleDateString()}
              </p>
            </div>
            <div class="flex items-center gap-4 shrink-0">
              <div class="text-center">
                <p
                  class="font-mono text-xs uppercase tracking-widest"
                  style="color: var(--text-muted);"
                >
                  Clicks
                </p>
                <p class="font-mono text-lg font-bold">{link.clicks ?? 0}</p>
              </div>
              <button
                onclick={() => copyToClipboard(link.id, link.short_code)}
                class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all"
                style="border: 2px solid var(--border); color: {copied ===
                link.id
                  ? '#22c55e'
                  : 'var(--text)'}; border-color: {copied === link.id
                  ? '#22c55e'
                  : 'var(--border)'};"
              >
                {copied === link.id ? "✓ Copied!" : "Copy"}
              </button>
              <button
                onclick={() => deleteLink(link.id)}
                class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 hover:bg-red-500 hover:text-white transition-all"
                style="border: 2px solid var(--border);"
              >
                Delete
              </button>
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
