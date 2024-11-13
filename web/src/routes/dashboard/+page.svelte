<script lang='ts'>
    let gamecode = '';
</script>

<div class="grid h-screen place-items-center">
    <div class="flex flex-col items-center space-y-4">
        <h1 class="text-5xl font-bold mb-4">Dashboard</h1>
        <button class="btn w-full">CREATE GAME</button>
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
            <button class="btn">JOIN GAME</button>
        </div>
    </div>
</div>