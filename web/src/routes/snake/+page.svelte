<script lang="ts">
    let socket: WebSocket | null = null;
    let message = "";

    // Fonction de connexion WebSocket
    function connectWebSocket() {
        if (typeof window !== "undefined") {
            // Cette vérification garantit que WebSocket est exécuté uniquement dans le navigateur
            socket = new WebSocket("ws://localhost:8080/ws");

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

    // Envoyer un message au serveur
    function sendMessage() {
        if (socket && socket.readyState === WebSocket.OPEN) {
            // const message = "Hello from SvelteKit!";
            socket.send(message);
            console.log("Message envoyé au serveur:", message);
        } else {
            console.warn("La connexion WebSocket est fermée");
        }
    }

    // Se connecter au serveur à l'initialisation
    // connectWebSocket();
</script>

<h1>TEST WEB SOCKET</h1>

<!-- <input type="text" bind:value={message} placeholder="Entrez un message" />

<button class="btn" on:click={sendMessage}>Envoyer un message</button> -->
