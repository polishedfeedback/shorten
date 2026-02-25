<script>
  import { BASE_URL } from "$lib/config.js";
  import { login } from "$lib/stores/auth.js";
  import { theme, toggleTheme } from "$lib/stores/theme.js";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  let username = $state("");
  let password = $state("");
  let error = $state("");
  let loading = $state(false);

  onMount(() => {
    theme.subscribe((t) => {
      document.documentElement.setAttribute("data-theme", t);
    });
  });

  async function handleSubmit(e) {
    e.preventDefault();
    loading = true;
    error = "";
    try {
      const response = await fetch(`${BASE_URL}/api/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });
      if (!response.ok) throw new Error("Invalid credentials");
      const data = await response.json();
      login(data.token);
      goto("/dashboard");
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>

<div
  class="min-h-screen flex items-center justify-center p-4"
  style="background: var(--bg); color: var(--text);"
>
  <!-- Theme toggle top right -->
  <div class="fixed top-4 right-4">
    <button
      onclick={toggleTheme}
      class="font-mono text-xs font-bold uppercase tracking-widest px-3 py-2 transition-all"
      style="border: 2px solid var(--border); color: var(--text);"
    >
      {#if $theme === "dark"}☀️{:else}🌙{/if}
    </button>
  </div>

  <div class="w-full max-w-md">
    <!-- Header -->
    <div class="mb-8">
      <p
        class="font-mono text-xs font-bold tracking-widest uppercase text-orange-500 mb-2"
      >
        // Link Shortener
      </p>
      <h1 class="font-mono text-4xl font-bold leading-none">
        Sign<br />
        <span style="-webkit-text-stroke: 2px var(--text); color: transparent;"
          >In</span
        >
      </h1>
    </div>

    <!-- Card -->
    <div
      class="p-8"
      style="background: var(--card-bg); border: 2.5px solid var(--border); box-shadow: 5px 5px 0 var(--border);"
    >
      {#if error}
        <div
          class="bg-red-500 border-2 border-red-700 text-white font-mono text-sm font-bold p-3 mb-6"
        >
          ✗ {error}
        </div>
      {/if}

      <form onsubmit={handleSubmit} class="flex flex-col gap-5">
        <div class="flex flex-col gap-1.5">
          <label
            for="username"
            class="font-mono text-xs font-bold uppercase tracking-widest"
            style="color: var(--text-muted);">Username</label
          >
          <input
            id="username"
            type="text"
            bind:value={username}
            placeholder="your username"
            required
            class="p-3 font-sans text-sm outline-none transition-all"
            style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"
          />
        </div>

        <div class="flex flex-col gap-1.5">
          <label
            for="password"
            class="font-mono text-xs font-bold uppercase tracking-widest"
            style="color: var(--text-muted);">Password</label
          >
          <input
            id="password"
            type="password"
            bind:value={password}
            placeholder="your password"
            required
            class="p-3 font-sans text-sm outline-none transition-all"
            style="background: var(--bg); border: 2px solid var(--border); color: var(--text);"
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          class="bg-orange-500 text-white font-mono font-bold text-sm uppercase tracking-widest py-3 px-6 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
          style="border: 2.5px solid var(--border); box-shadow: 4px 4px 0 var(--border);"
        >
          {loading ? "Signing in..." : "Sign In →"}
        </button>
      </form>
    </div>

    <p
      class="font-mono text-xs text-center mt-4"
      style="color: var(--text-muted);"
    >
      shorten.krishnasaimarupudi.com
    </p>
  </div>
</div>
