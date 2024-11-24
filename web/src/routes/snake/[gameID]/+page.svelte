<script lang='ts'>
    import { page } from '$app/stores';
    import { onDestroy, onMount } from 'svelte';
    import { closeSocket, gameID, joinGame, socket } from '../../../stores/websocket';
    import axios from 'axios';
    import { writable } from 'svelte/store';
    import Loader from '$lib/ui/Loader.svelte';
    import { userStore } from '../../../stores/user_store';

    const baseWsUrl = import.meta.env.VITE_WS_URL;
    const baseApiUrl = import.meta.env.VITE_API_URL;

    let currentGameID = '';

    type GameInfos = {
        id: string;
        gameLeaderID: number;
        mapHeight: number;
        mapWidth: number;
        nbPlayer: number;
        nbPlayerMax: number;
        gameState: number;
    };

    let gameInfos = writable<GameInfos>({
        id: '',
        gameLeaderID: 0,
        mapHeight: 20,
        mapWidth: 20,
        nbPlayer: 0,
        nbPlayerMax: 0,
        gameState: 0,
    });

    function messageHandler(event: MessageEvent) {
        if (event.data === 'ping') {
            // console.log('Ping received');
            return;
        }

        const message = JSON.parse(event.data);
        console.log('Message received:', message);

        switch (message.type) {
            case 'GameStart':
                gameInfos.update((value) => {
                    value.gameState = 1;
                    return value;
                });
                break;
            case 'StartBoard':
                console.log('StartBoard', message.data);
                grid.update((currentGrid) => {
                    message.data.forEach((cellUpdate: any) => {
                        const { x, y, content } = cellUpdate;
                        if (content.type === 1) {
                            currentGrid[y][x].state = 'snake';
                            currentGrid[y][x].color = content.color;
                        } else if (content.type === 2) {
                            currentGrid[y][x].state = 'food';
                        } else {
                            currentGrid[y][x].state = 'empty';
                        }
                    });

                    return currentGrid
                })
                break;
            default:
                console.warn('Unknown message type:', message.type);
        }
    }

    onMount(async () => {
        currentGameID = $page.params.gameID;
        gameID.set(currentGameID);
        joinGame(currentGameID, baseWsUrl, messageHandler);

        try {
            const response = await axios.get(`${baseApiUrl}/game/${currentGameID}/infos`, {withCredentials: true});
            gameInfos.set(response.data);

            grid.set(
                Array.from({ length: $gameInfos.mapHeight }, (_, y) =>
                    Array.from({ length: $gameInfos.mapWidth }, (_, x) => ({
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

    type GameMessage = {
        type: string;
        data: any;
    };

    function sendMessage(message: GameMessage) {
        console.log('Sending message:', message);
        if ($socket && $socket.readyState === WebSocket.OPEN) {
            $socket.send(JSON.stringify(message));
        } else {
            console.warn("La connexion WebSocket est fermée");
        }
    }

    // GAME
    type Cell = {
        x: number;
        y: number;
        state: string;
        color?: string;
    };
    let grid = writable<Cell[][]>([]);

    grid.subscribe((value) => {
        value = Array.from({ length: $gameInfos.mapHeight }, (_, y) =>
            Array.from({ length: $gameInfos.mapWidth }, (_, x) => ({
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

    <h1>{JSON.stringify($gameInfos)}</h1>

    {#if $gameInfos.gameState === 0}
        <div class="absolute top-0 left-0 w-full h-full bg-black bg-opacity-50 flex flex-col items-center justify-center z-10 gap-8">
            <h2 class="text-white text-2xl">Waiting for the leader to start the game</h2>
            <p>{$gameInfos.nbPlayer}/{$gameInfos.nbPlayerMax} player, waiting for {$gameInfos.nbPlayerMax - $gameInfos.nbPlayer} more player</p>

            {#if $gameInfos.gameLeaderID === $userStore?.id}
                <button class="btn btn-primary" on:click={() => sendMessage({type: "game_control", data: {"action": "start"}})}>Start the game</button>
            {:else}
                <Loader />
            {/if}
        </div>
    {/if}

    <!-- Conteneur du jeu -->
    <div
        class="grid bg-neutral p-2 rounded-lg min-w-96"
        style="grid-template-rows: repeat({$gameInfos.mapHeight}, {cellSize}px); grid-template-columns: repeat({$gameInfos.mapWidth}, {cellSize}px);"
    >
        {#each $grid as row}
            {#each row as cell}
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                <!-- svelte-ignore a11y_no_static_element_interactions -->
                <div
                    class="aspect-square bg-base-300 border border-base-200 transition-all"
                    style="width: {cellSize}px; height: {cellSize}px; background-color: {cell.color};"
                >
                    {#if cell.state === 'food'}
                        <img src="/default_food.png" alt="Food" class="w-full h-full p-1" />
                    {/if}
                </div>
                <!-- <div
                    class="aspect-square bg-base-300 border border-base-200 cursor-pointer transition-all"
                    class:bg-success={cell.state === 'snake'}
                    class:bg-error={cell.state === 'food'}
                    class:hover={cell.state === 'empty'}
                    on:click={() => updateCellState(cell.x, cell.y, 'snake')}
                    style="width: {cellSize}px; height: {cellSize}px;"
                ></div> -->
            {/each}
        {/each}
    </div>
</div>
