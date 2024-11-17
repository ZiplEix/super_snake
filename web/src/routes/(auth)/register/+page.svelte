<script lang='ts'>
    import axios from "axios";
    import { writable } from "svelte/store";
    import { userStore } from "../../../stores/user_store";

    const baseApiUrl = import.meta.env.VITE_API_URL;

    let name = '';
    let email = '';
    let password = '';
    let errorMessage = writable('');

    const register = async (e: Event) => {
        e.preventDefault();

        try {
            errorMessage.set('');

            const response = await axios.post(
                `${baseApiUrl}/register`,
                {
                    name,
                    email,
                    password,
                },
                {
                    withCredentials: true,
                }
            )

            if (response.status === 201) {
                // localStorage.setItem('jwt', response.data.token);
                userStore.set(response.data.user);
                window.location.href = '/dashboard';
            } else {
                errorMessage.set(response.data.message);
            }
        } catch (error: any) {
            errorMessage.set(error.response.data.message);
        }
    };
</script>

<div class="grid h-screen place-items-center">
    <div class="flex flex-col items-center space-y-4">
        <h1 class="text-5xl font-bold mb-16">Register</h1>
        <div class="card w-96 bg-base-100 shadow-xl custom-card-bg">
            <div class="card-body">
                <form class="flex flex-col items-center space-y-4">
                    <input
                        type="text"
                        placeholder="Name"
                        class="input input-bordered input-primary w-full max-w-xs"
                        id="name"
                        name="name"
                        bind:value={name}
                        required
                    />
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
                    <button type="submit" class="btn btn-primary" on:click={register}>Register</button>
                </form>
            </div>
        </div>
        <div class="text-center text-sm mt-4">
            <a href="login" class="hover:underline flex">
                <span>Already have an account?&nbsp;</span>
                <span class="text-blue-500">Sign in</span>
            </a>
        </div>
    </div>
</div>
