<script lang='ts'>
    import NavBar from "$lib/ui/NavBar.svelte";

    import { onMount } from "svelte";
    import { userStore } from "../../stores/user_store";
    import { verifyToken } from "$lib/verify_token";
    import { redirect } from "@sveltejs/kit";
    import { goto } from "$app/navigation";

    onMount(async () => {
        const user = await verifyToken();

        if (!user) {
            userStore.set(null);
            goto('/login');
        } else {
            userStore.set(user);
        }
    });
</script>

<NavBar />

<slot />
