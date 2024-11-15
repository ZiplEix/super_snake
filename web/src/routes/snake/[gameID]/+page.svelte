<script lang='ts'>
    // get the gameID from the URL
    import { page } from '$app/stores';
    import { onDestroy, onMount } from 'svelte';
    import { writable } from 'svelte/store';

    const baseApiUrl = import.meta.env.VITE_WS_URL;

    let gameID = '';
    let message = '';
    let socket: WebSocket | null = null;

    onMount(() => {
        gameID = $page.params.gameID;
        console.log('gameID:', gameID);

        joinGame(gameID);
    });

    onDestroy(() => {
        if (socket) {
            socket.close();
            socket = null;
            console.log('Connexion WebSocket fermée');
        }
    });

    async function joinGame(gameID: string) {
        if (typeof window !== "undefined") {
            socket = new WebSocket(`${baseApiUrl}/ws/${gameID}`);

            socket.onopen = () => {
                console.log("Connexion WebSocket établie");
            };

            socket.onmessage = (event) => {
                console.log("Message reçu du serveur:", event.data);
            };

            socket.onerror = (error) => {
                console.error("Erreur WebSocket:", error);
            };

            socket.onclose = () => {
                console.log("Connexion WebSocket fermée");
            };
        }
    }

    function sendMessage() {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(message);
            console.log("Message envoyé au serveur:", message);
        } else {
            console.warn("La connexion WebSocket est fermée");
        }
    }
</script>

<h1>TEST WEB SOCKET</h1>

<p>gameID : {gameID}</p>

<br>

<div class="flex gap-4">
    <input class="input input-bordered input-primary" type="text" bind:value={message} placeholder="Entrez un message" />
    <button class="btn" on:click={sendMessage}>Envoyer un message</button>
</div>
