<script lang='ts'>
    import { page } from '$app/stores';
    import { onDestroy, onMount } from 'svelte';
    import { closeSocket, gameID, joinGame, socket, message } from '../../../stores/websocket';
    import axios from 'axios';
    import { writable } from 'svelte/store';

    const baseWsUrl = import.meta.env.VITE_WS_URL;
    const baseApiUrl = import.meta.env.VITE_API_URL;

    let currentMessage = '';
    let currentGameID = '';

    type GameInfos = {
        id: string;
        gameLeaderID: number;
        mapHeight: number;
        mapWidth: number;
        nbPlayer: number;
    };

    let gameInfos: GameInfos = {
        id: '',
        gameLeaderID: 0,
        mapHeight: 20,
        mapWidth: 20,
        nbPlayer: 0,
    };

    onMount(async () => {
        currentGameID = $page.params.gameID;
        gameID.set(currentGameID);
        joinGame(currentGameID, baseWsUrl);

        try {
            const response = await axios.get(`${baseApiUrl}/game/${currentGameID}/infos`, {withCredentials: true});
            gameInfos = response.data;

            grid.set(
                Array.from({ length: gameInfos.mapHeight }, (_, y) =>
                    Array.from({ length: gameInfos.mapWidth }, (_, x) => ({
                        x,
                        y,
                        state: 'empty',
                    }))
                )
            );
        } catch (error) {
            console.error(error);
        }
    });

    onDestroy(() => {
        closeSocket();
    });

    function sendMessage() {
        if ($socket && $socket.readyState === WebSocket.OPEN) {
            $socket.send(currentMessage);
        } else {
            console.warn("La connexion WebSocket est fermée");
        }
    }

    // GAME
    let grid = writable<{x: number,y: number,state: string,}[][]>([]);

    grid.subscribe((value) => {
        value = Array.from({ length: gameInfos.mapHeight }, (_, y) =>
            Array.from({ length: gameInfos.mapWidth }, (_, x) => ({
                x,
                y,
                state: 'empty'
            }))
        );
    });

    let cellSize = 30;

    function updateCellState(x: number, y: number, newState: string) {
        $grid[y][x].state = newState;
    }
</script>

<div class="flex flex-col items-center gap-4">
    <h1>TEST WEB SOCKET</h1>

    <p>gameID : {currentGameID}</p>
    <p>{JSON.stringify(gameInfos)}</p>

    <div class="flex gap-4">
        <input
            class="input input-bordered input-primary"
            type="text"
            bind:value={currentMessage}
            placeholder="Entrez un message"
        />
        <button class="btn" on:click={sendMessage}>Envoyer un message</button>
    </div>

    <button class="btn" on:click={closeSocket}>Disconnect</button>

    <!-- Conteneur pour la grille -->
    <div
        class="grid bg-neutral p-2 rounded-lg min-w-96 max-w-[90vw] max-h-[90vh]"
        style="grid-template-rows: repeat({gameInfos.mapHeight}, {cellSize}px); grid-template-columns: repeat({gameInfos.mapWidth}, {cellSize}px);"
    >
        {#each $grid as row}
            {#each row as cell}
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                <!-- svelte-ignore a11y_no_static_element_interactions -->
                <div
                    class="aspect-square bg-base-300 border border-base-200 cursor-pointer transition-all"
                    class:bg-success={cell.state === 'snake'}
                    class:bg-error={cell.state === 'food'}
                    class:hover={cell.state === 'empty'}
                    on:click={() => updateCellState(cell.x, cell.y, 'snake')}
                    style="width: {cellSize}px; height: {cellSize}px;"
                ></div>
            {/each}
        {/each}
    </div>

</div>
