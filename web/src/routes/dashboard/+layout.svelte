<script context="module" lang='ts'>
    export async function load({ url }: { url: URL }) {
        return {
            queryParam: url.searchParams.get('redirectTo') || '/'
        }
    }
</script>

<script lang='ts'>
  import NavBar from "$lib/ui/NavBar.svelte";

    import { onMount } from "svelte";

    onMount(() => {
        const jwt = localStorage.getItem('jwt');

        if (!jwt) {
            window.location.href = 'login';
        }
    });
</script>

<NavBar />

<slot />
