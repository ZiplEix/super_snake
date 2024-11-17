<script lang='ts'>
    import axios from "axios";
    import { writable } from "svelte/store";
    import { userStore } from "../../../stores/user_store";

    const baseApiUrl = import.meta.env.VITE_API_URL;

    let email = '';
    let password = '';
    let errorMessage = writable('');
    let loading = writable(false);

    const login = async (e: Event) => {
        e.preventDefault();

        errorMessage.set('');
        loading.set(true);

        try {
            const response = await axios.post(
                `${baseApiUrl}/login`,
                {
                    email,
                    password,
                },
                {
                    withCredentials: true,
                }
            )

            if (response.status === 200) {
                // localStorage.setItem('jwt', response.data.token);
                userStore.set(response.data.user);
                window.location.href = '/dashboard';
            } else {
                errorMessage.set(response.data.message);
            }
        } catch (error: any) {
            errorMessage.set(error.response.data.message);
        } finally {
            loading.set(false);
        }
    };
</script>

<div class="grid h-screen place-items-center">
    <div class="flex flex-col items-center space-y-4">
        <h1 class="text-5xl font-bold mb-16">Sign In</h1>
        <div class="card w-96 bg-base-100 shadow-xl custom-card-bg">
            <div class="card-body">
                <form class="flex flex-col items-center space-y-4">
                    <input
                        type="email"
                        placeholder="Email"
                        class="input input-bordered input-primary w-full max-w-xs"
                        id="email"
                        name="email"
                        bind:value={email}
                        required
                    />
                    <input
                        type="password"
                        placeholder="Password"
                        class="input input-bordered input-primary w-full max-w-xs"
                        id="password"
                        name="password"
                        bind:value={password}
                        required
                    />
                    {#if $errorMessage}
                        <div class="mt-2 text-red-500">
                            <p>{$errorMessage}</p>
                        </div>
                    {/if}
                    <button type="submit" class="btn btn-primary h-full" on:click={login}>
                        {#if $loading}
                            <div class="mt-4 flex justify-center">
                                <div class="w-8 h-8 border-t-4 border-blue-300 border-solid rounded-full animate-spin"></div>
                            </div>
                        {:else}
                            Sign In
                        {/if}
                    </button>
                </form>
            </div>
        </div>
        <div class="text-center text-sm mt-4">
            <a href="register" class="hover:underline flex">
                <span>No account yet?&nbsp;</span>
                <span class="text-blue-500">Sign up</span>
            </a>
        </div>
    </div>
</div>
