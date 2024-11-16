<script lang='ts'>
  import { goto } from "$app/navigation";
  import { logout } from "$lib";
    import Loader from "$lib/ui/Loader.svelte";
    import axios from "axios";
    import { writable } from "svelte/store";

    const baseApiUrl = import.meta.env.VITE_API_URL;

    let gamecode = '';

    let createLoading = writable(false);
    let joinLoading = writable(false);

    async function createGame() {
        createLoading.set(true);

        try {
            const response = await axios.get(`${baseApiUrl}/game/create`, {withCredentials: true});

            goto(`/snake/${response.data.gameID}`);
        } catch (error: any) {
            console.error(error);
            if (error.response.status === 401) {
                logout();
            } else {
                alert('An error occurred while creating the game');
            }
        }

        createLoading.set(false);
    }

    function joinGame() {
        joinLoading.set(true);

        goto(`/snake/${gamecode}`);

        joinLoading.set(false);
    }
</script>

<div class="grid h-96 place-items-center">
    <div class="flex flex-col items-center space-y-4">
        <h1 class="text-5xl font-bold mb-4">Dashboard</h1>
        <button class="btn w-full" on:click={createGame}>
            {#if $createLoading}
                <Loader />
            {:else}
                CREATE GAME
            {/if}
        </button>
        <div class="divider">OR</div>
        <div class="flex flex-row gap-4">
            <input
                type="email"
                placeholder="XXXXXX"
                class="input input-bordered input-primary w-full"
                id="email"
                name="email"
                bind:value={gamecode}
                required
            />
            <button class="btn" on:click={joinGame}>
                {#if $joinLoading}
                    <Loader />
                {:else}
                    JOIN GAME
                {/if}
            </button>
        </div>
    </div>
</div>
