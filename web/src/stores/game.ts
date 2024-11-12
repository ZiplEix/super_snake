import { writable } from "svelte/store";

export const rows = 20;
export const cols = 20;

export const snake = writable([{ x: 10, y: 10 }]);

export const apple = writable(generateRandomApplePosition());

export const direction = writable({ x: 1, y: 0 });

function generateRandomApplePosition(): { x: number; y: number } {
    const x = Math.floor(Math.random() * cols);
    const y = Math.floor(Math.random() * rows);
    return { x, y };
}

export function moveSnake() {
    snake.update((segments) => {
        

        // Récupérer la direction actuelle
        const { x: dx, y: dy } = $direction;

        // Calculer la nouvelle position de la tête
        const newHead = {
            x: segments[0].x + dx,
            y: segments[0].y + dy
        };

        // Vérifier les collisions avec les murs
        if (newHead.x < 0 || newHead.x >= cols || newHead.y < 0 || newHead.y >= rows) {
            resetGame();
            return segments;
        }

        // Ajouter la nouvelle tête au début du serpent
        const newSnake = [newHead, ...segments];

        // Vérifier si le serpent mange la pomme
        if (newHead.x === $apple.x && newHead.y === $apple.y) {
            apple.set(generateRandomApplePosition()); // Générer une nouvelle pomme
            score.update(n => n + 10); // Augmenter le score
        } else {
            // Retirer la dernière partie du corps si le serpent n'a pas mangé
            newSnake.pop();
        }

        return newSnake;
    });
}