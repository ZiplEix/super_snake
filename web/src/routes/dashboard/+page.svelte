<script lang='ts'>
  import Loader from "$lib/ui/Loader.svelte";
import { writable } from "svelte/store";

    let gamecode = '';

    let createLoading = writable(false);
    let joinLoading = writable(false);

    function createGame() {
        createLoading.set(true);
        setTimeout(() => {
            createLoading.set(false);
        }, 2000);
    }

    function joinGame() {
        joinLoading.set(true);
        setTimeout(() => {
            joinLoading.set(false);
        }, 2000);
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
