<script lang='ts'>
    // get the gameID from the URL
    import { page } from '$app/stores';
    import { onDestroy, onMount } from 'svelte';
    import { closeSocket, gameID, joinGame, socket, message } from '../../../stores/websocket';
  import axios from 'axios';

    const baseWsUrl = import.meta.env.VITE_WS_URL;
    const baseApiUrl = import.meta.env.VITE_API_URL;

    let currentMessage = '';
    let currentGameID = '';

    type GameInfos = {
        id: string;
        gameLeaderID: number;
    };

    let gameInfos: GameInfos = {
        id: '',
        gameLeaderID: 0,
    };

    onMount(async () => {
        currentGameID = $page.params.gameID;
        gameID.set(currentGameID);
        joinGame(currentGameID, baseWsUrl);

        try {
            const response = await axios.get(`${baseApiUrl}/game/${currentGameID}/infos`, {withCredentials: true});
            gameInfos = response.data;
        } catch (error) {
            console.error(error);
        }
    });

    onDestroy(() => {
        closeSocket();
    });

    function sendMessage() {
        if ($socket && $socket.readyState === WebSocket.OPEN) {
            $socket.send($message);
        } else {
            console.warn("La connexion WebSocket est fermée");
        }
    }
</script>

<h1>TEST WEB SOCKET</h1>

<p>gameID : {currentGameID}</p>

<p>{JSON.stringify(gameInfos)}</p>

<br>

<div class="flex gap-4">
    <input
        class="input input-bordered input-primary"
        type="text"
        bind:value={currentMessage}
        placeholder="Entrez un message"
    />
    <button class="btn" on:click={sendMessage}>Envoyer un message</button>
</div>

<br>

<button class="btn" on:click={closeSocket}>Disconnect</button>
