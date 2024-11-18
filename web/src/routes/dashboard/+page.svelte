<script lang='ts'>
    import { goto } from "$app/navigation";
    import { logout } from "$lib";
    import Loader from "$lib/ui/Loader.svelte";
    import axios from "axios";
    import { writable } from "svelte/store";
    import { userStore } from "../../stores/user_store";

    const baseApiUrl = import.meta.env.VITE_API_URL;

    let gamecode = '';

    let createLoading = writable(false);
    let joinLoading = writable(false);
    const user = userStore;

    let nbPlayerMaxRangeStyle = writable("");
    let mapHeightInputStyle = writable("input-primary");
    let mapWidthInputStyle = writable("input-primary");

    let gameParams = writable({
        nbPlayerMax: 2,
        mapHeight: 20,
        mapWidth: 20,
    });

    gameParams.subscribe((value) => {
        if (value.nbPlayerMax < 2) {
            nbPlayerMaxRangeStyle.set("range-error");
        } else {
            nbPlayerMaxRangeStyle.set("");
        }

        if (value.mapHeight < 10) {
            mapHeightInputStyle.set("input-error");
        } else {
            mapHeightInputStyle.set("input-primary");
        }

        if (value.mapWidth < 10) {
            mapWidthInputStyle.set("input-error");
        } else {
            mapWidthInputStyle.set("input-primary");
        }
    });

    async function createGame(params: App.GameParamsRequest) {
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
        <h2 class="text-3xl">Welcome {$user?.name}</h2>

        <button class="btn w-full" on:click={() => {createGameModal.showModal()}}>
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

<dialog id="createGameModal" class="modal modal-bottom sm:modal-middle">
    <div class="modal-box space-y-6">
        <!-- Titre -->
        <h2 class="text-3xl font-bold text-center">Game Settings</h2>

        <!-- Contenu -->
        <div class="space-y-6">
            <!-- Max Players -->
            <div class="flex flex-col space-y-4">
                <label class="text-lg font-medium">Max Players:</label>
                <div class="flex flex-col space-y-2">
                    <!-- Slider -->
                    <input
                        type="range"
                        min="1"
                        max="4"
                        bind:value={$gameParams.nbPlayerMax}
                        class="range {$nbPlayerMaxRangeStyle}"
                        step="1"
                    />
                    <div class="flex w-full justify-between px-4 text-sm text-gray-600">
                        {#each {length: 4} as _, i}
                            <span>{i + 1}</span>
                        {/each}
                    </div>
                </div>
            </div>

            <!-- Map Size Inputs -->
            <div class="grid grid-cols-2 gap-4">
                <!-- Map Height -->
                <div class="flex flex-col space-y-2">
                    <label class="text-lg font-medium">Map Height:</label>
                    <input
                        type="number"
                        min="10"
                        max="50"
                        bind:value={$gameParams.mapHeight}
                        class="input input-bordered {$mapHeightInputStyle}"
                    />
                </div>
                <!-- Map Width -->
                <div class="flex flex-col space-y-2">
                    <label class="text-lg font-medium">Map Width:</label>
                    <input
                        type="number"
                        min="10"
                        max="50"
                        bind:value={$gameParams.mapWidth}
                        class="input input-bordered {$mapWidthInputStyle}"
                    />
                </div>
            </div>
        </div>

        <div class="modal-action">
            <form method="dialog">
                <button class="btn btn-outline">Close</button>
                <button class="btn btn-primary">Create Game</button>
            </form>
        </div>
    </div>
</dialog>
