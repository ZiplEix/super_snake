<script lang='ts'>
    import { logout } from "$lib";
</script>

<nav class="navbar bg-base-200">
    <div class="flex-1">
        <a href="/dashboard" class="btn btn-ghost text-xl">SuperSnake</a>
    </div>
    <div class="flex-none">
        <div class="dropdown dropdown-end">
        </div>
        <div class="dropdown dropdown-end">
            <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                <div class="w-10 rounded-full">
                    <img
                        alt="Tailwind CSS Navbar component"
                        src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
                    />
                </div>
            </div>
            <ul
                class="menu menu-sm dropdown-content bg-base-200 rounded-box z-[1] mt-3 w-52 p-2 shadow"
            >
                <li><a href="/dashboard">Settings</a></li>
                <li><button on:click={logout}>Logout</button></li>
            </ul>
        </div>
    </div>
</nav>
