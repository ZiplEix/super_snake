<script lang='ts'>
    // get the gameID from the URL
    import { page } from '$app/stores';
    import { onDestroy, onMount } from 'svelte';
    import { closeSocket, gameID, joinGame, socket, message } from '../../../stores/websocket';

    const baseApiUrl = import.meta.env.VITE_WS_URL;

    let currentMessage = '';
    let currentGameID = '';

    onMount(() => {
        currentGameID = $page.params.gameID;
        gameID.set(currentGameID);
        joinGame(currentGameID, baseApiUrl);
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
